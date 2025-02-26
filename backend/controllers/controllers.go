package controllers

import (
	"encoding/json"
	"errors"
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
		helpers.ErrorResponse(w, 500, "Unable to Get Advertisement", err)
		return
	}

	//Handle no data in advertisement table
	if len(advs) == 0 {
		helpers.ErrorResponse(w, 500, "No Data found", errors.New("no data available in advertisement table"))
		return
	}

	//Send Success response
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

	//Handles null/no values in request
	if clickData.Timestamp == "" || clickData.AdID == 0 {
		helpers.ErrorResponse(w, 400, "Missing required files ", errors.New("the request body is missing required fields. please include all necessary fields"))
		return
	}

	//Collect IP from Request
	clickData.IP = helpers.GetIP(r)

	//Store the recieved data in database
	err := database.AddClick(clickData)
	if err != nil {
		helpers.ErrorResponse(w, 500, "Unable to Add Click Data", err)
		return
	}

	//Send success response
	helpers.SuccessResponse(w, 201, "Click Logged Successfully", nil)

}
