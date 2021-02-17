package configs

var FeedConfigs *serviceConfigs

func init() {
	const (
		port   = 9093
		dbPort = 6379
	)
	FeedConfigs = &serviceConfigs{
		Addr:    "localhost:" + str(port),
		Port:    port,
		DBPort:  dbPort,
		Timeout: stdConnectionTimeout,
	}
}
