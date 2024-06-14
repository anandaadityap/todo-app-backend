package config

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDB() {
	dbUser := viper.GetString("DB_USERNAME")
	dbPassword := viper.GetString("DB_PASSWORD")
	dbHost := viper.GetString("DB_HOST")
	dbPort := viper.GetString("DB_PORT")
	dbName := viper.GetString("DB_NAME")

	dsn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v", dbUser, dbPassword, dbHost, dbPort, dbName)
	log.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	//errMigrate := db.AutoMigrate(&entity.Todo{})
	//if errMigrate != nil {
	//	log.Println(errMigrate, "Error migration")
	//}

	log.Println("Database Connected...")

	DB = db
}
