package configs

import "strconv"

var PostConfigs serviceConfigs

func init() {
	const (
		port   = 9092
		dbPort = 5434
	)
	PostConfigs = serviceConfigs{
		Addr:    "localhost:" + strconv.Itoa(port),
		StrPort: strconv.Itoa(port),
		Port:    port,
		DBPort:  dbPort,
		Timeout: stdConnectionTimeout,
	}
}
