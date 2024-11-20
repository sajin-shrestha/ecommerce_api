package response

import (
	"encoding/json"
	"net/http"
)

type IErrorResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Detail     string `json:"detail"`
}

func ERROR(w http.ResponseWriter, code int, errDetail interface{}) {
	messages := map[int]string{
		/* client-side errors */
		http.StatusBadRequest:            "Client_Error: Bad Request",
		http.StatusUnauthorized:          "Client_Error: Unauthorized - Authentication required",
		http.StatusPaymentRequired:       "Client_Error: Payment Required",
		http.StatusForbidden:             "Client_Error: Forbidden - Access denied",
		http.StatusNotFound:              "Client_Error: Not Found",
		http.StatusMethodNotAllowed:      "Client_Error: Method Not Allowed",
		http.StatusNotAcceptable:         "Client_Error: Response not acceptable.",
		http.StatusProxyAuthRequired:     "Client_Error: Proxy Authentication Required",
		http.StatusRequestTimeout:        "Client_Error: Request Timeout",
		http.StatusConflict:              "Client_Error: Resource conflict.",
		http.StatusGone:                  "Client_Error: Resource no longer available.",
		http.StatusRequestEntityTooLarge: "Client_Error: Request size too large.",
		http.StatusRequestURITooLong:     "Client_Error: Request URI is too long.",
		http.StatusUnsupportedMediaType:  "Client_Error: Unsupported Media Type",
		http.StatusMisdirectedRequest:    "Client_Error: Misdirected Request - Wrong server for request.",
		http.StatusLocked:                "Client_Error: Locked - Resource is locked.",
		http.StatusFailedDependency:      "Client_Error: Failed Dependency - Request failed due to dependency.",
		http.StatusTooManyRequests:       "Client_Error: Too Many Requests",

		/* server-side errors */
		http.StatusInternalServerError:           "Server_Error: Internal Server Error",
		http.StatusNotImplemented:                "Server_Error: Not Implemented - Feature not supported.",
		http.StatusBadGateway:                    "Server_Error: Bad Gateway - Invalid response from upstream server.",
		http.StatusServiceUnavailable:            "Server_Error: Service Unavailable - Service temporarily unavailable.",
		http.StatusGatewayTimeout:                "Server_Error: Gateway Timeout - Upstream server didn't respond.",
		http.StatusHTTPVersionNotSupported:       "Server_Error: HTTP Version Not Supported - Unsupported HTTP version.",
		http.StatusInsufficientStorage:           "Server_Error: Insufficient Storage - Server out of storage.",
		http.StatusLoopDetected:                  "Server_Error: Loop Detected - Infinite loop in request.",
		http.StatusNotExtended:                   "Server_Error: Not Extended - Further extensions required.",
		http.StatusNetworkAuthenticationRequired: "Server_Error: Network Authentication Required - Network authentication needed.",
	}

	message := messages[code]
	if message == "" {
		message = "ERROR: Something Went Wrong. Try Again Later!"
	}

	var detail string
	switch v := errDetail.(type) {
	case string:
		detail = v
	case error:
		detail = v.Error()
	default:
		detail = "No additional message provided regarding this error"
	}

	response := IErrorResponse{
		StatusCode: code,
		Message:    message,
		Detail:     detail,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}
