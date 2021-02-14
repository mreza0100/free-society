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

		StrPort: str(port),
		Port:    port,

		DBPort:    dbPort,
		StrDBPort: str(dbPort),

		RedisPort:    redisPort,
		StrRedisPort: str(redisPort),
	}
}
