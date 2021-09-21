package influxdb

import (
	"SmartHome_Adapter/errors"
	"SmartHome_Adapter/services"
	"fmt"

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
			SELECT "alarm_delay", "alarm_duaration", "arm_delay", "device_volume", "mac_address", "password_setting" 
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
