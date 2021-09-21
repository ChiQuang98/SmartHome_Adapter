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
		ResponseErr := models.ErrorResponse{Error: base.BAD_REQUEST}
		response, _ := json.Marshal(&ResponseErr)
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
		} else {
			ResponseErr := models.ErrorResponse{Error: string(res)}
			response, _ := json.Marshal(&ResponseErr)
			w.Write(response)
		}
	}
}

func DeleteDevice(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	req := models.DeviceDelete{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responseJson(w, http.StatusBadRequest, map[string]string{
			"error": base.BAD_REQUEST,
		})
		return
	}

	if err := req.Validate(); err != nil {
		responseJson(w, http.StatusBadRequest, map[string]string{
			"error": base.BAD_REQUEST,
		})
		return
	}

	token := r.Header.Get("Authorization")

	thingExists, err := services.ThingExists(token, req.ThingID)
	if err != nil {
		if errors.Is(errors.KindInvalidToken, err) {
			responseJson(w, http.StatusUnauthorized, map[string]interface{}{
				"error": base.UNAUTHORIZED,
			})
			return
		}

		responseJson(w, http.StatusNotFound, map[string]interface{}{
			"error": base.SERVER_ERROR,
		})
		return
	}

	if !thingExists {
		responseJson(w, http.StatusNotFound, map[string]interface{}{
			"error": "Thing not found",
		})
		return
	}

	channelExists, err := services.ChannelExists(token, req.ChannelID)
	if err != nil {
		if errors.Is(errors.KindInvalidToken, err) {
			responseJson(w, http.StatusUnauthorized, map[string]interface{}{
				"error": base.UNAUTHORIZED,
			})
			return
		}

		responseJson(w, http.StatusInternalServerError, map[string]interface{}{
			"error": base.SERVER_ERROR,
		})
		return
	}

	if !channelExists {
		responseJson(w, http.StatusNotFound, map[string]interface{}{
			"error": "channel not found",
		})
		return
	}

	if err := services.DeleteDevice(token, req.ThingID, req.ChannelID); err != nil {
		if errors.Is(errors.KindInvalidToken, err) {
			responseJson(w, http.StatusUnauthorized, map[string]interface{}{
				"error": base.UNAUTHORIZED,
			})
			return
		}

		responseJson(w, http.StatusInternalServerError, map[string]interface{}{
			"error": base.SERVER_ERROR,
		})
		return
	}

	responseJson(w, http.StatusOK, map[string]interface{}{
		"result": "success",
	})
}
func DeviceAlarmOff(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	deviceAlarmOff := new(models.DeviceOffThing)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(deviceAlarmOff)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil || deviceAlarmOff.MacAddress == "" || deviceAlarmOff.ChannelID == "" {
		w.WriteHeader(http.StatusBadRequest)
		ResponseErr := models.ErrorResponse{Error: base.BAD_REQUEST}
		response, _ := json.Marshal(&ResponseErr)
		w.Write(response)
	} else {
		token := r.Header.Get("Authorization")
		status, res := services.DeviceAlarmOff(token,deviceAlarmOff)
		w.WriteHeader(status)
		if status == http.StatusOK {
			w.Write(res)
		} else {
			ResponseErr := models.ErrorResponse{Error: string(res)}
			response, _ := json.Marshal(&ResponseErr)
			w.Write(response)
		}
	}
}

func DeviceSettingApp(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	deviceSettingApp := new(models.DeviceSettingApp)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(deviceSettingApp)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil || deviceSettingApp.ChannelID == "" || deviceSettingApp.MacAddress == "" {
		w.WriteHeader(http.StatusBadRequest)
		ResponseErr := models.ErrorResponse{Error: base.BAD_REQUEST}
		response, _ := json.Marshal(&ResponseErr)
		w.Write(response)
	} else {
		token := r.Header.Get("Authorization")
		status, res := services.DeviceSettingApp(token,deviceSettingApp)
		w.WriteHeader(status)
		if status == http.StatusOK {
			w.Write(res)
		} else {
			ResponseErr := models.ErrorResponse{Error: string(res)}
			response, _ := json.Marshal(&ResponseErr)
			w.Write(response)
		}
	}
}
func HelloWorld(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("HelloWolrd"))
}
