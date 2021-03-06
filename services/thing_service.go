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

var addressMainflux = settings.GetMainfluxInfo().Address
var portMainflux = settings.GetMainfluxInfo().Port
var urlMainflux = fmt.Sprintf("http://%s:%d", addressMainflux, portMainflux)

func GetThingMainflux(thingID string, token string) (int, models.ThingMainflux, []byte) {
	client := &http.Client{}
	urlGetThingInfo := fmt.Sprintf(urlMainflux+settings.GetEndPoints().Things+"/%s", thingID)
	req, err := http.NewRequest(http.MethodGet, urlGetThingInfo, nil)
	if err != nil {
		glog.Error("Fail to request to API Get Thing Info")
		return http.StatusInternalServerError, models.ThingMainflux{}, []byte(err.Error())
	}
	req.Header = http.Header{
		"Authorization": []string{token},
	}
	res, err := client.Do(req)
	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)
	if res.StatusCode != 200 {
		return res.StatusCode, models.ThingMainflux{}, data
	}
	thingResponse := models.ThingResponse{}
	err = json.Unmarshal(data, &thingResponse)
	if err != nil {
		glog.Error(err)
		return http.StatusInternalServerError, models.ThingMainflux{}, []byte(err.Error())
	}
	thingMainflux := models.ThingMainflux{
		ThingID:  *thingResponse.ID,
		ThingKey: *thingResponse.Key,
	}
	return res.StatusCode, thingMainflux, nil
}

func CreateThingMainflux(deviceCreate *models.DeviceCreate) (int, []byte) {
	token := deviceCreate.Token
	metadata := &models.Metadata{
		Type: "SmartHome",
	}
	bodyThing := &models.ThingRequest{
		Name:     deviceCreate.MacAddress,
		Metadata: metadata,
	}
	client := &http.Client{}
	jsonThing, _ := json.Marshal(bodyThing)
	req, err := http.NewRequest(http.MethodPost, urlMainflux+settings.GetEndPoints().Things, bytes.NewBuffer(jsonThing))
	if err != nil {
		glog.Error("Fail request api Add Thing mainflux", err)
		return http.StatusInternalServerError, []byte(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header = http.Header{
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{token},
	}
	resThingCreate, err := client.Do(req)
	defer resThingCreate.Body.Close()
	if resThingCreate.StatusCode != 201 {
		data, _ := ioutil.ReadAll(resThingCreate.Body)
		return resThingCreate.StatusCode, data
	}
	location := resThingCreate.Header.Get("Location")
	arrLoc := strings.Split(location, "/")
	thingID := arrLoc[2]
	return resThingCreate.StatusCode, []byte(thingID)
}

func DeleteMainfluxThingById(token, thingId string) error {
	op := errors.Op("services.DeleteThingMainflux")

	req, err := http.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("%s%s/%s", urlMainflux, settings.GetEndPoints().Things, thingId),
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
		return errors.E(op, errors.KindInvalidToken, fmt.Errorf("invalid token"))
	}

	return nil
}

func ThingExists(token, thingId string) (bool, error) {
	op := errors.Op("services.ThingExists")

	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s%s/%s", urlMainflux, settings.GetEndPoints().Things, thingId),
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
