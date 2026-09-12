package main

import (
	"AuthInGoService/app"
	config "AuthInGoService/config/env"
)

func main() {
	config.Load()
	cfg := app.NewConfig()
	app := app.NewApplication(cfg)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
