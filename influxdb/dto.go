package influxdb

import (
	"SmartHome_Adapter/services"
	"encoding/json"
	"fmt"
)

func toSmartHomeAppLog(v []interface{}) (services.AppLog, error) {
	if len(v) != 8 {
		return services.AppLog{}, fmt.Errorf("invalid value length: %d", len(v))
	}

	if nilExists(v) {
		return services.AppLog{}, fmt.Errorf("a field is nil: %v", v)
	}

	alarmDelay, err := v[1].(json.Number).Int64()
	if err != nil {
		return services.AppLog{}, fmt.Errorf("invalid value of alarm_delay: %T", v[1])
	}

	alarmDuaration, err := v[2].(json.Number).Int64()
	if err != nil {
		return services.AppLog{}, fmt.Errorf("invalid value of alarm_duaration: %T", v[2])
	}

	alarmStatus, err := v[3].(json.Number).Int64()
	if err != nil {
		return services.AppLog{}, fmt.Errorf("invalid value of arm_delay: %T", v[3])
	}

	armDelay, err := v[4].(json.Number).Int64()
	if err != nil {
		return services.AppLog{}, fmt.Errorf("invalid value of arm_delay: %T", v[4])
	}

	deviceVolume, err := v[5].(json.Number).Int64()
	if err != nil {
		return services.AppLog{}, fmt.Errorf("invalid value of device_volume: %T", v[5])
	}

	macAddr, ok := v[6].(string)
	if !ok {
		return services.AppLog{}, fmt.Errorf("invalid value of device_volume: %T", v[6])
	}

	passwordSetting, ok := v[7].(string)
	if !ok {
		return services.AppLog{}, fmt.Errorf("invalid value of device_volume: %T", v[7])
	}

	return services.AppLog{
		MacAddress:      macAddr,
		PasswordSetting: passwordSetting,
		DeviceVolume:    deviceVolume,
		ArmDelay:        armDelay,
		AlarmDelay:      alarmDelay,
		AlarmDuaration:  alarmDuaration,
		AlarmStatus:     alarmStatus,
	}, nil
}

func nilExists(values []interface{}) bool {
	for _, v := range values {
		if v == nil {
			return true
		}
	}
	return false
}
