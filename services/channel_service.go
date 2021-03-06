package services

import (
	"SmartHome_Adapter/core_libs/models"
	"SmartHome_Adapter/core_libs/settings"
	"SmartHome_Adapter/errors"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/golang/glog"
)

func SendMessage(token string, deviceSetting *models.DeviceSettingApp) (int, []byte) {
	///http/channels/79f4a262-fe70-460c-8575-3969f0047135/messages
	client := &http.Client{}
	urlSendMessage := urlMainflux + "/http/channels/" + deviceSetting.ChannelID + "/messages/SmartHomeAppLogs"
	body := &models.DeviceSettingAppBody{
		MacAddress:      deviceSetting.MacAddress,
		DeviceVolume:    deviceSetting.DeviceVolume,
		PasswordSetting: deviceSetting.PasswordSetting,
		ArmDelay:        deviceSetting.ArmDelay,
		AlarmDelay:      deviceSetting.AlarmDelay,
		AlarmDuaration:  deviceSetting.AlarmDuaration,
		AlarmStatus:     deviceSetting.AlarmStatus,
	}
	jsonDeviceSetting, _ := json.Marshal(body)
	req, err := http.NewRequest(http.MethodPost, urlSendMessage, bytes.NewBuffer(jsonDeviceSetting))
	if err != nil {
		glog.Error("Fail request api send Message Mainflux", err)
		return http.StatusInternalServerError, []byte(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header = http.Header{
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{token},
	}
	resChannel, err := client.Do(req)
	defer resChannel.Body.Close()
	if resChannel.StatusCode != 202 {
		data, _ := ioutil.ReadAll(resChannel.Body)
		return resChannel.StatusCode, data
	}
	return resChannel.StatusCode, nil
}
func SendMessageDeviceAlarmOff(token string, deviceOff *models.DeviceOffThing) (int, []byte) {
	///http/channels/79f4a262-fe70-460c-8575-3969f0047135/messages
	client := &http.Client{}
	urlSendMessage := urlMainflux + "/http/channels/" + deviceOff.ChannelID + "/messages/SmartHomeThingLogs"
	body := &models.DeviceOffThingBody{
		MacAddress:         deviceOff.MacAddress,
		HomeAway:           deviceOff.HomeAway,
		AlarmDoorbell:      deviceOff.AlarmDoorbell,
		PinVolt:            deviceOff.PinVolt,
		ArmingDisarming:    deviceOff.ArmingDisarming,
		DoorStatus:         deviceOff.DoorStatus,
		Boot:               deviceOff.Boot,
		RestoreFactory:     deviceOff.RestoreFactory,
		FirmwareVersion:    deviceOff.FirmwareVersion,
		OtaFirmwareTrigger: deviceOff.OtaFirmwareTrigger,
		OtaFirmwareReport:  deviceOff.OtaFirmwareReport,
		AlarmStatus:        deviceOff.AlarmStatus,
	}
	jsonDeviceOff, _ := json.Marshal(body)
	req, err := http.NewRequest(http.MethodPost, urlSendMessage, bytes.NewBuffer(jsonDeviceOff))
	if err != nil {
		glog.Error("Fail request api send Message Mainflux", err)
		return http.StatusInternalServerError, []byte(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header = http.Header{
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{token},
	}
	resChannel, err := client.Do(req)
	defer resChannel.Body.Close()
	if resChannel.StatusCode != 202 {
		data, _ := ioutil.ReadAll(resChannel.Body)
		return resChannel.StatusCode, data
	}
	return resChannel.StatusCode, nil
}
func CreateChannel(deviceCreate *models.DeviceCreate) (int, []byte) {
	token := deviceCreate.Token
	client := &http.Client{}
	metadata := &models.Metadata{
		Type: "SmartHome",
	}
	body := &models.ThingRequest{
		Name:     deviceCreate.MacAddress,
		Metadata: metadata,
	}
	jsonChannel, _ := json.Marshal(body)
	urlCreateChannels := urlMainflux + settings.GetEndPoints().Channels
	req, err := http.NewRequest(http.MethodPost, urlCreateChannels, bytes.NewBuffer(jsonChannel))
	if err != nil {
		glog.Error("Fail request api Add Thing mainflux", err)
		return http.StatusInternalServerError, []byte(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header = http.Header{
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{token},
	}
	resChannelCreate, err := client.Do(req)
	defer resChannelCreate.Body.Close()
	if resChannelCreate.StatusCode != 201 {
		data, _ := ioutil.ReadAll(resChannelCreate.Body)
		return resChannelCreate.StatusCode, data
	}
	location := resChannelCreate.Header.Get("Location")
	arrLoc := strings.Split(location, "/")
	channelID := arrLoc[2]
	return resChannelCreate.StatusCode, []byte(channelID)
}

func DeleteMainfluxChannelById(token, channelId string) error {
	op := errors.Op("services.DeleteMainfluxChannelById")

	req, err := http.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("%s%s/%s", urlMainflux, settings.GetEndPoints().Channels, channelId),
		nil,
	)

	if err != nil {
		return errors.E(op, err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.E(op, err)
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return errors.E(op, errors.KindInvalidToken, err)
	}

	return nil
}

func ChannelExists(token, channelId string) (bool, error) {
	op := errors.Op("services.ThingExists")

	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s%s/%s", urlMainflux, settings.GetEndPoints().Channels, channelId),
		nil,
	)

	if err != nil {
		return false, errors.E(op, err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, errors.E(op, err)
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return false, errors.E(op, errors.KindInvalidToken, fmt.Errorf("invalid token"))
	}

	if resp.StatusCode == http.StatusNotFound {
		var body struct {
			Error string `json:"error"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
			return false, errors.E(op, err)
		}

		if body.Error != "non-existent entity" {
			return false, errors.E(op, err)
		}

		return false, nil
	}

	if resp.StatusCode == http.StatusOK {
		return true, nil
	}

	return false, errors.E(op, fmt.Errorf("unexpected error: %d", resp.StatusCode))
}
