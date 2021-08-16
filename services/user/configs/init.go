package configs

import (
	"freeSociety/configs"
)

const serviceName = "user"

var Configs *configs.ServiceConfigs

func init() {
	Configs = new(configs.ServiceConfigs)
	Configs.Name = serviceName
	Configs.SetConfigFile()
}
