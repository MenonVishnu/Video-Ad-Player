package helpers

import (
	"encoding/json"
	"net/http"
)

// GET Request Data
type AdvData struct {
	AdID      int    `json:"ad_id"`
	ImageUrl  string `json:"image_url"`
	TargetUrl string `json:"target_url"`
}

// POST Request Data
type ClickData struct {
	AdID           int     `json:"ad_id"`
	Timestamp      float64 `json:"timestamp"`
	IP             string  `json:"ip"`
	VideoTimeStamp float64 `json:"video_timestamp"`
}

// Response Data
type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
}

//Function for Error response
func ErrorResponse(w http.ResponseWriter, statusCode int, message string, errors interface{}) {
	var response Response

	response.Message = message
	response.Error = errors
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}


//Function for Success response
func SuccessResponse(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	var response Response

	response.Message = message
	response.Data = data
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
