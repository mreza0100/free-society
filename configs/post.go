package configs

var PostConfigs *serviceConfigs

func init() {
	const (
		port   = 9091
		dbPort = 5434
	)
	PostConfigs = &serviceConfigs{
		Addr:      "localhost:" + str(port),
		StrPort:   str(port),
		Port:      port,
		DBPort:    dbPort,
		StrDBPort: str(dbPort),
		Timeout:   stdConnectionTimeout,
	}
}
