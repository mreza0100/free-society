package db

import (
	fmt "fmt"

	models "microServiceBoilerplate/services/user/models"

	postgres "gorm.io/driver/postgres"
	gorm "gorm.io/gorm"
	schema "gorm.io/gorm/schema"
)

var (
	DB *gorm.DB
)

const (
	host   = " host=localhost "
	user   = " user=postgres "
	dbname = " dbname=postgres "
	port   = " port=5433 "
)

var (
	dsn = host + user + dbname + port
)

func init() {
	DB = connectDB()
	fmt.Println("user service: ", "✅db is connected✅")
}

func getConfigs() (gormConfigs *gorm.Config, driverConfigs gorm.Dialector) {
	driverConfigs = postgres.New(postgres.Config{
		DSN: dsn,
	})

	gormConfigs = &gorm.Config{
		NamingStrategy:         schema.NamingStrategy{},
		SkipDefaultTransaction: true,
	}

	return
}

func connectDB() *gorm.DB {
	gormConfigs, driverConfigs := getConfigs()

	db, err := gorm.Open(driverConfigs, gormConfigs)
	if err != nil {
		panic("failed to connect database")
	}

	migrations(db)

	return db
}

func migrations(db *gorm.DB) {
	if err := db.AutoMigrate(&models.User{}); err != nil {
		fmt.Println(err)
		fmt.Println("\n\n\n\n\n\n\n\n ")
		panic("db migration failed")
	}
}
