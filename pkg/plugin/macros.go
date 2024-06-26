package plugin

import (
	"fmt"
	"github.com/emnify/cloud-trail-lake/pkg/plugin/models"
	"strings"
	"time"

	"github.com/grafana/grafana-plugin-sdk-go/backend/gtime"
	"github.com/grafana/sqlds/v2"
	"github.com/pkg/errors"
	"github.com/viant/toolbox"
)

const (
	timestampFormat   = "'yyyy-MM-dd HH:mm:ss'"
	goTimestampFormat = "2006-01-02 15:04:05"
)

func parseTime(target, format string) string {
	if format == "" {
		return target
	} else if format == timestampFormat {
		return fmt.Sprintf("TIMESTAMP %s", target)
	}
	return fmt.Sprintf("parse_datetime(%s,%s)", target, format)
}

func parseTimeGroup(query *models.Query, args []string) (time.Duration, string, error) {
	if len(args) < 2 {
		return 0, "", errors.WithMessagef(sqlds.ErrorBadArgumentCount, "macro $__timeGroup needs time column and interval")
	}

	interval, err := gtime.ParseInterval(strings.Trim(args[1], `'`))
	if err != nil {
		return 0, "", fmt.Errorf("error parsing interval %v", args[1])
	}

	timeVar := args[0]
	if len(args) == 3 {
		timeVar = parseTime(args[0], args[2])
	}

	return interval, timeVar, nil
}

func macroTimeGroup(query *models.Query, args []string) (string, error) {
	interval, timeVar, err := parseTimeGroup(query, args)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("FROM_UNIXTIME(FLOOR(TO_UNIXTIME(%s)/%v)*%v)", timeVar, interval.Seconds(), interval.Seconds()), nil
}

func macroUnixEpochGroup(query *models.Query, args []string) (string, error) {
	interval, timeVar, err := parseTimeGroup(query, args)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("FROM_UNIXTIME(FLOOR(%s/%v)*%v)", timeVar, interval.Seconds(), interval.Seconds()), nil
}

func macroParseTime(query *models.Query, args []string) (string, error) {
	if len(args) < 1 {
		return "", errors.WithMessagef(sqlds.ErrorBadArgumentCount, "expected at least one argument")
	}

	var (
		column     = args[0]
		timeFormat = timestampFormat
	)

	if len(args) == 2 {
		timeFormat = args[1]
	}

	return parseTime(column, timeFormat), nil
}

func macroTimeFilter(query *models.Query, args []string) (string, error) {
	if len(args) < 1 {
		return "", errors.WithMessagef(sqlds.ErrorBadArgumentCount, "expected at least one argument")
	}

	var (
		column     = args[0]
		timeFormat = ""
		from       = query.TimeRange.From.Format(goTimestampFormat)
		to         = query.TimeRange.To.Format(goTimestampFormat)
	)

	if len(args) > 1 {
		timeFormat = args[1]
	}
	timeVar := parseTime(column, timeFormat)

	return fmt.Sprintf("%s BETWEEN TIMESTAMP '%s' AND TIMESTAMP '%s'", timeVar, from, to), nil
}

func macroUnixEpochFilter(query *models.Query, args []string) (string, error) {
	if len(args) != 1 {
		return "", errors.WithMessagef(sqlds.ErrorBadArgumentCount, "expected one argument")
	}

	var (
		column = args[0]
		from   = query.TimeRange.From.UTC().Unix()
		to     = query.TimeRange.To.UTC().Unix()
	)

	return fmt.Sprintf("%s BETWEEN %d AND %d", column, from, to), nil
}

func macroTimeFrom(query *models.Query, args []string) (string, error) {
	return fmt.Sprintf("TIMESTAMP '%s'", query.TimeRange.From.UTC().Format(goTimestampFormat)), nil

}

func macroRawTimeFrom(query *models.Query, args []string) (string, error) {
	format := timestampFormat
	if len(args) == 1 && args[0] != "" {
		format = args[0]
	}
	timeLayout := toolbox.DateFormatToLayout(format)
	return query.TimeRange.From.UTC().Format(timeLayout), nil
}

func macroTimeTo(query *models.Query, args []string) (string, error) {
	return fmt.Sprintf("TIMESTAMP '%s'", query.TimeRange.To.UTC().Format(goTimestampFormat)), nil
}

func macroRawTimeTo(query *models.Query, args []string) (string, error) {
	format := timestampFormat
	if len(args) == 1 && args[0] != "" {
		format = args[0]
	}
	timeLayout := toolbox.DateFormatToLayout(format)
	return query.TimeRange.To.UTC().Format(timeLayout), nil
}

func macroDateFilter(query *models.Query, args []string) (string, error) {
	if len(args) != 1 {
		return "", errors.WithMessagef(sqlds.ErrorBadArgumentCount, "expected 1 argument, received %d", len(args))
	}

	var (
		column = args[0]
		from   = query.TimeRange.From.UTC().Format("2006-01-02")
		to     = query.TimeRange.To.UTC().Format("2006-01-02")
	)

	return fmt.Sprintf("%s BETWEEN date '%s' AND date '%s'", column, from, to), nil
}

func macroEdsId(query *models.Query, args []string) (string, error) {
	// args contains one element of empty string
	if len(args) > 1 || args[0] != "" {
		return "", errors.WithMessagef(sqlds.ErrorBadArgumentCount, "expected 0 arguments, received %d", len(args))
	}

	return query.EdsId, nil

	// return fmt.Sprintf("be078ded-503b-4baa-bdad-0f7697864340"), nil
}

type Macros map[string]MacroFunc

type MacroFunc func(*models.Query, []string) (string, error)

var macros = map[string]MacroFunc{
	"dateFilter":  macroDateFilter,
	"parseTime":   macroParseTime,
	"timeFilter":  macroTimeFilter,
	"rawTimeFrom": macroRawTimeFrom,
	"timeFrom":    macroTimeFrom,
	"timeGroup":   macroTimeGroup,
	"rawTimeTo":   macroRawTimeTo,
	"timeTo":      macroTimeTo,
	"edsId":       macroEdsId,
}

func (s *CtlDatasource) Macros() Macros {
	return macros
}
