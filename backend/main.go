package main

import (
	"log"
	"net/http"

	"github.com/MenonVishnu/Video-Ad-Player/backend/controllers"
)

func main() {

	http.HandleFunc("/api/v1/ads", controllers.GetAds)
	http.HandleFunc("/api/v1/ads/click", controllers.LogClick)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
