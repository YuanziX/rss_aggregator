package main

import (
	"net/http"

	"github.com/yuanzix/rss_aggregator/handlers"
)

func InitializeRoutes() *http.ServeMux {
	router := http.NewServeMux()

	v1Router := http.NewServeMux()
	v1Router.HandleFunc("/health", handlers.HandlerReadiness)
	v1Router.HandleFunc("/err", handlers.HandlerErr)

	router.Handle("/v1/", http.StripPrefix("/v1", v1Router))

	return router
}
