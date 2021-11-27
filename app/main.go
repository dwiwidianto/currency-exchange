package main

import (
	"currency-exchange/app/routes"
	"currency-exchange/repository/driver/postgres"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"

	_userUsecase "currency-exchange/business/users"
	_userController "currency-exchange/controllers/users"
	_userRepo "currency-exchange/repository/databases/users"
	// _currencyUsecase "currency-exchange/business/currency"
	// _currencyController "currency-exchange/controllers/currency"
)

func init() {
	viper.SetConfigFile("app/config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func dbMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&_userRepo.User{},
	)
	if err != nil {
		panic(err)
	}

}

func main() {
	configDb := postgres.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.dbname`),
	}

	db := configDb.InitialDb()
	dbMigrate(db)

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	// Penghubung Layer
	e := echo.New()
	userRepository := _userRepo.NewPosgresUserRepository(db)
	userUsecase := _userUsecase.NewUseCase(userRepository, timeoutContext)
	userCtrl := _userController.NewUserController(e, userUsecase)

	routesInit := routes.RouteControllerList{
		UserController: *userCtrl,
	}

	routesInit.RouteRegiester(e)
	log.Fatal(e.Start(viper.GetString("server.address")))
}
