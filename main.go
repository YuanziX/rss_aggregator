package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/yuanzix/rss_aggregator/handlers"
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

	router := http.NewServeMux()

	corsHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" {
			origin = "*"
		}
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Expose-Headers", "Link")
		w.Header().Set("Access-Control-Max-Age", "300")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		router.ServeHTTP(w, r)
	})

	v1Router := http.NewServeMux()
	v1Router.HandleFunc("/health", handlers.HandlerReadiness)
	v1Router.HandleFunc("/err", handlers.HandlerErr)

	router.Handle("/v1/", http.StripPrefix("/v1", v1Router))

	server := &http.Server{
		Handler: corsHandler,
		Addr:    ":" + port,
	}

	log.Printf("Server starting on port %v.\n", port)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
