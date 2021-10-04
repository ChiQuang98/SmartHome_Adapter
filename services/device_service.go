package services

import (
	"SmartHome_Adapter/core_libs/base"
	"SmartHome_Adapter/core_libs/models"
	"SmartHome_Adapter/errors"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang/glog"
)

func CreateDevice(deviceCreate *models.DeviceCreate) (int, []byte) {
	token := deviceCreate.Token
	//TODO: Create Thing API
	statusCreateThing, res := CreateThingMainflux(deviceCreate)
	if statusCreateThing != 201 {
		if statusCreateThing == 401 {
			return statusCreateThing, []byte(base.UNAUTHORIZED)
		}
		if statusCreateThing == 400 {
			return statusCreateThing, []byte(base.BAD_REQUEST)
		}
		return http.StatusInternalServerError, []byte(base.SERVER_ERROR)
	}
	//Todo: Get Thing Info
	thingID := string(res)
	statusCodeGetThing, thingMainflux, err := GetThingMainflux(thingID, token)
	if err != nil {
		if statusCodeGetThing == 401 {
			return statusCodeGetThing, []byte(base.UNAUTHORIZED)
		}
		if statusCodeGetThing == 400 {
			return statusCodeGetThing, []byte(base.BAD_REQUEST)
		}
		return http.StatusInternalServerError, []byte(base.SERVER_ERROR)
	}
	//TODO Create Channel
	statusCreateChannel, res := CreateChannel(deviceCreate)
	if statusCreateChannel != 201 {
		if statusCreateChannel == 401 {
			return statusCreateChannel, []byte(base.UNAUTHORIZED)
		}
		if statusCreateChannel == 400 {
			return statusCreateChannel, []byte(base.BAD_REQUEST)
		}
		return http.StatusInternalServerError, []byte(base.SERVER_ERROR)
	}
	channelID := string(res)
	responseCreateDevice := &models.ResponseCreateDevice{
		ThingID:   thingMainflux.ThingID,
		ThingKey:  thingMainflux.ThingKey,
		ChannelID: channelID,
	}
	//TODO Connect Thing To Channel
	statusConnect, res := ConnectThingToChannel(token, *responseCreateDevice)
	if statusConnect != 200 {
		if statusConnect == 401 {
			return statusConnect, []byte(base.UNAUTHORIZED)
		}
		if statusConnect == 400 {
			return statusConnect, []byte(base.BAD_REQUEST)
		}
		return http.StatusInternalServerError, []byte(base.SERVER_ERROR)
	}
	response, err1 := json.Marshal(&responseCreateDevice)
	if err1 != nil {
		glog.Error(err1.Error())
		return http.StatusInternalServerError, []byte(base.SERVER_ERROR)
	}
	return 200, response
}

func DeleteDevice(token, thingId, channelId string) error {
	op := errors.Op("services.DeleteDevice")

	if err := DeleteMainfluxThingById(token, thingId); err != nil {
		if errors.Is(errors.KindInvalidToken, err) {
			return errors.E(op, errors.KindInvalidToken, err)
		}
		return errors.E(op, err)
	}

	if err := DeleteMainfluxChannelById(token, channelId); err != nil {
		if errors.Is(errors.KindInvalidToken, err) {
			return errors.E(op, errors.KindInvalidToken, err)
		}
		return errors.E(op, err)
	}

	return nil
}
func DeviceAlarmOff(thingToken string, deviceOff *models.DeviceOffThing) (int, []byte) {
	statusSendMessage, res := SendMessageDeviceAlarmOff(thingToken, deviceOff)
	if statusSendMessage != 202 && res != nil {
		if statusSendMessage == 403 || statusSendMessage == 503 {
			return 401, []byte(base.UNAUTHORIZED)
		}
		if statusSendMessage == 400 {
			return 400, []byte(base.BAD_REQUEST)
		}
		return http.StatusInternalServerError, []byte(base.SERVER_ERROR)
	}
	responseSettingDevice := &models.SuccessResponse{
		Result: "success",
	}
	response, err1 := json.Marshal(&responseSettingDevice)
	if err1 != nil {
		glog.Error(err1.Error())
		return http.StatusInternalServerError, []byte(base.SERVER_ERROR)
	}
	return 200, response
}
func DeviceSettingApp(thingToken string, deviceSetting *models.DeviceSettingApp) (int, []byte) {
	statusSendMessage, res := SendMessage(thingToken, deviceSetting)
	if statusSendMessage != 202 && res != nil {
		if statusSendMessage == 403 || statusSendMessage == 503 {
			return 401, []byte(base.UNAUTHORIZED)
		}
		if statusSendMessage == 400 {
			return 400, []byte(base.BAD_REQUEST)
		}
		return http.StatusInternalServerError, []byte(base.SERVER_ERROR)
	}
	responseSettingDevice := &models.SuccessResponse{
		Result: "success",
	}
	response, err1 := json.Marshal(&responseSettingDevice)
	if err1 != nil {
		glog.Error(err1.Error())
		return http.StatusInternalServerError, []byte(base.SERVER_ERROR)
	}
	return 200, response
}

func SendMessageDeviceSettings(token, chanID string, setting ThingLog) error {
	op := errors.Op("services.SendMessageDeviceSettings")

	target := fmt.Sprintf(
		"%s/http/channels/%s/messages/SmartHomeThingLogs",
		urlMainflux,
		chanID,
	)

	body := map[string]interface{}{
		"mac_address":        setting.MacAddress,
		"home_away":          setting.HomeAway,
		"alarm_doorbell":     setting.AlarmDoorbell,
		"door_status":        setting.DoorStatus,
		"pin_volt":           setting.PinVolt,
		"ArmingDisarming":    setting.ArmingDisarming,
		"Boot":               setting.Boot,
		"RestoreFactory":     setting.RestoreFactory,
		"FirmwareVersion":    setting.FirmwareVersion,
		"OtaFirmwareTrigger": setting.OtaFirmwareTrigger,
		"OtaFirmwareReport":  setting.OtaFirmwareReport,
		"AlarmStatus":        setting.AlarmStatus,
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return errors.E(op, err)
	}

	req, err := http.NewRequest(http.MethodPost, target, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return errors.E(op, err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header = http.Header{
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{token},
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.E(op, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 503 {
		return errors.E(op, errors.KindUnauthorization, fmt.Errorf("request unauthorization"))
	}

	return nil
}
