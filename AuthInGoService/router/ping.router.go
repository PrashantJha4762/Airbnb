package router
import (
	"github.com/go-chi/chi/v5"
	"AuthInGoService/controllers"
)

func SetUpRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/ping",controllers.Pinghandler)
	return router
}
