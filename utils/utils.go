package utils

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// Message .
func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

// Respond .
func Respond(w http.ResponseWriter, data map[string]interface{}, status uint) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
	w.Header().Add("X-XSS-Protection", "1; mode=block")
	w.Header().Add("X-Frame-Options", "Deny")
	w.Header().Add("Content-Security-Policy", "script-src 'self'")
	w.Header().Add("X-Content-Type-Options", "nosniff")
	w.Header().Add("Referrer-Policy", "no-referrer")
	w.Header().Add("Feature-Policy", "vibrate 'none'; geolocation 'none'")

	switch status {
	case 200: // break 200 Accept Request
		w.WriteHeader(http.StatusOK)
		break
	case 201: // break 201 created POST
		w.WriteHeader(http.StatusCreated)
		break
	case 204: // break 204 No Content (Just Delete Http)
		w.WriteHeader(http.StatusNoContent)
		break
	case 301: // break 301 Moved Permanently
		w.WriteHeader(http.StatusMovedPermanently)
		break
	case 400: // break 400 Bad Request
		w.WriteHeader(http.StatusBadRequest)
		break
	case 401: // break 401 Unauthorized
		w.WriteHeader(http.StatusUnauthorized)
		break
	case 403: // break 403 Forbidden
		w.WriteHeader(http.StatusForbidden)
		break
	case 404: // break 404 Not Found
		w.WriteHeader(http.StatusNotFound)
		break
	case 500: // break 500 Internal Server Error
		w.WriteHeader(http.StatusInternalServerError)
		break
	}

	data["status"] = status
	json.NewEncoder(w).Encode(data)
}

// ConverStringToUint .
func ConverStringToUint(text string) uint {
	temp, _ := strconv.ParseUint(text, 10, 32)
	return uint(temp)
}
