package handlers

import (
	"net/http"

	"github.com/yuanzix/rss_aggregator/utils"
)

func HandlerErr(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	utils.RespondWithError(w, 200, "Something went wrong")
}
