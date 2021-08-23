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

func CreateChannel(deviceCreate *models.DeviceCreate) (int, []byte) {
	token := *deviceCreate.Token
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
