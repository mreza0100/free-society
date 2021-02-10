package configs

import "time"

type serviceConfigs struct {
	Addr    string
	StrPort string
	Port    int
	DBPort  int
	Timeout time.Duration
}

const (
	LogPath = "./logs/all.log"
)

var (
	// standard connection timeout for services
	stdConnectionTimeout = time.Duration(2 * time.Second)
)
