package controllers

import (
	"SmartHome_Adapter/bind"
	"encoding/json"
	"net/http"
)

func Query(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	var req struct {
		Table string `json:"table"`
		Limit int    `json:"limit"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responseJson(w, http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	if req.Limit == 0 {
		responseJson(w, http.StatusBadRequest, map[string]string{
			"error": "limit must be greater than zero",
		})
		return
	}

	if req.Table != "SmartHomeAppLogs" && req.Table != "SmartHomeThingLogs" {
		responseJson(w, http.StatusBadRequest, map[string]string{
			"error": "table must be SmartHomeAppLogs or SmartHomeThingLogs",
		})
		return
	}

	logs, err := bind.QueryReadModel().Query(req.Table, req.Limit)
	if err != nil {
		responseJson(w, http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}

	responseJson(w, http.StatusOK, logs)
}
