package handlers

import (
	"net/http"

	"github.com/yuanzix/rss_aggregator/utils"
)

func HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	utils.RespondWithJSON(w, 200, struct{}{})
}
