package mock

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/athena"
	"github.com/aws/aws-sdk-go/service/cloudtrail"
	"github.com/aws/aws-sdk-go/service/cloudtrail/cloudtrailiface"
)

const DESCRIBE_STATEMENT_FAILED = "DESCRIBE_STATEMENT_FAILED"
const DESCRIBE_STATEMENT_SUCCEEDED = "DESCRIBE_STATEMENT_FINISHED"

// Define a mock struct to be used in your unit tests of myFunc.
type MockCtlClient struct {
	CalledTimesCounter   int
	CalledTimesCountDown int
	Databases            []string
	Columns              []string
	Cancelled            bool
	cloudtrailiface.CloudTrailAPI
}

func (m *MockCtlClient) GetQueryExecutionWithContext(ctx aws.Context, input *cloudtrail.GetQueryResultsInput, opts ...request.Option) (*cloudtrail.GetQueryResultsOutput, error) {
	// mock response/functionality
	m.CalledTimesCountDown--
	m.CalledTimesCounter++

	output := &cloudtrail.GetQueryResultsOutput{}
	if m.CalledTimesCountDown == 0 {
		if *input.QueryId == DESCRIBE_STATEMENT_FAILED {
			output.QueryStatus = aws.String(cloudtrail.QueryStatusFailed)
		} else {
			output.QueryStatus = aws.String(cloudtrail.QueryStatusFinished)
		}
	} else {
		output.QueryStatus = aws.String(cloudtrail.QueryStatusRunning)
	}
	return output, nil
}

const FAKE_ERROR = "FAKE_ERROR"
const FAKE_SUCCESS = "FAKE_SUCCESS"

func (m *MockCtlClient) StartQueryExecutionWithContext(ctx aws.Context, input *cloudtrail.StartQueryInput, opts ...request.Option) (*cloudtrail.StartQueryOutput, error) {
	output := &cloudtrail.StartQueryOutput{
		QueryId: input.QueryStatement,
	}
	if *input.QueryStatement == FAKE_ERROR {
		return nil, errors.New(FAKE_ERROR)
	}

	return output, nil
}

const EMPTY_ROWS = "EMPTY_ROWS"
const ROWS_WITH_NEXT = "RowsWithNext"

func (m *MockCtlClient) GetQueryResults(input *cloudtrail.GetQueryResultsInput) (*cloudtrail.GetQueryResultsOutput, error) {
	if *input.QueryId == FAKE_ERROR {
		return nil, errors.New(FAKE_ERROR)
	}

	output := &cloudtrail.GetQueryResultsOutput{
		NextToken: input.NextToken,
		QueryResultRows: [][]map[string]*string{},
	}

	if *input.QueryId == ROWS_WITH_NEXT {
		next := "oneMorePage"
		output.NextToken = &next
		fakeVarChar := "someString"
		fakeDatum := athena.Datum{
			VarCharValue: &fakeVarChar,
		}
		fakeColumnName := "_col0"
		fakeColumn := athena.Datum{
			VarCharValue: &fakeColumnName,
		}
		output.QueryResultRows = append(output.QueryResultRows,
			[][]map[string]*string
			&athena.Row{
				Data: []*athena.Datum{&fakeColumn},
			},
			&athena.Row{
				Data: []*athena.Datum{&fakeDatum},
			},
		)
		fakeNullable := "NULLABLE"
		fakePrecision := int64(1)
		fakeType := "varchar"
		fakeName := "name"
		fakeColumnInfo := athena.ColumnInfo{
			Name:      &fakeName,
			Nullable:  &fakeNullable,
			Precision: &fakePrecision,
			Type:      &fakeType,
		}
		output.ResultSet.ResultSetMetadata.ColumnInfo = []*athena.ColumnInfo{&fakeColumnInfo}
	}

	return output, nil
}

func (m *MockCtlClient) GetQueryResultsWithContext(ctx context.Context, input *athena.GetQueryResultsInput, opts ...request.Option) (*athena.GetQueryResultsOutput, error) {
	return &athena.GetQueryResultsOutput{
		ResultSet: &athena.ResultSet{
			ResultSetMetadata: &athena.ResultSetMetadata{},
			Rows:              []*athena.Row{},
		},
	}, nil
}

func (m *MockCtlClient) ListDataCatalogsWithContext(ctx aws.Context, input *athena.ListDataCatalogsInput, opts ...request.Option) (*athena.ListDataCatalogsOutput, error) {
	r := &athena.ListDataCatalogsOutput{}
	for _, c := range m.Catalogs {
		r.DataCatalogsSummary = append(r.DataCatalogsSummary, &athena.DataCatalogSummary{CatalogName: aws.String(c)})
	}
	return r, nil
}

func (m *MockCtlClient) ListDatabasesWithContext(ctx aws.Context, input *athena.ListDatabasesInput, opts ...request.Option) (*athena.ListDatabasesOutput, error) {
	r := &athena.ListDatabasesOutput{}
	for _, c := range m.Databases {
		r.DatabaseList = append(r.DatabaseList, &athena.Database{Name: aws.String(c)})
	}
	return r, nil
}

func (m *MockCtlClient) ListWorkGroupsWithContext(ctx aws.Context, input *athena.ListWorkGroupsInput, opts ...request.Option) (*athena.ListWorkGroupsOutput, error) {
	r := &athena.ListWorkGroupsOutput{}
	for _, c := range m.Workgroups {
		r.WorkGroups = append(r.WorkGroups, &athena.WorkGroupSummary{Name: aws.String(c)})
	}
	return r, nil
}

func (m *MockCtlClient) ListTableMetadataWithContext(ctx aws.Context, input *athena.ListTableMetadataInput, opts ...request.Option) (*athena.ListTableMetadataOutput, error) {
	r := &athena.ListTableMetadataOutput{}
	for _, c := range m.TableMetadataList {
		r.TableMetadataList = append(r.TableMetadataList, &athena.TableMetadata{Name: aws.String(c)})
	}
	return r, nil
}

func (m *MockCtlClient) GetTableMetadataWithContext(ctx aws.Context, input *athena.GetTableMetadataInput, opts ...request.Option) (*athena.GetTableMetadataOutput, error) {
	r := &athena.GetTableMetadataOutput{}
	r.TableMetadata = &athena.TableMetadata{Name: aws.String("fake table metadata")}
	for _, c := range m.Columns {
		r.TableMetadata.Columns = append(r.TableMetadata.Columns, &athena.Column{Name: aws.String(c)})
	}
	return r, nil
}

func (m *MockCtlClient) StopQueryExecution(input *athena.StopQueryExecutionInput) (*athena.StopQueryExecutionOutput, error) {
	m.Cancelled = true
	return &athena.StopQueryExecutionOutput{}, nil
}
