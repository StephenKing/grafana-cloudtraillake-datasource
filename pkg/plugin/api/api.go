package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/aws/aws-sdk-go/service/cloudtrail"
	"github.com/emnify/cloud-trail-lake/pkg/plugin/models"
	"github.com/grafana/grafana-aws-sdk/pkg/awsds"
	"github.com/grafana/grafana-aws-sdk/pkg/sql/api"
	sqlModels "github.com/grafana/grafana-aws-sdk/pkg/sql/models"
	"github.com/grafana/grafana-plugin-sdk-go/backend"
	sdkhttpclient "github.com/grafana/grafana-plugin-sdk-go/backend/httpclient"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"github.com/grafana/sqlds/v2"
	"strings"
)

type API struct {
	Client   *cloudtrail.CloudTrail
	settings *models.CloudTrailLakeDataSourceSettings
}

func New(sessionCache *awsds.SessionCache, settings sqlModels.Settings) (api.AWSAPI, error) {
	ctlSettings := settings.(*models.CloudTrailLakeDataSourceSettings)

	httpClientProvider := sdkhttpclient.NewProvider()
	httpClientOptions, err := ctlSettings.Config.HTTPClientOptions()
	if err != nil {
		backend.Logger.Error("failed to create HTTP client options", "error", err.Error())
		return nil, err
	}
	httpClient, err := httpClientProvider.New(httpClientOptions)
	if err != nil {
		backend.Logger.Error("failed to create HTTP client", "error", err.Error())
		return nil, err
	}

	sess, err := sessionCache.GetSession(awsds.SessionConfig{
		HTTPClient:    httpClient,
		Settings:      ctlSettings.AWSDatasourceSettings,
		UserAgentName: aws.String("Athena"),
	})
	if err != nil {
		return nil, err
	}

	return &API{Client: cloudtrail.New(sess), settings: ctlSettings}, nil
}

func (c *API) Execute(ctx context.Context, input *api.ExecuteQueryInput) (*api.ExecuteQueryOutput, error) {
	ctlInput := &cloudtrail.StartQueryInput{
		QueryStatement: aws.String(input.Query),
	}

	output, err := c.Client.StartQueryWithContext(ctx, ctlInput)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", api.ExecuteError, err)
	}

	return &api.ExecuteQueryOutput{ID: *output.QueryId}, nil
}

/// FIXME TODO check if this is also true for CTL (comment still from Athena DS)

// GetQueryID always returns not found. To actually check if the query has already been called would require calling
// ListQueryExecutions, which has a limit of 5 calls per second. This leads to throttling when there are many panels
// and/or many query executions to page through
func (c *API) GetQueryID(ctx context.Context, query string, args ...interface{}) (bool, string, error) {
	return false, "", nil
}

func (c *API) Status(ctx aws.Context, output *api.ExecuteQueryOutput) (*api.ExecuteQueryStatus, error) {
	statusResp, err := c.Client.DescribeQueryWithContext(ctx, &cloudtrail.DescribeQueryInput{
		QueryId: aws.String(output.ID),
	})
	if err != nil {
		return nil, fmt.Errorf("%w: %v", api.ExecuteError, err)
	}

	var finished bool
	state := *statusResp.QueryStatus
	switch state {
	case cloudtrail.QueryStatusFailed, cloudtrail.QueryStatusCancelled, cloudtrail.QueryStatusTimedOut:
		finished = true
		err = errors.New(*statusResp.ErrorMessage)
	case cloudtrail.QueryStatusFinished:
		finished = true
	default:
		finished = false
	}
	return &api.ExecuteQueryStatus{
		ID:       output.ID,
		State:    state,
		Finished: finished,
	}, err
}

func (c *API) CancelQuery(ctx context.Context, options sqlds.Options, queryID string) error {
	return c.Stop(&api.ExecuteQueryOutput{ID: queryID})
}

func (c *API) Stop(output *api.ExecuteQueryOutput) error {
	_, err := c.Client.CancelQuery(&cloudtrail.CancelQueryInput{
		QueryId: &output.ID,
	})
	if err != nil {
		return fmt.Errorf("%w: unable to stop query", err)
	}
	return nil
}

// regions from https://docs.aws.amazon.com/general/latest/gr/athena.html
var standardRegions = []string{
	"af-south-1",
	"ap-east-1",
	"ap-northeast-1",
	"ap-northeast-2",
	"ap-northeast-3",
	"ap-south-1",
	"ap-southeast-1",
	"ap-southeast-2",
	"ca-central-1",
	"eu-central-1",
	"eu-north-1",
	"eu-south-1",
	"eu-west-1",
	"eu-west-2",
	"eu-west-3",
	"me-south-1",
	"sa-east-1",
	"us-east-1",
	"us-east-2",
	"us-gov-east-1",
	"us-gov-west-1",
	"us-west-1",
	"us-west-2",
}

func (c *API) Regions(aws.Context) ([]string, error) {
	return standardRegions, nil
}

func (c *API) getOptionWithDefault(options sqlds.Options, option string) string {
	v, ok := options[option]
	if !ok {
		return ""
	}
	if v == sqlModels.DefaultKey {
		switch option {
		case "region":
			v = c.settings.DefaultRegion
		}
	}
	return v
}

func (c *API) EventDataStores(ctx aws.Context) ([]models.EventDataStore, error) {
	log.DefaultLogger.Debug("Listing Event Data Stores")

	dbResp, err := c.Client.ListEventDataStoresWithContext(ctx, &cloudtrail.ListEventDataStoresInput{})
	if err != nil {
		return nil, fmt.Errorf("%w: %v", api.ExecuteError, err)
	}

	res := []models.EventDataStore{}
	for _, eventStore := range dbResp.EventDataStores {
		// we only want the resource part of the EDS ARN
		edsArn, _ := arn.Parse(aws.StringValue(eventStore.EventDataStoreArn))

		name := *eventStore.Name
		id := strings.Replace(edsArn.Resource, "eventdatastore/", "", 1)

		log.DefaultLogger.Debug("Found EDS", "name", name, "id", id)

		res = append(res, models.EventDataStore{
			Name: name,
			Id:   id,
		})
	}
	return res, nil
}

func (c *API) Databases(ctx aws.Context, options sqlds.Options) ([]string, error) {
	// TODO is that enough?
	res := []string{}
	res = append(res, "mytestsdatabase")
	return res, nil
}

func (c *API) Schemas(ctx aws.Context, options sqlds.Options) ([]string, error) {
	// TODO is that enough?
	res := []string{}
	res = append(res, "mytestschema")
	return res, nil
}
func (c *API) Tables(ctx aws.Context, options sqlds.Options) ([]string, error) {
	// TODO is that enough?
	res := []string{}
	res = append(res, "mytesttable")
	return res, nil
}

func (c *API) Columns(ctx aws.Context, options sqlds.Options) ([]string, error) {
	// TODO is that enough?
	res := []string{}
	res = append(res, "mytestcolumn")
	return res, nil
}
