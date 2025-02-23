package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/MenonVishnu/Video-Ad-Player/backend/database"
	"github.com/MenonVishnu/Video-Ad-Player/backend/helpers"
)

func GetAds(w http.ResponseWriter, r *http.Request) {
	//Set Headers
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	//Take data from database
	var advs []helpers.AdvData
	advs, err := database.GetAllAdv()
	if err != nil {
		helpers.ErrorResponse(w, 503, "Unable to Get Advertisement", err)
	}

	//Send response
	helpers.SuccessResponse(w, 200, "Data retrieved Successfully", advs)

}

func LogClick(w http.ResponseWriter, r *http.Request) {
	//Set Headers
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	//Take the data from body
	var clickData helpers.ClickData
	_ = json.NewDecoder(r.Body).Decode(&clickData)

	//Store it in sqlite
	err := database.AddClick(clickData)
	if err != nil {
		helpers.ErrorResponse(w, 503, "Unable to Add Click Data", err)
	}

	//Send response
	helpers.SuccessResponse(w, 201, "Click Logged Successfully", nil)

}
