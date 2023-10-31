package api

import (
	"github.com/aws/aws-sdk-go/service/cloudtrail"
	"github.com/emnify/cloud-trail-lake/pkg/plugin/models"
)

// NewFake returns an API object with the given args
func NewFake(cli cloudtrail.CloudTrail, settings *models.CloudTrailLakeDataSourceSettings) *API {
	return &API{Client: &cli, settings: settings}
}
