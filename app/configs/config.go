package config

import (
	model "widi443/currency-exchange/app/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBConnection *gorm.DB

func InitDB() {
	dsn := "host=localhost user=root password=root dbname=db_currency-exchange port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	var err error
	DBConnection, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	initMigrate()
}

func initMigrate() {
	DBConnection.AutoMigrate(
		&model.User{},
		&model.Wallet{},
		&model.Transaction{},
	)

}
