package repository

import (
	fmt "fmt"

	"microServiceBoilerplate/configs"
	"microServiceBoilerplate/services/relation/instances"
	models "microServiceBoilerplate/services/relation/models"

	"github.com/mreza0100/golog"
	postgres "gorm.io/driver/postgres"
	gorm "gorm.io/gorm"
	schema "gorm.io/gorm/schema"
)

func NewRepo(lgr *golog.Core) *instances.Repository {
	var (
		db     *gorm.DB
		readQ  *read
		writeQ *write
	)

	{
		db = getConnection(lgr)
		lgr = lgr.With("In Repository ->")
	}
	{
		readQ = &read{
			lgr: lgr.With("In Read ->"),
			db:  db,
		}
		writeQ = &write{
			lgr: lgr.With("In Write ->"),
			db:  db,
		}
		readQ.write = writeQ
		writeQ.read = readQ
	}

	return &instances.Repository{
		Read:  readQ,
		Write: writeQ,
	}
}

func getConfigs() (driverConfigs gorm.Dialector, gormConfigs *gorm.Config) {
	DSN := fmt.Sprintf("host=localhost user=postgres dbname=postgres port=%v", configs.RelationConfigs.DBPort)
	driverConfigs = postgres.New(postgres.Config{
		DSN: DSN,
	})

	gormConfigs = &gorm.Config{
		NamingStrategy:         schema.NamingStrategy{},
		SkipDefaultTransaction: true,
		// PrepareStmt:            false,
	}

	return
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
		if err := db.AutoMigrate(&models.Followers{}); err != nil {
			lgr.Fatal(err)
		}
	}

	lgr.SuccessLog("Connected to DB")

	return db
}
