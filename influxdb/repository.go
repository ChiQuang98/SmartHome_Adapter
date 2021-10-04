package influxdb

import (
	"SmartHome_Adapter/errors"
	"SmartHome_Adapter/services"
	"encoding/json"
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
	influxdb1 "github.com/influxdata/influxdb1-client/v2"
)

type Repository struct {
	c         influxdb1.Client
	db        string
	precision string
}

func NewRepository(c influxdb1.Client) Repository {
	return Repository{
		c:         c,
		db:        "mainflux",
		precision: "",
	}
}

func (r Repository) GetLatestSmartHomeAppLog(macAddr string) (services.AppLog, error) {
	op := errors.Op("influxdb.Repository.GetLatestSmartHomeAppLog")

	influxQuery := influxdb1.Query{
		Command: `
			SELECT "alarm_delay", "alarm_duaration", "alarm_status", "arm_delay", "device_volume", "mac_address", "password_setting" 
				FROM "SmartHomeAppLogs"
				WHERE "mac_address" = $mac_addr
				ORDER BY "time" DESC
				LIMIT 1
		`,
		Database: r.db,
		Parameters: map[string]interface{}{
			"mac_addr": macAddr,
		},
		Precision: "rfc3339",
	}

	response, err := r.c.Query(influxQuery)
	if err != nil {
		return services.AppLog{}, errors.E(op, err)
	}

	values := [][]interface{}{}
	for _, result := range response.Results {
		for _, se := range result.Series {
			values = append(values, se.Values...)
		}
	}

	if len(values) != 1 {
		return services.AppLog{}, errors.E(
			op,
			errors.KindNotFound,
			fmt.Errorf("no log for this mac: %s", macAddr),
		)
	}

	ret, err := toSmartHomeAppLog(values[0])
	if err != nil {
		return services.AppLog{}, errors.E(op, err)
	}

	return ret, nil
}

func (r Repository) Query(table string, limit int) ([]map[string]interface{}, error) {
	op := errors.Op("influxdb.Repository.Query")

	influxQuery := influxdb1.Query{
		Command: fmt.Sprintf(`
			SELECT * FROM "%s"
				ORDER BY "time" DESC
				LIMIT %d
		`, table, limit),
		Database:  r.db,
		Precision: "rfc3339",
	}

	response, err := r.c.Query(influxQuery)
	if err != nil {
		return nil, errors.E(op, err)
	}

	spew.Dump(response)

	ret := []map[string]interface{}{}

	for _, result := range response.Results {
		for _, se := range result.Series {
			for _, value := range se.Values {
				row := map[string]interface{}{}
				for i := range se.Columns {
					if se.Columns[i] == "time" {
						nsec, err := value[i].(json.Number).Int64()
						if err != nil {
							continue
						}

						ts := time.Unix(0, nsec)

						row[se.Columns[i]] = ts.Format("02/01/06 15:04:05 -0700")
						continue
					}
					row[se.Columns[i]] = value[i]
				}
				ret = append(ret, row)
			}
		}
	}

	return ret, nil
}
