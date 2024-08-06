package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/yuanzix/rss_aggregator/internal/database"
	"github.com/yuanzix/rss_aggregator/utils"
)

func (apiCfg *ApiConfig) HandlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})

	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Couldn't create feed follow: %v", err))
		return
	}

	utils.RespondWithJSON(w, 201, utils.DatabaseFeedFollowToFeedFollow(feedFollow))
}

func (apiCfg *ApiConfig) HandlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)

	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Couldn't get feed follows: %v", err))
		return
	}

	utils.RespondWithJSON(w, 200, utils.DatabaseFeedsFollowsToFeedsFollows(feedFollows))
}

func (apiCfg *ApiConfig) HandlerDeleteFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowID, err := uuid.Parse(r.PathValue("feedFollowID"))
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Couldn't read supplied feed follow ID: %v", err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})

	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Couldn't delete feed follow: %v", err))
		return
	}

	utils.RespondWithJSON(w, 200, struct{}{})
}
