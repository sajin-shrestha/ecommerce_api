package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ISuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SUCCESS(w http.ResponseWriter, code int, data interface{}) {
	messages := map[int]string{
		http.StatusOK:             fmt.Sprintf("Success: OK"),
		http.StatusCreated:        fmt.Sprintf("Success: Created"),
		http.StatusAccepted:       fmt.Sprintf("Success: Request Accepted"),
		http.StatusNoContent:      fmt.Sprintf("Success: No Content Returned"),
		http.StatusResetContent:   fmt.Sprintf("Success: Reset View"),
		http.StatusPartialContent: fmt.Sprintf("Success: Partial Content Returned"),
	}

	message := messages[code]
	if message == "" {
		message = fmt.Sprintf("Success: Operation completed (%d)", code)
	}

	response := ISuccessResponse{
		Message: message,
	}

	response.Data = data

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}
