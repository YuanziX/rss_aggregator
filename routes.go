package main

import (
	"net/http"

	"github.com/yuanzix/rss_aggregator/handlers"
)

func InitializeRoutes(apiCfg handlers.ApiConfig) *http.ServeMux {
	router := http.NewServeMux()

	v1Router := http.NewServeMux()
	v1Router.HandleFunc("GET /health", handlers.HandlerReadiness)

	v1Router.HandleFunc("GET /err", handlers.HandlerErr)

	v1Router.HandleFunc("GET /users", apiCfg.MiddlewareAuth(apiCfg.HandlerGetUserByApiKey))
	v1Router.HandleFunc("POST /users", apiCfg.HandlerCreateUser)

	v1Router.HandleFunc("GET /feeds", apiCfg.HandlerGetFeeds)
	v1Router.HandleFunc("POST /feeds", apiCfg.MiddlewareAuth(apiCfg.HandlerCreateFeed))

	v1Router.HandleFunc("GET /posts", apiCfg.MiddlewareAuth(apiCfg.HandlerGetPostsForUser))

	v1Router.HandleFunc("GET /feed_follows", apiCfg.MiddlewareAuth(apiCfg.HandlerGetFeedFollows))
	v1Router.HandleFunc("POST /feed_follows", apiCfg.MiddlewareAuth(apiCfg.HandlerCreateFeedFollow))
	v1Router.HandleFunc("DELETE /feed_follows/{feedFollowID}", apiCfg.MiddlewareAuth(apiCfg.HandlerDeleteFeedFollows))

	router.Handle("/v1/", http.StripPrefix("/v1", v1Router))

	return router
}
