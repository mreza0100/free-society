package configs

var UserConfigs *serviceConfigs

func init() {
	const (
		port   = 9090
		dbPort = 5433
	)
	UserConfigs = &serviceConfigs{
		Addr:      "localhost:" + str(port),
		StrPort:   str(port),
		Port:      port,
		DBPort:    dbPort,
		StrDBPort: str(dbPort),
		Timeout:   stdConnectionTimeout,
	}
}
