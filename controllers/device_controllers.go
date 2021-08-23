package controllers

import (
	"SmartHome_Adapter/core_libs/models"
	"SmartHome_Adapter/services"
	"encoding/json"
	"net/http"
)

func CreateDevice(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)  {
	macAddress:=new(models.MacAddress)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(macAddress)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err!=nil || macAddress.MacAddress == nil || *macAddress.MacAddress == ""{
		w.WriteHeader(http.StatusBadRequest)
	} else {
		token:=r.Header.Get("Authorization")
		deviceCreate:=new(models.DeviceCreate)
		deviceCreate.MacAddress = macAddress.MacAddress
		deviceCreate.Token = token
		status, res := services.CreateDevice(deviceCreate)
		w.WriteHeader(status)
		if(status == http.StatusOK){
			w.Write(res)
		}
	}
}
func HelloWorld(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)  {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("HelloWolrd"))
}
