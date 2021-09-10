package configs

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type ServiceConfigs struct {
	Name string

	Addr              string
	LogPath           string
	ConnectionTimeout time.Duration

	Service_port  int `json:"service_port"`
	Postgres_port int `json:"postgres_port"`
	Mongo_port    int `json:"mongo_port"`
	Redis_port    int `json:"redis_port"`
}

func ConfigToMap(path string) (rawConfigs map[string]interface{}) {
	rawConfigs = make(map[string]interface{})

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&rawConfigs)
	if err != nil {
		panic(err)
	}

	return rawConfigs
}

func (sc *ServiceConfigs) SetConfigFile() {
	if sc.Name == "" {
		panic("Service name is empty")
	}

	configPath := os.Getenv(fmt.Sprintf("%s_CONFIG_FILE_PATH", strings.ToUpper(sc.Name)))
	rawConfigs := ConfigToMap(configPath)

	{
		sc.Service_port = int(rawConfigs["service_port"].(float64))
		sc.Postgres_port = int(rawConfigs["postgres_port"].(float64))
		sc.Redis_port = int(rawConfigs["redis_port"].(float64))
		sc.Mongo_port = int(rawConfigs["mongo_port"].(float64))

		sc.Addr = fmt.Sprintf("localhost:%v", sc.Service_port)
	}
	{
		rawTimeout, found := rawConfigs["timeout"]
		if found {
			sc.ConnectionTimeout = time.Duration(rawTimeout.(float64)) * time.Second
		} else {
			sc.ConnectionTimeout = time.Duration(time.Second * 5)
		}
	}
	{
		rawLogPath, found := rawConfigs["log_path"]
		if found {
			sc.LogPath = fmt.Sprintf("%v%v", ROOT, rawLogPath)
		} else {
			sc.LogPath = fmt.Sprintf("%s/logs/%s.log", ROOT, sc.Name)
		}
	}
}
