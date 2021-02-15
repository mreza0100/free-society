package db

import (
	fmt "fmt"
	"microServiceBoilerplate/configs"
	models "microServiceBoilerplate/services/security/models"

	"github.com/go-redis/redis"
	"github.com/mreza0100/golog"
	postgres "gorm.io/driver/postgres"
	gorm "gorm.io/gorm"
	"gorm.io/gorm/logger"
	schema "gorm.io/gorm/schema"
)

var (
	psDB    *gorm.DB
	redisDB *redis.Client
)

func getDSN() string {
	var (
		host   = " host=localhost "
		user   = " user=postgres "
		dbname = " dbname=postgres "
		port   = " port=" + configs.SecurityConfigs.StrDBPort
	)
	return host + user + dbname + port
}

func getConfigs() (gormConfigs *gorm.Config, driverConfigs gorm.Dialector) {
	driverConfigs = postgres.New(postgres.Config{
		DSN: getDSN(),
	})

	gormConfigs = &gorm.Config{
		NamingStrategy:         schema.NamingStrategy{},
		SkipDefaultTransaction: true,
		// PrepareStmt:            false,
		Logger: logger.Default,
	}

	return
}

func ConnectPS(lgr *golog.Core) {
	gormConfigs, driverConfigs := getConfigs()

	var err error
	psDB, err = gorm.Open(driverConfigs, gormConfigs)
	if err != nil {
		panic("failed to connect database")
	}
	lgr.SuccessLog("Connected to ps db")

	if err := psDB.AutoMigrate(&models.Password{}, &models.Session{}); err != nil {
		fmt.Println(err)
		fmt.Println("\n\n\n\n\n\n\n\n ")
		panic("db migration failed")
	}
}

func ConnecRedis(lgr *golog.Core) {
	redisDB = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("localhost:%v", configs.SecurityConfigs.RedisPort),
		Password: "",
		DB:       0,

		OnConnect: func(c *redis.Conn) error {
			lgr.GreenLog("redis is connected successfuly")
			return nil
		},
	})
}
