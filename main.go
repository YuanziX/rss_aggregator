package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/yuanzix/rss_aggregator/utils"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Unable to load environment variables.")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT not configured in the environment.")
	}

	router := InitializeRoutes()

	server := &http.Server{
		Handler: utils.CORSHandler(router),
		Addr:    ":" + port,
	}

	log.Printf("Server starting on port %v.\n", port)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
