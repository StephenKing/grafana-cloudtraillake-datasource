package plugin

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/emnify/cloud-trail-lake/pkg/plugin/api"
	"github.com/emnify/cloud-trail-lake/pkg/plugin/driver"
	"github.com/emnify/cloud-trail-lake/pkg/plugin/models"
	"github.com/grafana/grafana-aws-sdk/pkg/awsds"
	sqlAPI "github.com/grafana/grafana-aws-sdk/pkg/sql/api"
	"github.com/grafana/grafana-aws-sdk/pkg/sql/datasource"
	asyncDriver "github.com/grafana/grafana-aws-sdk/pkg/sql/driver/async"
	sqlModels "github.com/grafana/grafana-aws-sdk/pkg/sql/models"
	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/instancemgmt"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"github.com/grafana/grafana-plugin-sdk-go/data"
	"github.com/grafana/grafana-plugin-sdk-go/data/sqlutil"
	"github.com/grafana/sqlds/v2"
	"github.com/pkg/errors"
)

// // Make sure CtlDatasource implements required interfaces. This is important to do
// // since otherwise we will only get a not implemented error response from plugin in
// // runtime. In this example datasource instance implements backend.QueryDataHandler,
// // backend.CheckHealthHandler interfaces. Plugin should not implement all these
// // interfaces- only those which are required for a particular task.
//var (
//	_ backend.QueryDataHandler   = (*CtlDatasource)(nil)
//	_ backend.CheckHealthHandler = (*CtlDatasource)(nil)
//)

type CtlDatasourceIface interface {
	sqlds.Driver
	sqlds.Completable
	sqlAPI.Resources
	awsds.AsyncDriver
}

type awsDSClient interface {
	Init(config backend.DataSourceInstanceSettings)
	GetAsyncDB(id int64, options sqlds.Options, settingsLoader sqlModels.Loader, apiLoader sqlAPI.Loader, driverLoader asyncDriver.Loader) (awsds.AsyncDB, error)
	GetAPI(id int64, options sqlds.Options, settingsLoader sqlModels.Loader, apiLoader sqlAPI.Loader) (sqlAPI.AWSAPI, error)
}

// NewDatasource creates a new datasource instance.
// FIXME needed?
func NewDatasource(_ backend.DataSourceInstanceSettings) (instancemgmt.Instance, error) {
	log.DefaultLogger.Info("datasource / NewDataSource")
	return &CtlDatasource{}, nil
}

//	func New() CtlDatasourceIface {
//		return &CtlDatasource{awsDS: datasource.New()}
//	}
func New() *CtlDatasource {
	return &CtlDatasource{awsDS: datasource.New()}
}

// CtlDatasource is an example datasource which can respond to data queries, reports
// its health and has streaming skills.
type CtlDatasource struct {
	awsDS awsDSClient
}

func (s *CtlDatasource) Settings(_ backend.DataSourceInstanceSettings) sqlds.DriverSettings {
	return sqlds.DriverSettings{
		FillMode: &data.FillMissing{
			Mode: data.FillModeNull,
		},
	}
}

func (s *CtlDatasource) Converters() (sc []sqlutil.Converter) {
	return sc
}

// Connect opens a sql.DB connection using datasource settings
func (s *CtlDatasource) Connect(config backend.DataSourceInstanceSettings, queryArgs json.RawMessage) (*sql.DB, error) {

	log.DefaultLogger.Info("XXXXXX datasource / Connect")
	// TODO what to do with this method? Only needed for sync queries?
	return nil, errors.New("did not expect to call Connect method")
}

func (s *CtlDatasource) GetAsyncDB(config backend.DataSourceInstanceSettings, queryArgs json.RawMessage) (awsds.AsyncDB, error) {
	log.DefaultLogger.Info("XXXXXX datasource / GetAsyncDB")

	s.awsDS.Init(config)
	args, err := sqlds.ParseOptions(queryArgs)
	if err != nil {
		return nil, err
	}
	args["updated"] = config.Updated.String()

	return s.awsDS.GetAsyncDB(config.ID, args, models.New, api.New, driver.New)
}

func (s *CtlDatasource) getApi(ctx context.Context, options sqlds.Options) (*api.API, error) {
	log.DefaultLogger.Info("XXXXXX datasource / getApi")

	id := datasource.GetDatasourceID(ctx)
	args := sqlds.Options{}
	for key, val := range options {
		args[key] = val
	}
	// the updated time makes sure that we don't use a token for a stale version of the datasource
	args["updated"] = datasource.GetDatasourceLastUpdatedTime(ctx)

	res, err := s.awsDS.GetAPI(id, args, models.New, api.New)
	if err != nil {
		return nil, err
	}

	return res.(*api.API), err
}

func (s *CtlDatasource) Regions(ctx context.Context) ([]string, error) {
	api, err := s.getApi(ctx, sqlds.Options{})
	if err != nil {
		return nil, err
	}
	regions, err := api.Regions(ctx)
	if err != nil {
		return nil, err
	}
	return regions, nil
}

func (s *CtlDatasource) Databases(ctx context.Context, options sqlds.Options) ([]string, error) {
	api, err := s.getApi(ctx, options)
	if err != nil {
		return nil, err
	}
	dbs, err := api.Databases(ctx, options)
	if err != nil {
		return nil, err
	}
	return dbs, nil
}

func (s *CtlDatasource) CancelQuery(ctx context.Context, options sqlds.Options, queryID string) error {
	api, err := s.getApi(ctx, options)
	if err != nil {
		return err
	}
	return api.CancelQuery(ctx, options, queryID)
}

func (s *CtlDatasource) Schemas(ctx context.Context, options sqlds.Options) ([]string, error) {
	api, err := s.getApi(ctx, options)
	if err != nil {
		return nil, err
	}
	schemas, err := api.Schemas(ctx, options)
	if err != nil {
		return nil, err
	}
	return schemas, nil
}

func (s *CtlDatasource) Tables(ctx context.Context, options sqlds.Options) ([]string, error) {
	api, err := s.getApi(ctx, options)
	if err != nil {
		return nil, err
	}
	tables, err := api.Tables(ctx, options)
	if err != nil {
		return nil, err
	}
	return tables, nil
}

func (s *CtlDatasource) Columns(ctx context.Context, options sqlds.Options) ([]string, error) {
	api, err := s.getApi(ctx, options)
	if err != nil {
		return nil, err
	}
	cols, err := api.Columns(ctx, options)
	if err != nil {
		return nil, err
	}
	return cols, nil
}

// Dispose here tells plugin SDK that plugin wants to clean up resources when a new instance
// created. As soon as datasource settings change detected by SDK old datasource instance will
// be disposed and a new one will be created using NewSampleDatasource factory function.
//func (d *CtlDatasource) Dispose() {
//	// Clean up datasource instance resources.
//}

//// QueryData handles multiple queries and returns multiple responses.
//// req contains the queries []DataQuery (where each query contains RefID as a unique identifier).
//// The QueryDataResponse contains a map of RefID to the response for each query, and each response
//// contains Frames ([]*Frame).
//func (d *CtlDatasource) QueryData(ctx context.Context, req *backend.QueryDataRequest) (*backend.QueryDataResponse, error) {
//	// create response struct
//	response := backend.NewQueryDataResponse()
//
//	// loop over queries and execute them individually.
//	for _, q := range req.Queries {
//		res := d.query(ctx, req.PluginContext, q)
//
//		// save the response in a hashmap
//		// based on with RefID as identifier
//		response.Responses[q.RefID] = res
//	}
//
//	return response, nil
//}

//func (d *CtlDatasource) query(_ context.Context, pCtx backend.PluginContext, query backend.DataQuery) backend.DataResponse {
//	var response backend.DataResponse
//
//	// Unmarshal the JSON into our queryModel.
//	var qm queryModel
//
//	err := json.Unmarshal(query.JSON, &qm)
//	if err != nil {
//		return backend.ErrDataResponse(backend.StatusBadRequest, fmt.Sprintf("json unmarshal: %v", err.Error()))
//	}
//
//	// create data frame response.
//	// For an overview on data frames and how grafana handles them:
//	// https://grafana.com/docs/grafana/latest/developers/plugins/data-frames/
//	frame := data.NewFrame("response")
//
//	// add fields.
//	frame.Fields = append(frame.Fields,
//		data.NewField("time", nil, []time.Time{query.TimeRange.From, query.TimeRange.To}),
//		data.NewField("content", nil, []string{"Foo", "Bar"}),
//	)
//
//	// add the frames to the response.
//	response.Frames = append(response.Frames, frame)
//
//	return response
//}

// CheckHealth handles health checks sent from Grafana to the plugin.
// The main use case for these health checks is the test button on the
// datasource configuration page which allows users to verify that
// a datasource is working as expected.
func (d *CtlDatasource) CheckHealth(_ context.Context, req *backend.CheckHealthRequest) (*backend.CheckHealthResult, error) {
	var status = backend.HealthStatusOk
	var message = "Data source is working"

	return &backend.CheckHealthResult{
		Status:  status,
		Message: message,
	}, nil
}
