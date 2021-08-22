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
	GlogConfig *GlogConfigs
	RestfulApiPort       int
	RestfulApiHost       string
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
func GetGlogConfig() *GlogConfigs {
	return settings.GlogConfig
}
func GetRestfulApiPort() int {
	return settings.RestfulApiPort
}
func GetRestfulApiHost() string {
	return settings.RestfulApiHost
}
