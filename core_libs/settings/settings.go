package settings

import (
	"encoding/json"
	"io/ioutil"
)

type GlogConfigs struct {
	LogDir  string
	MaxSize uint64
	V       int
}

type Settings struct {
	GlogConfig     *GlogConfigs
	RestfulApiPort int
	RestfulApiHost string
	MainfluxInfo   *MainfluxInfo
	InfluxInfo     InfluxInfo
}

type Endpoints struct {
	Token    string
	Things   string
	Channels string
}

type MainfluxInfo struct {
	Address   string
	Port      int
	Endpoints *Endpoints
}

type InfluxInfo struct {
	Password string
	Username string
	Host     string
}

var settings Settings = Settings{}

func init() {
	content, err := ioutil.ReadFile("setting.json")
	if err != nil {
		panic(err)
	}
	settings = Settings{}
	jsonErr := json.Unmarshal(content, &settings)
	if jsonErr != nil {
		panic(jsonErr)
	}
}
func GetMainfluxInfo() *MainfluxInfo {
	return settings.MainfluxInfo
}

func GetInfluxInfo() InfluxInfo {
	return settings.InfluxInfo
}
func GetEndPoints() *Endpoints {
	return settings.MainfluxInfo.Endpoints
}
func GetGlogConfig() *GlogConfigs {
	return settings.GlogConfig
}
func GetRestfulApiPort() int {
	return settings.RestfulApiPort
}
func GetRestfulApiHost() string {
	return settings.RestfulApiHost
}
