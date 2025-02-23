package main

import (
	"log"
	"net/http"
	"os"

	"github.com/MenonVishnu/Video-Ad-Player/backend/controllers"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not find .env file")
	}

	http.HandleFunc("/api/v1/ads", controllers.GetAds)
	http.HandleFunc("/api/v1/ads/click", controllers.LogClick)

	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), nil))
}
