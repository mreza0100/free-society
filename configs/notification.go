package configs

var NotificationConfigs *serviceConfigs

func init() {
	const (
		port   = 9095
		dbPort = 5437
	)
	NotificationConfigs = &serviceConfigs{
		Addr:    "localhost:" + str(port),
		Port:    port,
		DBPort:  dbPort,
		Timeout: stdConnectionTimeout,
	}
}
