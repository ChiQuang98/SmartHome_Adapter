package influxdb

import (
	"SmartHome_Adapter/services"
	"encoding/json"
	"fmt"
)

func toSmartHomeAppLog(v []interface{}) (services.AppLog, error) {
	var err error

	if len(v) != 8 {
		return services.AppLog{}, fmt.Errorf("invalid value length: %d", len(v))
	}

	alarmDelay := int64(5)
	if v[1] != nil {
		alarmDelay, err = v[1].(json.Number).Int64()
		if err != nil {
			return services.AppLog{}, fmt.Errorf("invalid value of alarm_delay: %T", v[1])
		}
	}

	alarmDuaration := int64(5)
	if v[2] != nil {
		alarmDuaration, err = v[2].(json.Number).Int64()
		if err != nil {
			return services.AppLog{}, fmt.Errorf("invalid value of alarm_duaration: %T", v[2])
		}
	}

	alarmStatus := int64(0)
	if v[3] != nil {
		alarmStatus, err = v[3].(json.Number).Int64()
		if err != nil {
			return services.AppLog{}, fmt.Errorf("invalid value of arm_delay: %T", v[3])
		}
	}

	armDelay := int64(5)
	if v[4] != nil {
		armDelay, err = v[4].(json.Number).Int64()
		if err != nil {
			return services.AppLog{}, fmt.Errorf("invalid value of arm_delay: %T", v[4])
		}
	}

	deviceVolume := int64(1)
	if v[5] != nil {
		deviceVolume, err = v[5].(json.Number).Int64()
		if err != nil {
			return services.AppLog{}, fmt.Errorf("invalid value of device_volume: %T", v[5])
		}
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
