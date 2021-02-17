package configs

var SecurityConfigs *serviceConfigs

func init() {
	const (
		port      = 9094
		dbPort    = 5436
		redisPort = 6380
	)
	SecurityConfigs = &serviceConfigs{
		Addr:    "localhost:" + str(port),
		Timeout: stdConnectionTimeout,

		Port: port,

		DBPort: dbPort,

		RedisPort: redisPort,
	}
}
