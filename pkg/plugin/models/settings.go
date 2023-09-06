package models

import (
	"encoding/json"
	"fmt"
	"github.com/grafana/grafana-aws-sdk/pkg/awsds"
	"github.com/grafana/grafana-aws-sdk/pkg/sql/models"
	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/sqlds/v2"
)

const (
	Region = "region"
)

type CloudTrailLakeDataSourceSettings struct {
	awsds.AWSDatasourceSettings
	Config backend.DataSourceInstanceSettings
}

func New() models.Settings {
	return &CloudTrailLakeDataSourceSettings{}
}

func (s *CloudTrailLakeDataSourceSettings) Load(config backend.DataSourceInstanceSettings) error {
	if config.JSONData != nil && len(config.JSONData) > 1 {
		if err := json.Unmarshal(config.JSONData, s); err != nil {
			return fmt.Errorf("could not unmarshal DatasourceSettings json: %w", err)
		}
	}

	s.AccessKey = config.DecryptedSecureJSONData["accessKey"]
	s.SecretKey = config.DecryptedSecureJSONData["secretKey"]

	s.Config = config

	return nil
}

func (s *CloudTrailLakeDataSourceSettings) Apply(args sqlds.Options) {
	region := args[Region]
	if region != "" {
		if region == models.DefaultKey {
			s.Region = s.DefaultRegion
		} else {
			s.Region = region
		}
	}

}
