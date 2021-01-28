package configs

import "strconv"

type hellgateConfigs struct {
	Addr    string
	StrPort string
	Port    int
	timeout int
}

var HellgateConfigs hellgateConfigs

func init() {
	const port = 10000
	HellgateConfigs = hellgateConfigs{
		Addr:    "localhost:" + strconv.Itoa(port),
		StrPort: strconv.Itoa(port),
		Port:    port,
		timeout: int(connectionTimeout),
	}
}
