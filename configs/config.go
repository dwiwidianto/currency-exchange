package config

import (
	model "widi443/currency-exchange/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBConnection *gorm.DB

func InitDB() {
	dsn := "host=172.21.0.2 user=root password=root dbname=db_currency-exchange port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	var err error
	DBConnection, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	initMigrate()
}

func initMigrate() {
	DBConnection.AutoMigrate(&model.User{})

}
