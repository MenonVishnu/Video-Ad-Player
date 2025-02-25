package helpers

import (
	"encoding/json"
	"net"
	"net/http"
	"strings"
)

// Function for Error response
func ErrorResponse(w http.ResponseWriter, statusCode int, message string, errors interface{}) {
	var response Response

	response.Message = message
	response.Error = errors
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

// Function for Success response
func SuccessResponse(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	var response Response

	response.Message = message
	response.Data = data
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

// Function to get IP address
func GetIP(r *http.Request) string {
	// Check X-Forwarded-For first
	forwarded := r.Header.Get("X-Forwarded-For")
	if forwarded != "" {
		// X-Forwarded-For may contain multiple IPs, take the first one
		ips := strings.Split(forwarded, ",")
		return strings.TrimSpace(ips[0])
	}

	// Check X-Real-IP
	realIP := r.Header.Get("X-Real-IP")
	if realIP != "" {
		return realIP
	}

	// Fallback to RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}
