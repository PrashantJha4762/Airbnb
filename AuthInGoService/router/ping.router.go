package router
import (
	"github.com/go-chi/chi/v5"
	"AuthInGoService/controllers"
)

type Router interface {
	Register( r *chi.Mux)
}

func SetUpRouter(userRouter Router) *chi.Mux {
	chirouter := chi.NewRouter()
	chirouter.Get("/ping",controllers.Pinghandler)

	userRouter.Register(chirouter)
	return chirouter
}
