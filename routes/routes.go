package routes

import (
	"github.com/gorilla/mux"
	"github.com/sajin-shrestha/ecommerce_api/handlers"
)

func InitRoutes(router *mux.Router) {
	// Public routes
	public := router.PathPrefix("/").Subrouter()
	public.HandleFunc("/", handlers.WelcomeHandler).Methods("GET")

	// Private routes (require authentication middleware)
	private := router.PathPrefix("/api").Subrouter()
	// private.Use(middlewares.AuthMiddleware)
	private.HandleFunc("/dashboard", handlers.DashboardHandler).Methods("GET")
}
