package services

import (
	"SmartHome_Adapter/core_libs/models"
	"SmartHome_Adapter/core_libs/settings"
	"fmt"
	"github.com/golang/glog"
	"io/ioutil"
	"net/http"
)

func ConnectThingToChannel(token string,responseCreateDevice models.ResponseCreateDevice)(int,[]byte)  {
	client := &http.Client{}
	urlConnectThingToChannel := fmt.Sprintf(urlMainflux+settings.GetEndPoints().Channels+"/%s"+
		settings.GetEndPoints().Things+"/%s",responseCreateDevice.ChannelID,responseCreateDevice.ThingID)
	req, err := http.NewRequest(http.MethodPut,urlConnectThingToChannel , nil)
	if err !=nil{
		glog.Error("Fail to request to API Connect Thing To Channel")
		return http.StatusInternalServerError,[]byte(err.Error())
	}
	req.Header = http.Header{
		"Authorization": []string{token},
	}
	res, err := client.Do(req)
	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)
	if res.StatusCode != 200{
		return res.StatusCode,data
	}
	return 200,nil
}
