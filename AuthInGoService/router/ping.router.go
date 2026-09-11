package router

import (
	"AuthInGoService/controllers"
	"AuthInGoService/middleware"

	"github.com/go-chi/chi/v5"
)

type Router interface {
	Register( r *chi.Mux)
}

func SetUpRouter(userRouter Router) *chi.Mux {
	chirouter := chi.NewRouter()
	chirouter.Use(middleware.RateLimiter); 
	chirouter.Get("/ping",controllers.Pinghandler)

	userRouter.Register(chirouter)
	return chirouter
}
