package bind

import (
	"SmartHome_Adapter/influxdb"
	"SmartHome_Adapter/services"
)

var (
	influxRepository influxdb.Repository
)

func RegisterInfluxRepository(r influxdb.Repository) {
	influxRepository = r
}

func LastestSmartHomeAppLogReadModel() services.LasAppLogReadModel {
	return influxRepository
}

func QueryReadModel() services.QueryReadmodel {
	return influxRepository
}
