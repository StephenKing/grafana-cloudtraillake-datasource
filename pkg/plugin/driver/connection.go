package driver

import (
	"context"
	"database/sql/driver"
	"fmt"
	"github.com/aws/aws-sdk-go/service/cloudtrail"
	"github.com/emnify/cloud-trail-lake/pkg/plugin/api"
	"github.com/grafana/grafana-aws-sdk/pkg/awsds"
	sqlAPI "github.com/grafana/grafana-aws-sdk/pkg/sql/api"
	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
)

var _ awsds.AsyncDB = &conn{}

type conn struct {
	api    *api.API
	closed bool
}

func newConnection(api *api.API) *conn {
	return &conn{
		api: api,
	}
}

func (c *conn) StartQuery(ctx context.Context, query string, args ...interface{}) (string, error) {
	output, err := c.api.Execute(ctx, &sqlAPI.ExecuteQueryInput{Query: query})
	if err != nil {
		return "", err
	}
	log.DefaultLogger.Warn("Submitted query", "queryId", output.ID, "query", query)
	return output.ID, nil
}

func (c *conn) GetQueryID(ctx context.Context, query string, args ...interface{}) (bool, string, error) {
	return c.api.GetQueryID(ctx, query, args)
}

func (c *conn) QueryStatus(ctx context.Context, queryID string) (awsds.QueryStatus, error) {
	status, err := c.api.Status(ctx, &sqlAPI.ExecuteQueryOutput{ID: queryID})
	if err != nil {
		return awsds.QueryUnknown, err
	}
	var returnStatus awsds.QueryStatus
	switch status.State {
	case cloudtrail.QueryStatusQueued:
		returnStatus = awsds.QuerySubmitted
	case cloudtrail.QueryStatusRunning:
		returnStatus = awsds.QueryRunning
	case cloudtrail.QueryStatusFinished:
		returnStatus = awsds.QueryFinished
	case cloudtrail.QueryStatusCancelled:
		returnStatus = awsds.QueryCanceled
	case cloudtrail.QueryStatusFailed:
		returnStatus = awsds.QueryFailed
	case cloudtrail.QueryStatusTimedOut:
		returnStatus = awsds.QueryFailed
	}
	backend.Logger.Warn("QueryStatus", "state", status.State, "queryID", queryID)
	return returnStatus, nil
}

func (c *conn) CancelQuery(ctx context.Context, queryID string) error {
	log.DefaultLogger.Warn("Query canceled", "queryId", queryID)
	return c.api.Stop(&sqlAPI.ExecuteQueryOutput{ID: queryID})

}

func (c *conn) GetRows(ctx context.Context, queryID string) (driver.Rows, error) {
	return NewRows(ctx, c.api.Client, queryID)
}

func (c *conn) Ping(ctx context.Context) error {
	_, err := c.api.Execute(ctx, &sqlAPI.ExecuteQueryInput{Query: "SELECT 1"})
	if err != nil {
		return err
	}
	return nil
}

func (c *conn) Begin() (driver.Tx, error) {
	return nil, fmt.Errorf("CloudTrail Lake driver doesn't support begin statements")
}

func (c *conn) Prepare(query string) (driver.Stmt, error) {
	return nil, fmt.Errorf("CloudTrail Lake driver doesn't support prepared statements")
}

func (c *conn) Close() error {
	c.closed = true
	return nil
}
