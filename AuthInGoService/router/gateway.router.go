package router

import (
	config "AuthInGoService/config/env"
	"AuthInGoService/middleware"
	"AuthInGoService/utils"

	"github.com/go-chi/chi/v5"
)

// RegisterGatewayRoutes exposes backend services only through JWT-authenticated
// gateway paths. Configure the upstream services with HOTEL_SERVICE_URL and
// BOOKING_SERVICE_URL, for example http://127.0.0.1:3001.
func RegisterGatewayRoutes(r *chi.Mux) {
	registerProxy(r, "/hotels", config.GetString("HOTEL_SERVICE_URL", ""))
	registerProxy(r, "/bookings", config.GetString("BOOKING_SERVICE_URL", ""))
}

func registerProxy(r *chi.Mux, pathPrefix, targetURL string) {
	handler := middleware.JwtAuth(utils.ReverseProxy(targetURL, pathPrefix))
	r.Handle(pathPrefix, handler)
	r.Handle(pathPrefix+"/*", handler)
}
