package driver

import (
	"context"
	"database/sql/driver"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudtrail"
	"github.com/aws/aws-sdk-go/service/cloudtrail/cloudtrailiface"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"io"
	"time"
)

type Rows struct {
	service cloudtrailiface.CloudTrailAPI
	queryID string

	done   bool
	result *cloudtrail.GetQueryResultsOutput
}

func NewRows(ctx context.Context, service cloudtrailiface.CloudTrailAPI, queryID string) (*Rows, error) {
	r := Rows{
		service: service,
		queryID: queryID,
	}

	if err := r.fetchNextPage(nil); err != nil {
		return nil, err
	}
	return &r, nil
}

// fetchNextPage fetches the next statement result page and adds the result to the row
func (r *Rows) fetchNextPage(token *string) error {
	var err error

	r.result, err = r.service.GetQueryResults(&cloudtrail.GetQueryResultsInput{
		QueryId:   aws.String(r.queryID),
		NextToken: token,
	})

	if err != nil {
		return err
	}

	return nil
}

// Close closes the rows iterator.
func (r *Rows) Close() error {
	r.done = true
	return nil
}

// Next is called to populate the next row of data into
// the provided slice. The provided slice will be the same
// size as the Columns() are wide. io.EOF should be returned when there are no more rows.
func (r *Rows) Next(dest []driver.Value) error {
	if r.done {
		log.DefaultLogger.Info("Next done")
		return io.EOF
	}

	// If nothing left to iterate...
	if len(r.result.QueryResultRows) == 0 {
		// And if nothing more to paginate...
		if r.result.NextToken == nil || *r.result.NextToken == "" {
			r.done = true
			return io.EOF
		}

		err := r.fetchNextPage(r.result.NextToken)
		if err != nil {
			return err
		}
	}

	current := r.result.QueryResultRows[0]
	err := convertRow(current, dest)
	if err != nil {
		return err
	}

	r.result.QueryResultRows = r.result.QueryResultRows[1:]
	return nil
}

func convertRow(data []map[string]*string, ret []driver.Value) error {
	// Read the row and iterate over all the column->value pairs to store the value
	i := 0
	for _, row := range data {
		for key, value := range row {
			if key == "eventTime" || key == "time" {
				log.DefaultLogger.Warn("Attempting to parse datetime", "key", key, "value", value)
				t, err := time.Parse("2006-01-02 15:04:05", *value)
				if err != nil {
					return err
				}
				ret[i] = t
			} else {
				ret[i] = *value
			}
			i++
		}
	}
	return nil
}

// Columns returns the names of the columns.
func (r *Rows) Columns() []string {
	columnNames := []string{}

	if len(r.result.QueryResultRows) == 0 {
		return columnNames
	}

	for _, columnData := range r.result.QueryResultRows[0] {
		for key, _ := range columnData {
			columnNames = append(columnNames, key)
		}
	}
	return columnNames
}
