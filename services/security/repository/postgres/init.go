package postgres

import (
	fmt "fmt"
	"microServiceBoilerplate/configs"
	instances "microServiceBoilerplate/services/security/instances"
	"microServiceBoilerplate/services/security/models"
	"time"

	"github.com/mreza0100/golog"
	postgres "gorm.io/driver/postgres"
	gorm "gorm.io/gorm"
	"gorm.io/gorm/logger"
	schema "gorm.io/gorm/schema"
)

func New(lgr *golog.Core) *instances.Repo_Postgres {
	var (
		db     *gorm.DB
		readQ  *read
		writeQ *write
	)

	{
		db = getConnection(lgr)
		autoSessionExpirator(db)
		lgr = lgr.With("In Repository->")
	}
	{
		readQ = &read{
			lgr: lgr.With("In Read->"),
			db:  db,
		}
		writeQ = &write{
			lgr: lgr.With("In Write->"),
			db:  db,
		}
	}

	return &instances.Repo_Postgres{
		Read:  readQ,
		Write: writeQ,
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
		if err := db.AutoMigrate(&models.Session{}, &models.Password{}); err != nil {
			lgr.Fatal(err)
		}
	}

	lgr.SuccessLog("Connected to DB")

	return db
}

func getConfigs() (driverConfigs gorm.Dialector, gormConfigs *gorm.Config) {
	DSN := fmt.Sprintf("host=localhost user=postgres dbname=postgres port=%v", configs.SecurityConfigs.DBPort)
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

func autoSessionExpirator(db *gorm.DB) {
	go func() {
		ticker := time.NewTicker(configs.Token_expire_auto_remove_duration)

		for {
			select {
			case <-ticker.C:
				now := time.Now().Add(time.Hour * 26).Format(time.RFC3339)

				tx := db.Exec(`
							DELETE
							FROM sessions
							WHERE expire_at < ?
						`, now)
				if tx.Error != nil {
					panic(tx.Error)
				}
			}
		}
	}()
}
