package controllers

import (
	"SmartHome_Adapter/core_libs/models"
	"SmartHome_Adapter/services"
	"encoding/json"
	"net/http"
)

func CreateDevice(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)  {
	deviceCreate:=new(models.DeviceCreate)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(deviceCreate)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err!=nil || deviceCreate.Token == nil || *deviceCreate.Token == "" || deviceCreate.MacAddress == nil || *deviceCreate.MacAddress == ""{
		w.WriteHeader(http.StatusBadRequest)
	} else {
		status, res := services.CreateDevice(deviceCreate)
		w.WriteHeader(status)
		w.Write(res)
	}

}
func HelloWorld(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)  {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("HelloWolrd"))
}
