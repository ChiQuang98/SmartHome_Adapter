package controllers

import (
	"SmartHome_Adapter/core_libs/base"
	"SmartHome_Adapter/core_libs/models"
	"SmartHome_Adapter/errors"
	"SmartHome_Adapter/services"
	"encoding/json"
	"net/http"
)

func CreateDevice(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	macAddress := new(models.MacAddress)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(macAddress)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil || macAddress.MacAddress == nil || *macAddress.MacAddress == "" {
		w.WriteHeader(http.StatusBadRequest)
		ResponseErr:=models.ErrorResponse{Error: base.BAD_REQUEST}
		response,_ := json.Marshal(&ResponseErr)
		w.Write(response)
	} else {
		token := r.Header.Get("Authorization")
		deviceCreate := new(models.DeviceCreate)
		deviceCreate.MacAddress = macAddress.MacAddress
		deviceCreate.Token = token
		status, res := services.CreateDevice(deviceCreate)
		w.WriteHeader(status)
		if status == http.StatusOK {
			w.Write(res)
		} else{
			ResponseErr:=models.ErrorResponse{Error: string(res)}
			response,_ := json.Marshal(&ResponseErr)
			w.Write(response)
		}
	}
}

func DeleteDevice(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	op := errors.Op("controllers.DeleteDevice")

	req := models.DeviceDelete{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responseJson(w, http.StatusBadRequest, map[string]string{
			"error": errors.E(op, err).Error(),
		})
		return
	}

	token := r.Header.Get("Authorization")

	if err := services.DeleteDevice(token, req.ThingID, req.ChannelID); err != nil {
		if errors.Is(errors.KindInvalidToken, err) {
			responseJson(w, http.StatusUnauthorized, map[string]interface{}{
				"error":   errors.E(op, err).Error(),
				"success": false,
			})
			return
		}

		responseJson(w, http.StatusInternalServerError, map[string]interface{}{
			"error":   errors.E(op, err).Error(),
			"success": false,
		})
		return
	}

	responseJson(w, http.StatusOK, map[string]interface{}{
		"success": true,
	})
}

func HelloWorld(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("HelloWolrd"))
}
