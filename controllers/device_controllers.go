package controllers

import (
	"SmartHome_Adapter/core_libs/models"
	"SmartHome_Adapter/errors"
	"SmartHome_Adapter/services"
	"encoding/json"
	"net/http"
)

func CreateDevice(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	deviceCreate := new(models.DeviceCreate)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(deviceCreate)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil || deviceCreate.Token == nil || *deviceCreate.Token == "" || deviceCreate.MacAddress == nil || *deviceCreate.MacAddress == "" {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		status, res := services.CreateDevice(deviceCreate)
		w.WriteHeader(status)
		w.Write(res)
	}
}

func DeleteDevice(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	op := errors.Op("controllers.DeleteDevice")

	req := models.DeviceDelete{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responseJson(w, http.StatusInternalServerError, map[string]string{
			"error": errors.E(op, err).Error(),
		})
		return
	}

	if err := services.DeleteDevice(req); err != nil {
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
