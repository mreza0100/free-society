package configs

import "strconv"

type userConfigs struct {
	Addr    string
	StrPort string
	Port    int
	DBPort  int
	timeout int
}

var UserConfigs userConfigs

func init() {
	const port = 9090
	UserConfigs = userConfigs{
		Addr:    "localhost:" + strconv.Itoa(port),
		StrPort: strconv.Itoa(port),
		Port:    9090,
		DBPort:  5433,
		timeout: int(connectionTimeout),
	}
}
