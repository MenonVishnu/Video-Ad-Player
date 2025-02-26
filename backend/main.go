package main

import (
	"log"
	"net/http"
	"os"

	"github.com/MenonVishnu/Video-Ad-Player/backend/controllers"
	"github.com/joho/godotenv"
)

func main() {
	//Handle env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not find .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	//Handle Routes
	http.HandleFunc("/api/v1/ads", controllers.GetAds)
	http.HandleFunc("/api/v1/ads/click", controllers.LogClick)

	//Start Backend Server
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
