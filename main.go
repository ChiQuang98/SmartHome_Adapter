package main

import (
	"SmartHome_Adapter/core_libs/settings"
	"SmartHome_Adapter/routers"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/codegangsta/negroni"
	"github.com/golang/glog"
	"github.com/rs/cors"
)

func init() {
	//glog
	//create logs folder
	os.Mkdir("./logs", 0777)
	flag.Lookup("stderrthreshold").Value.Set("[INFO|WARN|FATAL]")
	flag.Lookup("logtostderr").Value.Set("false")
	flag.Lookup("alsologtostderr").Value.Set("true")
	flag.Lookup("log_dir").Value.Set("./logs")
	glog.MaxSize = 1024 * 1024 * settings.GetGlogConfig().MaxSize
	flag.Lookup("v").Value.Set(fmt.Sprintf("%d", settings.GetGlogConfig().V))
	flag.Parse()
}
func main() {
	flag.Parse()
	routerApi := routers.InitRoutes()
	nApi := negroni.Classic()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{"DELETE", "PUT", "GET", "HEAD", "OPTIONS", "POST"},
	})
	nApi.Use(c)
	nApi.UseHandler(routerApi)
	host := fmt.Sprint(settings.GetRestfulApiHost()+":", strconv.Itoa(settings.GetRestfulApiPort()))
	http.ListenAndServe(host, nApi)
	glog.Info("Service Started")
}
