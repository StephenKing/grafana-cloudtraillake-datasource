package mock

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/cloudtrail"
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
	cloudtrail.CloudTrail
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
		NextToken:       input.NextToken,
		QueryResultRows: [][]map[string]*string{},
	}

	//if *input.QueryId == ROWS_WITH_NEXT {
	//	next := "oneMorePage"
	//	output.NextToken = &next
	//	fakeVarChar := "someString"
	//	fakeDatum := aws.String(fakeVarChar)
	//	fakeColumnName := "_col0"
	//	fakeColumn := athena.Datum{
	//		VarCharValue: &fakeColumnName,
	//	}
	//	output.QueryResultRows = append(output.QueryResultRows,
	//		[]map[string]*string{{"foo": aws.String("bar")}},
	//	)
	//		// FIXME
	//	//	[][]map[string]*string
	//	//	&athena.Row{
	//	//		Data: []*athena.Datum{&fakeColumn},
	//	//	},
	//	//	&athena.Row{
	//	//		Data: []*athena.Datum{&fakeDatum},
	//	//	},
	//	)
	//	fakeNullable := "NULLABLE"
	//	fakePrecision := int64(1)
	//	fakeType := "varchar"
	//	fakeName := "name"
	//	fakeColumnInfo := athena.ColumnInfo{
	//		Name:      &fakeName,
	//		Nullable:  &fakeNullable,
	//		Precision: &fakePrecision,
	//		Type:      &fakeType,
	//	}
	//	output.ResultSet.ResultSetMetadata.ColumnInfo = []*athena.ColumnInfo{&fakeColumnInfo}
	//}

	return output, nil
}

func (m *MockCtlClient) GetQueryResultsWithContext(ctx context.Context, input *cloudtrail.GetQueryResultsInput, opts ...request.Option) (*cloudtrail.GetQueryResultsOutput, error) {
	return &cloudtrail.GetQueryResultsOutput{
		QueryResultRows: [][]map[string]*string{0: {0: {"foo": aws.String("bar")}}},
	}, nil
}

func (m *MockCtlClient) StopQueryExecution(input *cloudtrail.CancelQueryInput) (*cloudtrail.CancelQueryOutput, error) {
	m.Cancelled = true
	return &cloudtrail.CancelQueryOutput{}, nil
}
