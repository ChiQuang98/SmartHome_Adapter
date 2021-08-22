package routers

import (
	"SmartHome_Adapter/controllers"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetDeviceRouter(router *mux.Router) *mux.Router {
	router.Handle("/smarthome/v1/create-device",
		negroni.New(
			negroni.HandlerFunc(controllers.CreateDevice),
		)).Methods("POST")
	router.Handle("/smarthome/v1/testHello",
		negroni.New(
			negroni.HandlerFunc(controllers.HelloWorld),
		)).Methods("GET")
	return router
}