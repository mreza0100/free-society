package db

import (
	fmt "fmt"

	"microServiceBoilerplate/configs"
	models "microServiceBoilerplate/services/user/models"

	"github.com/mreza0100/golog"
	postgres "gorm.io/driver/postgres"
	gorm "gorm.io/gorm"
	schema "gorm.io/gorm/schema"
)

var (
	db *gorm.DB
)

func getDSN() string {
	var (
		host   = " host=localhost "
		user   = " user=postgres "
		dbname = " dbname=postgres "
		port   = " port=" + configs.UserConfigs.StrDBPort
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
	}

	return
}

func ConnectDB(lgr *golog.Core) {
	gormConfigs, driverConfigs := getConfigs()

	var err error
	db, err = gorm.Open(driverConfigs, gormConfigs)
	if err != nil {
		panic("failed to connect database")
	}

	lgr.SuccessLog("Connected to DB")

	migrations()
}

func migrations() {
	if err := db.AutoMigrate(&models.User{}); err != nil {
		fmt.Println(err)
		fmt.Println("\n\n\n\n\n\n\n\n ")
		panic("db migration failed")
	}
}
