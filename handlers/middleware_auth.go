package handlers

import (
	"fmt"
	"net/http"

	"github.com/yuanzix/rss_aggregator/internal/auth"
	"github.com/yuanzix/rss_aggregator/internal/database"
	"github.com/yuanzix/rss_aggregator/utils"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *ApiConfig) MiddlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)

		if err != nil {
			utils.RespondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByApiKey(r.Context(), apiKey)

		if err != nil {
			utils.RespondWithError(w, 403, fmt.Sprintf("Couldn't get user: %v", err))
			return
		}

		handler(w, r, user)
	}
}
