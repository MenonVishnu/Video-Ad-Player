package helpers

// GET Request Data
type AdvData struct {
	AdID      int    `json:"ad_id"`
	ImageUrl  string `json:"image_url"`
	TargetUrl string `json:"target_url"`
}

// POST Request Data
type ClickData struct {
	AdID           int     `json:"ad_id"`
	Timestamp      string  `json:"timestamp"`
	IP             string  `json:"ip"`
	VideoTimeStamp float64 `json:"video_timestamp"`
}

// Response Data
type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
}
