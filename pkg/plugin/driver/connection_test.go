package driver

//func TestConnection_QueryContext(t *testing.T) {
//	c := &conn{
//		api: api.NewFake(ctlclientmock.MockCtlClient{CalledTimesCountDown: 1},
//			&models.CloudTrailLakeDataSourceSettings{
//				AWSDatasourceSettings: awsds.AWSDatasourceSettings{},
//			}),
//	}
//
//	failedOutput, err := c.StartQuery(context.Background(), ctlclientmock.FAKE_ERROR, []driver.NamedValue{})
//	if !errors.Is(err, sqlAPI.ExecuteError) {
//		t.Errorf("unexpected err %v", err)
//	}
//	assert.Equal(t, failedOutput, "")
//
//	_, err = c.StartQuery(context.Background(), ctlclientmock.FAKE_SUCCESS, []driver.NamedValue{})
//	assert.Equal(t, err, nil)
//}
