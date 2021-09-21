package routers

import (
	"SmartHome_Adapter/controllers"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetDeviceRouter(router *mux.Router) *mux.Router {
	router.Handle("/aiot-smarthome/v1/app/create-device",
		negroni.New(
			negroni.HandlerFunc(controllers.CreateDevice),
		)).Methods("POST")

	router.Handle("/aiot-smarthome/v1/app/delete-device",
		negroni.New(
			negroni.HandlerFunc(controllers.DeleteDevice),
		)).Methods("POST")
	router.Handle("/aiot-smarthome/v1/app/device-setting",
		negroni.New(
			negroni.HandlerFunc(controllers.DeviceSettingApp),
		)).Methods("POST")
	router.Handle("/aiot-smarthome/v1/app/device-alarmoff",
		negroni.New(
			negroni.HandlerFunc(controllers.DeviceSettingApp),
		)).Methods("POST")
	router.Handle("/aiot-smarthome/v1/thing/device-alarmoff",
		negroni.New(
			negroni.HandlerFunc(controllers.DeviceAlarmOff),
		)).Methods("POST")
	router.Handle("/smarthome/v1/testHello",
		negroni.New(
			negroni.HandlerFunc(controllers.HelloWorld),
		)).Methods("GET")
	return router
}
