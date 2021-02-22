package configs

var FeedConfigs *serviceConfigs

func init() {
	const (
		port      = 9093
		dbPort    = 6379
		redisPort = dbPort
	)
	FeedConfigs = &serviceConfigs{
		Addr:      "localhost:" + str(port),
		Port:      port,
		DBPort:    dbPort,
		RedisPort: redisPort,
		Timeout:   stdConnectionTimeout,
	}
}
