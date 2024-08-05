package handlers

import (
	"net/http"

	"github.com/yuanzix/rss_aggregator/utils"
)

func HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, 200, struct{}{})
}
