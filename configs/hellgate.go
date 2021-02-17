package configs

import "strconv"

var HellgateConfigs serviceConfigs

func init() {
	const (
		port = 10000
	)
	HellgateConfigs = serviceConfigs{
		Addr:    "localhost:" + strconv.Itoa(port),
		Port:    port,
		Timeout: stdConnectionTimeout,
	}
}
