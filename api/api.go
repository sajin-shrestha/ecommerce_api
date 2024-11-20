package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sajin-shrestha/ecommerce_api/database"
	"github.com/sajin-shrestha/ecommerce_api/middlewares"
	"github.com/sajin-shrestha/ecommerce_api/routes"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{addr: addr}
}

func (s *APIServer) Run() error {
	err := database.DBConnection()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err.Error())
		return err
	}

	router := mux.NewRouter()
	router.Use(middlewares.CORS)
	routes.InitRoutes(router)

	const underline = "\033[4m"
	const reset = "\033[0m"

	localURL := "http://localhost"
	apiURL := localURL + s.addr

	log.Printf("Starting server on: %s%s%s", underline, apiURL, reset)
	return http.ListenAndServe(s.addr, router)
}
