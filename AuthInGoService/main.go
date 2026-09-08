package main

import (

	"AuthInGoService/app"
	config "AuthInGoService/config/env"
	dbconfig "AuthInGoService/config/db"
)

func main() {
	config.Load()
	cfg := app.NewConfig()
	app := app.NewApplication(cfg)
	dbconfig.SetUpDB()
	app.Run()
}
