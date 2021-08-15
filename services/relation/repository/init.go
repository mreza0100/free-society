package repository

import (
	fmt "fmt"

	"freeSociety/configs"
	"freeSociety/services/relation/instances"
	models "freeSociety/services/relation/models"

	"github.com/mreza0100/golog"
	postgres "gorm.io/driver/postgres"
	gorm "gorm.io/gorm"
	schema "gorm.io/gorm/schema"
)

func NewRepo(lgr *golog.Core) *instances.Repository {
	var (
		db *gorm.DB

		f_read  *followers_read
		f_write *followers_write

		l_read  *likes_read
		l_write *likes_write
	)

	{
		db = getConnection(lgr)
		lgr = lgr.With("In Repository->")
	}
	{
		f_read = &followers_read{
			lgr: lgr.With("In followers Read->"),
			db:  db,
		}
		f_write = &followers_write{
			lgr: lgr.With("In followers Write->"),
			db:  db,
		}
	}
	{
		l_read = &likes_read{
			lgr: lgr.With("In likes Read->"),
			db:  db,
		}
		l_write = &likes_write{
			lgr: lgr.With("In likes Write->"),
			db:  db,
		}
	}

	return &instances.Repository{
		Followers_read:  f_read,
		Followers_write: f_write,

		Likes_read:  l_read,
		Likes_write: l_write,
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
		if err := db.AutoMigrate(&models.Followers{}, &models.Like{}); err != nil {
			lgr.Fatal(err)
		}
	}

	lgr.SuccessLog("Connected to DB")

	return db
}
