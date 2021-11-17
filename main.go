package main

import (
	config "widi443/currency-exchange/app/configs"
	route "widi443/currency-exchange/app/routes"

)

func main() {
	config.InitDB()
	e := route.New()
	e.Start(":9000")
}
