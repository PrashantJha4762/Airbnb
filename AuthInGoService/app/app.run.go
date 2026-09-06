package app

import (
	config "AuthInGoService/config/env"
	"AuthInGoService/router"
	"fmt"
	"net/http"
	"time"
)

type Config struct {
	Addr string
}

type Application struct {
	Config Config
}

func NewConfig() Config {
	PORT:=config.GetString("PORT",":8080")
	return Config{
		Addr: PORT,
	}
}

func NewApplication(cfg Config) *Application {
	return &Application{
		Config: cfg,
	}
}

func (app *Application) Run() error {
	server := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      router.SetUpRouter(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("The server is running on", app.Config.Addr)
	return server.ListenAndServe()
}
