package services

import (
	"SmartHome_Adapter/core_libs/models"
	"SmartHome_Adapter/core_libs/settings"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	"io/ioutil"
	"net/http"
	"strings"
)

var addressMainflux = settings.GetMainfluxInfo().Address
var portMainflux= settings.GetMainfluxInfo().Port
var urlMainflux =fmt.Sprintf("http://%s:%d", addressMainflux, portMainflux)
func GetThingMainflux(thingID string,token string)(int,models.ThingMainflux,[]byte){
	client := &http.Client{}
	urlGetThingInfo := fmt.Sprintf(urlMainflux+settings.GetEndPoints().Things+"/%s",thingID)
	req, err := http.NewRequest(http.MethodGet,urlGetThingInfo , nil)
	if err !=nil{
		glog.Error("Fail to request to API Get Thing Info")
		return http.StatusInternalServerError,models.ThingMainflux{},[]byte(err.Error())
	}
	req.Header = http.Header{
		"Authorization": []string{token},
	}
	res, err := client.Do(req)
	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)
	if res.StatusCode != 200{
		return res.StatusCode,models.ThingMainflux{},data
	}
	thingResponse:=models.ThingResponse{}
	err=json.Unmarshal(data,&thingResponse)
	if err!=nil{
		glog.Error(err)
		return http.StatusInternalServerError,models.ThingMainflux{},[]byte(err.Error())
	}
	thingMainflux:=models.ThingMainflux{
		ThingID:  *thingResponse.ID,
		ThingKey: *thingResponse.Key,
	}
	return res.StatusCode,thingMainflux,nil
}
func CreateThingMainflux(deviceCreate *models.DeviceCreate) (int,[]byte){
	token:=deviceCreate.Token
	metadata:=&models.Metadata{
		Type: "SmartHome",
	}
	bodyThing:=&models.ThingRequest{
		Name:     deviceCreate.MacAddress,
		Metadata: metadata,
	}
	client := &http.Client{}
	jsonThing, _ := json.Marshal(bodyThing)
	req, err := http.NewRequest(http.MethodPost,urlMainflux+settings.GetEndPoints().Things , bytes.NewBuffer(jsonThing))
	if err !=nil{
		glog.Error("Fail request api Add Thing mainflux",err)
		return http.StatusInternalServerError,[]byte(err.Error())
	}
	req.Header.Set("Content-Type","application/json")
	req.Header = http.Header{
		"Content-Type":[]string{"application/json"},
		"Authorization": []string{*token},
	}
	resThingCreate, err := client.Do(req)
	defer resThingCreate.Body.Close()
	if resThingCreate.StatusCode != 201{
		data, _ := ioutil.ReadAll(resThingCreate.Body)
		return resThingCreate.StatusCode,data
	}
	location:=resThingCreate.Header.Get("Location")
	arrLoc:=strings.Split(location,"/")
	thingID:=arrLoc[2]
	return resThingCreate.StatusCode,[]byte(thingID)
}
