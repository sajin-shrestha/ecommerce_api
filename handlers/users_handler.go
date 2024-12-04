package handlers

import (
	"net/http"

	"github.com/sajin-shrestha/ecommerce_api/modals"
	"github.com/sajin-shrestha/ecommerce_api/response"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	allUsers := []modals.User{
		{
			CommonModal: modals.CommonModal{
				Status:   true,
				Priority: 1,
			},
			UserDto: modals.UserDto{
				FirstName: "Sajin",
				LastName:  "Shrestha",
				Age:       22,
				Address:   "Kathmandu",
				Email:     "shresthasajin59@gmail.com",
				Password:  "testpassw",
			},
		},
	}
	response.SUCCESS(w, http.StatusOK, allUsers)
}
