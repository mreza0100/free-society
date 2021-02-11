package db

import (
	fmt "fmt"

	"microServiceBoilerplate/configs"
	models "microServiceBoilerplate/services/relation/models"

	postgres "gorm.io/driver/postgres"
	gorm "gorm.io/gorm"
	schema "gorm.io/gorm/schema"
)

var (
	DB *gorm.DB
)

func getDSN() string {
	var (
		host   = " host=localhost "
		user   = " user=postgres "
		dbname = " dbname=postgres "
		port   = " port=" + configs.RelationConfigs.StrDBPort + " "
	)

	return host + user + dbname + port
}

func init() {
	DB = connectDB()
	fmt.Println("Relation service: ", "✅ db is connected")
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
	if err := db.AutoMigrate(&models.Followers{}); err != nil {
		fmt.Println(err)
		fmt.Println("\n\n\n\n\n\n\n\n ")
		panic("db migration failed")
	}
}
