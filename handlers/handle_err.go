package handlers

import (
	"net/http"

	"github.com/yuanzix/rss_aggregator/utils"
)

func HandlerErr(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithError(w, 200, "Something went wrong")
}
