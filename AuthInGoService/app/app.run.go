package app

import (
	config "AuthInGoService/config/env"
	dbconfig "AuthInGoService/config/db"
	"AuthInGoService/controllers"
	db "AuthInGoService/db/repositories"
	"AuthInGoService/router"
	"AuthInGoService/services"
	"fmt"
	"net/http"
	"time"
	dbb "AuthInGoService/db/repositories"
)

type Config struct {
	Addr string
}

type Application struct {
	Config Config
	Store  *db.Storage
}

func NewConfig() Config {
	PORT := config.GetString("PORT", ":8080")
	return Config{
		Addr: PORT,
	}
}

func NewApplication(cfg Config) *Application {
	return &Application{
		Config: cfg,
		Store:  db.NewStorage(),
	}
}

func (app *Application) Run() error {
	db,err:=dbconfig.SetUpDB()
	if(err!=nil){
		fmt.Println("Something went wrong")
		return err;
	}
	ur:=dbb.NewUserRepository(db)
	userService := services.NewUserService(ur)
	userController := controllers.NewUserController(userService)
	userRouter := router.NewUserRouter(*userController)
	server := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      router.SetUpRouter(userRouter),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("The server is running on", app.Config.Addr)
	return server.ListenAndServe()
}
