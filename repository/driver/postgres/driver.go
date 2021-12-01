package postgres

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ConfigDB struct {
	DB_Username string
	DB_Password string
	DB_Host     string
	DB_Database string
	DB_Port     string
}

func (config *ConfigDB) InitialDb() *gorm.DB {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Shanghai",
		config.DB_Host,
		config.DB_Username,
		config.DB_Password,
		config.DB_Database,
		config.DB_Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
