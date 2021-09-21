package influxdb

import (
	"SmartHome_Adapter/core_libs/settings"
	"SmartHome_Adapter/errors"
	"time"

	influxdb1 "github.com/influxdata/influxdb1-client/v2"
)

func Connect(opts settings.InfluxInfo) (influxdb1.Client, error) {
	op := errors.Op("InfluxDB Connect")

	client, err := influxdb1.NewHTTPClient(influxdb1.HTTPConfig{
		Addr:     opts.Host,
		Username: opts.Username,
		Password: opts.Password,
	})

	if err != nil {
		return nil, errors.E(op, err)
	}

	if _, _, err := client.Ping(2 * time.Second); err != nil {
		return nil, errors.E(op, err)
	}

	return client, nil
}
