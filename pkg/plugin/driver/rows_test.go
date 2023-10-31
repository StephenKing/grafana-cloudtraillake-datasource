package driver

import (
	"context"
	"database/sql/driver"
	"fmt"
	"io"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	ctlservicemock "github.com/emnify/cloud-trail-lake/pkg/plugin/driver/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOnePageSuccess(t *testing.T) {
	ctlServiceMock := &ctlservicemock.CtlService{}
	ctlServiceMock.CalledTimesCountDown = 1
	rows, rowErr := NewRows(context.Background(), ctlServiceMock, ctlservicemock.SinglePageResponseQueryId)
	require.NoError(t, rowErr)
	cnt := 0
	for {
		var col1, col2 string
		err := rows.Next([]driver.Value{
			&col1,
			&col2,
		})
		if err != nil {
			require.ErrorIs(t, io.EOF, err)
			break
		}
		require.NoError(t, err)
		cnt++
	}
	require.Equal(t, 2, cnt)
}

func TestMultiPageSuccess(t *testing.T) {
	ctlServiceMock := &ctlservicemock.CtlService{}
	ctlServiceMock.CalledTimesCountDown = 5
	rows, rowErr := NewRows(context.Background(), ctlServiceMock, ctlservicemock.MultiPageResponseQueryId)
	require.NoError(t, rowErr)
	cnt := 0
	for {
		var col1, col2 string
		err := rows.Next([]driver.Value{
			&col1,
			&col2,
		})
		if err != nil {
			require.ErrorIs(t, io.EOF, err)
			break
		}
		require.NoError(t, err)
		cnt++
	}
	require.Equal(t, 10, cnt)
	require.Equal(t, 5, ctlServiceMock.CalledTimesCounter)
}

func Test_convertRow(t *testing.T) {

	tests := []struct {
		name          string
		data          *map[string]*string
		expectedType  string
		expectedValue string
	}{
		{
			name: "numeric type int",
			data: &map[string]*string{
				"my-int": aws.String("1"),
			},
			expectedType:  "int32",
			expectedValue: "1",
		},
		//{
		//	name: "numeric type float4",
		//	metadata: &redshiftdataapiservice.ColumnMetadata{
		//		Name:     aws.String("other"),
		//		TypeName: aws.String(REDSHIFT_FLOAT4),
		//	},
		//	data: &redshiftdataapiservice.Field{
		//		DoubleValue: aws.Float64(1.1),
		//	},
		//	expectedType:  "float64",
		//	expectedValue: "1.1",
		//},
		//{
		//	name: "numeric type float",
		//	metadata: &redshiftdataapiservice.ColumnMetadata{
		//		Name:     aws.String("other"),
		//		TypeName: aws.String(REDSHIFT_FLOAT),
		//	},
		//	data: &redshiftdataapiservice.Field{
		//		DoubleValue: aws.Float64(1.3),
		//	},
		//	expectedType:  "float64",
		//	expectedValue: "1.3",
		//},
		{
			name: "string",
			data: &map[string]*string{
				"my-string": aws.String("1"),
			},
			expectedType:  "string",
			expectedValue: "foo",
		},
		{
			name: "date",
			// FIXME should have a column called 'time'
			data: &map[string]*string{
				"my-date": aws.String("2008-01-01"),
			},
			expectedType:  "time.Time",
			expectedValue: "2008-01-01 00:00:00 +0000 UTC",
		},
		{
			name: "timestamp",
			// FIXME should have a column called 'time'
			data: &map[string]*string{
				"my-timestamp": aws.String("2008-01-01 20:00:00.000"),
			},
			expectedType:  "time.Time",
			expectedValue: "2008-01-01 20:00:00 +0000 UTC",
		},
		{
			name:          "null",
			data:          nil,
			expectedType:  "<nil>",
			expectedValue: "<nil>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := make([]driver.Value, 1)
			err := convertRow(
				[]map[string]*string{*tt.data},
				res,
			)
			require.NoError(t, err)
			assert.Equal(t, tt.expectedType, fmt.Sprintf("%T", res[0]))
			assert.Equal(t, tt.expectedValue, fmt.Sprintf("%v", res[0]))
		})
	}
}
