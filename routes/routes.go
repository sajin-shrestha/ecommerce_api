package routes

import (
	"github.com/gorilla/mux"
	"github.com/sajin-shrestha/ecommerce_api/handlers"
)

// public routes
func initPublicRoutes(router *mux.Router) {
	public := router.PathPrefix("/").Subrouter()
	public.HandleFunc("/", handlers.WelcomeHandler).Methods("GET")
	public.HandleFunc("/error/", handlers.ErrorTestHandler).Methods("GET")
	public.HandleFunc("/users/", handlers.GetAllUsers).Methods("GET")
}

// private routes
func initPrivateRoutes(router *mux.Router) {
	private := router.PathPrefix("/api").Subrouter()
	// Uncomment the middleware when ready
	// private.Use(middlewares.AuthMiddleware)
	private.HandleFunc("/dashboard/", handlers.DashboardHandler).Methods("GET")
}

func InitRoutes(router *mux.Router) {
	initPublicRoutes(router)  // Initializes public routes
	initPrivateRoutes(router) // Initializes private routes
}
