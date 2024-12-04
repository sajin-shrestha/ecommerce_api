package handlers

import (
	"net/http"

	"github.com/sajin-shrestha/ecommerce_api/response"
)

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"user": []string{"user1", "user2", "user3"},
		"page": 1,
		"next": nil,
		"prev": nil,
	}

	response.SUCCESS(w, http.StatusOK, data)

}

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"message": "Welcome to the Private Dashboard!",
	}

	response.SUCCESS(w, http.StatusOK, data)
}

func ErrorTestHandler(w http.ResponseWriter, r *http.Request) {
	errDetail := "This is a simulated error for testing purposes."
	response.ERROR(w, http.StatusBadRequest, errDetail)
}
