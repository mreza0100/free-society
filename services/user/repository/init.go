package repository

import (
	"fmt"
	"microServiceBoilerplate/configs"
	"microServiceBoilerplate/services/user/instanses"
	models "microServiceBoilerplate/services/user/models"

	"github.com/mreza0100/golog"
	postgres "gorm.io/driver/postgres"
	gorm "gorm.io/gorm"
	"gorm.io/gorm/logger"
	schema "gorm.io/gorm/schema"
)

func NewRepo(lgr *golog.Core) *instanses.Repository {
	var (
		db        *gorm.DB
		readCQRS  *read
		writeCQRS *write
	)

	{
		db = getConnection(lgr)
		lgr = lgr.With("In Repository -> ")
	}
	{
		readCQRS = &read{
			lgr: lgr.With("In Read -> "),
			db:  db,
		}
		writeCQRS = &write{
			lgr: lgr.With("In Read -> "),
			db:  db,
		}
	}

	return &instanses.Repository{
		Read:  readCQRS,
		Write: writeCQRS,
	}
}

func getConnection(lgr *golog.Core) *gorm.DB {
	var (
		err error
		db  *gorm.DB
	)

	{
		db, err = gorm.Open(getConfigs())
		if err != nil {
			lgr.Fatal(err)
		}
	}
	{
		if err := db.AutoMigrate(&models.User{}); err != nil {
			lgr.Fatal(err)
		}
	}

	lgr.SuccessLog("Connected to DB")

	return db
}

func getConfigs() (driverConfigs gorm.Dialector, gormConfigs *gorm.Config) {
	DSN := fmt.Sprintf("host=localhost user=postgres dbname=postgres port=%v", configs.UserConfigs.DBPort)
	driverConfigs = postgres.New(postgres.Config{
		DSN: DSN,
	})

	gormConfigs = &gorm.Config{
		NamingStrategy:         schema.NamingStrategy{},
		SkipDefaultTransaction: true,
		Logger:                 logger.Default,
		// PrepareStmt:            false,
	}

	return
}
