package handlers

import "net/http"

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Public API!"))
}

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Private Dashboard!"))
}
