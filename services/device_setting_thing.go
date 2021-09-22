package services

import "SmartHome_Adapter/errors"

type AppLog struct {
	MacAddr         string
	PasswordSetting string
	DeviceVolume    int64
	ArmDelay        int64
	AlarmDelay      int64
	AlarmDuaration  int64
	AlarmStatus     int64
}

type ThingLog struct {
	MacAddr            string
	FirmwareVersion    string
	HomeAway           int64
	AlarmDoorbell      int64
	PinVolt            int64
	ArmingDisarming    int64
	Boot               int64
	RestoreFactory     int64
	OtaFirmwareTrigger int64
	OtaFirmwareReport  int64
	AlarmStatus        int64
}

type LasAppLogReadModel interface {
	GetLatestSmartHomeAppLog(string) (AppLog, error)
}

func DeviceSettingThing(readmodel LasAppLogReadModel, token string, chanid string, l ThingLog) (AppLog, error) {
	op := errors.Op("services.DeviceSettingThing")

	if err := SendMessageDeviceSettings(token, chanid, l); err != nil {
		if errors.Is(errors.KindUnauthorization, err) {
			return AppLog{}, errors.E(op, errors.KindUnauthorization, err)
		}
		return AppLog{}, errors.E(op, err)
	}

	latest, err := readmodel.GetLatestSmartHomeAppLog(l.MacAddr)
	if err != nil {
		if errors.Is(errors.KindNotFound, err) {
			return AppLog{}, errors.E(op, errors.KindNotFound, err)
		}
		return AppLog{}, errors.E(op, err)
	}

	return latest, nil
}
