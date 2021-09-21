package models

import "fmt"

type DeviceCreate struct {
	Token      string  `json:"token"`
	MacAddress *string `json:"mac_address"`
}
type MacAddress struct {
	MacAddress *string `json:"mac_address"`
}

type DeviceDelete struct {
	ThingID   string `json:"thing_id"`
	ChannelID string `json:"channel_id"`
}

func (r DeviceDelete) Validate() error {
	if r.ThingID == "" {
		return fmt.Errorf("ThingID empty")
	}

	if r.ChannelID == "" {
		return fmt.Errorf("ChannelID empty")
	}

	return nil
}

type DeviceSettingThing struct {
	ChannelID          string `json:"channel_id"`
	MacAddr            string `json:"mac_address"`
	HomeAway           int64  `json:"home_away"`
	AlarmDoorbell      int64  `json:"alarm_doorbell"`
	PinVolt            int64  `json:"pin_volt"`
	ArmingDisarming    int64  `json:"ArmingDisarming"`
	Boot               int64  `json:"Boot"`
	RestoreFactory     int64  `json:"RestoreFactory"`
	FirmwareVersion    int64  `json:"FirmwareVersion"`
	OtaFirmwareTrigger int64  `json:"OtaFirmwareTrigger"`
	OtaFirmwareReport  int64  `json:"OtaFirmwareReport"`
	AlarmStatus        int64  `json:"AlarmStatus"`
}

func (r DeviceSettingThing) Validate() error {
	if r.MacAddr == "" {
		return fmt.Errorf("MacAddr empty")
	}

	if r.ChannelID == "" {
		return fmt.Errorf("ChannelID empty")
	}

	return nil
}

type DeviceSettingApp struct {
	ChannelID       string `json:"channel_id"`
	MacAddress      string `json:"mac_address"`
	DeviceVolume    int    `json:"device_volume"`
	PasswordSetting string `json:"password_setting"`
	ArmDelay        int    `json:"arm_delay"`
	AlarmDelay      int    `json:"alarm_delay"`
	AlarmDuaration  int    `json:"alarm_duaration"`
	AlarmStatus     int    `json:"alarm_status"`
}
type DeviceOffThing struct {
	ChannelID          string `json:"channel_id"`
	MacAddress         string `json:"mac_address"`
	HomeAway           int    `json:"home_away"`
	AlarmDoorbell      int    `json:"alarm_doorbell"`
	PinVolt            int    `json:"pin_volt"`
	ArmingDisarming    int    `json:"arming_disarming"`
	DoorStatus         int    `json:"door_status"`
	Boot               int    `json:"boot"`
	RestoreFactory     int    `json:"restore_factory"`
	FirmwareVersion    string `json:"firmware_version"`
	OtaFirmwareTrigger int    `json:"ota_firmware_trigger"`
	OtaFirmwareReport  int    `json:"ota_firmware_report"`
	AlarmStatus        int    `json:"alarm_status"`
}
type DeviceOffThingBody struct {
	MacAddress         string `json:"mac_address"`
	HomeAway           int    `json:"home_away"`
	AlarmDoorbell      int    `json:"alarm_doorbell"`
	PinVolt            int    `json:"pin_volt"`
	ArmingDisarming    int    `json:"arming_disarming"`
	DoorStatus         int    `json:"door_status"`
	Boot               int    `json:"boot"`
	RestoreFactory     int    `json:"restore_factory"`
	FirmwareVersion    string `json:"firmware_version"`
	OtaFirmwareTrigger int    `json:"ota_firmware_trigger"`
	OtaFirmwareReport  int    `json:"ota_firmware_report"`
	AlarmStatus        int    `json:"alarm_status"`
}
type DeviceSettingAppBody struct {
	MacAddress      string `json:"mac_address"`
	DeviceVolume    int    `json:"device_volume"`
	PasswordSetting string `json:"password_setting"`
	ArmDelay        int    `json:"arm_delay"`
	AlarmDelay      int    `json:"alarm_delay"`
	AlarmDuaration  int    `json:"alarm_duaration"`
	AlarmStatus     int    `json:"alarm_status"`
}

type ThingRequest struct {
	Key      *string   `json:"key"`
	Name     *string   `json:"name"`
	Metadata *Metadata `json:"metadata"`
}
type ThingResponse struct {
	ID       *string   `json:"id"`
	Key      *string   `json:"key"`
	Name     *string   `json:"name"`
	Metadata *Metadata `json:"metadata"`
}
type ThingMainflux struct {
	ThingID  string `json:"thing_id"`
	ThingKey string `json:"thing_key"`
}
type ResponseCreateDevice struct {
	ThingID   string `json:"thing_id"`
	ThingKey  string `json:"thing_key"`
	ChannelID string `json:"channel_id"`
}
type Metadata struct {
	Type string `json:"type"`
}
type ChannelRequest struct {
	Name     string    `json:"name"`
	Metadata *Metadata `json:"metadata"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Result string `json:"result"`
}
