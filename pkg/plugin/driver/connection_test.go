package driver

import (
	"context"
	"database/sql/driver"
	"errors"
	"testing"

	"github.com/emnify/cloud-trail-lake/pkg/plugin/api"
	ctlclientmock "github.com/emnify/cloud-trail-lake/pkg/plugin/api/mock"
	"github.com/emnify/cloud-trail-lake/pkg/plugin/models"
	"github.com/grafana/grafana-aws-sdk/pkg/awsds"
	sqlAPI "github.com/grafana/grafana-aws-sdk/pkg/sql/api"
	"gotest.tools/assert"
)

func TestConnection_QueryContext(t *testing.T) {
	c := &conn{
		api: api.NewFake(ctlclientmock.MockCtlClient{CalledTimesCountDown: 1},
			&models.CloudTrailLakeDataSourceSettings{
				AWSDatasourceSettings: awsds.AWSDatasourceSettings{},
			}),
	}

	failedOutput, err := c.StartQuery(context.Background(), ctlclientmock.FAKE_ERROR, []driver.NamedValue{})
	if !errors.Is(err, sqlAPI.ExecuteError) {
		t.Errorf("unexpected err %v", err)
	}
	assert.Equal(t, failedOutput, "")

	_, err = c.StartQuery(context.Background(), ctlclientmock.FAKE_SUCCESS, []driver.NamedValue{})
	assert.Equal(t, err, nil)
}
