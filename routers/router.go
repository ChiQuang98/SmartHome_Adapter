package routers

import (
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router:=mux.NewRouter()
	router = SetDeviceRouter(router)
	return router
}