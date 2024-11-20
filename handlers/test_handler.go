package handlers

import (
	"net/http"

	"github.com/sajin-shrestha/ecommerce_api/response"
)

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"message": "Welcome to the Public Api!",
	}

	response.SUCCESS(w, http.StatusOK, data)

}

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"message": "Welcome to the Private Dashboard!",
	}

	response.SUCCESS(w, http.StatusOK, data)
}
