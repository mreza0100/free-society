package configs

import "strconv"

var UserConfigs serviceConfigs

func init() {
	const (
		port   = 9090
		dbPort = 5433
	)
	UserConfigs = serviceConfigs{
		Addr:    "localhost:" + strconv.Itoa(port),
		StrPort: strconv.Itoa(port),
		Port:    port,
		DBPort:  dbPort,
		Timeout: stdConnectionTimeout,
	}
}
