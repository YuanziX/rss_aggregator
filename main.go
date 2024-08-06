package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"

	"github.com/yuanzix/rss_aggregator/handlers"
	"github.com/yuanzix/rss_aggregator/internal/database"
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

	db, err := sql.Open("sqlite3", "./sql/rss_aggregator_db.db")
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}

	queries := database.New(db)

	apiCfg := handlers.ApiConfig{
		DB: queries,
	}

	router := InitializeRoutes(apiCfg)

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
