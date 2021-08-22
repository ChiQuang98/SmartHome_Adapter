package controllers

import (
	"net/http"
)

func CreateDevice(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)  {
	
}
func HelloWorld(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)  {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("HelloWolrd"))
}
