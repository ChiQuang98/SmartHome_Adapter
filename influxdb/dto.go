package influxdb

import (
	"SmartHome_Adapter/services"
	"encoding/json"
	"fmt"
)

func toSmartHomeAppLog(v []interface{}) (services.AppLog, error) {
	fmt.Println(v)
	if len(v) != 7 {
		return services.AppLog{}, fmt.Errorf("invalid value length: %d", len(v))
	}

	alarmDelay, err := v[1].(json.Number).Int64()
	if err != nil {
		return services.AppLog{}, fmt.Errorf("invalid value of alarm_delay: %T", v[1])
	}

	alarmDuaration, err := v[2].(json.Number).Int64()
	if err != nil {
		return services.AppLog{}, fmt.Errorf("invalid value of alarm_duaration: %T", v[2])
	}

	armDelay, err := v[3].(json.Number).Int64()
	if err != nil {
		return services.AppLog{}, fmt.Errorf("invalid value of arm_delay: %T", v[3])
	}

	deviceVolume, err := v[4].(json.Number).Int64()
	if err != nil {
		return services.AppLog{}, fmt.Errorf("invalid value of device_volume: %T", v[4])
	}

	macAddr, ok := v[5].(string)
	if !ok {
		return services.AppLog{}, fmt.Errorf("invalid value of device_volume: %T", v[5])
	}

	passwordSetting, ok := v[6].(string)
	if !ok {
		return services.AppLog{}, fmt.Errorf("invalid value of device_volume: %T", v[6])
	}

	return services.AppLog{
		MacAddr:         macAddr,
		PasswordSetting: passwordSetting,
		DeviceVolume:    deviceVolume,
		ArmDelay:        armDelay,
		AlarmDelay:      alarmDelay,
		AlarmDuaration:  alarmDuaration,
	}, nil
}
