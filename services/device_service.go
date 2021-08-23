package services

import (
	"SmartHome_Adapter/core_libs/models"
	"encoding/json"
	"github.com/golang/glog"
	"net/http"
)



func CreateDevice(deviceCreate *models.DeviceCreate) (int,[]byte) {
	token:=*deviceCreate.Token
	//TODO: Create Thing API
	statusCreateThing,res:=CreateThingMainflux(deviceCreate)
	if statusCreateThing!=201{
		return statusCreateThing,res
	}
	//Todo: Get Thing Info
	thingID:=string(res)
	statusCodeGetThing,thingMainflux,err:=GetThingMainflux(thingID,token)
	if err!=nil{
		return statusCodeGetThing,err
	}
	//TODO Create Channel
	statusCreateChannel,res:=CreateChannel(deviceCreate)
	if statusCreateChannel!=201{
		return statusCreateChannel,res
	}
	channelID:=string(res)
	responseCreateDevice:=&models.ResponseCreateDevice{
		ThingID:   thingMainflux.ThingID,
		ThingKey:  thingMainflux.ThingKey,
		ChannelID: channelID,
	}
	//TODO Connect Thing To Channel
	statusConnect,res := ConnectThingToChannel(token,*responseCreateDevice)
	if statusConnect!=200{
		return statusConnect,res
	}
	response,err1 :=json.Marshal(&responseCreateDevice)
	if err1!=nil{
		glog.Error(err1.Error())
		return http.StatusInternalServerError,[]byte(err1.Error())
	}
	return 200,response;
}
