package main

import (
	config "widi443/currency-exchange/configs"
	route "widi443/currency-exchange/routes"
)

func main() {
	config.InitDB()
	e := route.New()
	e.Start(":9000")
}
