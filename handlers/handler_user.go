package handlers

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/yuanzix/rss_aggregator/internal/database"
	"github.com/yuanzix/rss_aggregator/utils"
)

func (apiCfg *ApiConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		ApiKey:    generateRandomSha256Hex(),
	})
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	utils.RespondWithJSON(w, 201, utils.DatabaseUserToUser(user))
}

func generateRandomSha256Hex() string {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}

	hash := sha256.Sum256(randomBytes)

	return hex.EncodeToString(hash[:])
}

func (apiCfg *ApiConfig) HandlerGetUserByApiKey(w http.ResponseWriter, r *http.Request, user database.User) {
	utils.RespondWithJSON(w, 200, utils.DatabaseUserToUser(user))
}

func (apiCfg *ApiConfig) HandlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	log.Printf("Getting posts for user %s", user.ID)
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})

	if err != nil {
		utils.RespondWithError(w, 500, fmt.Sprintf("Error getting posts: %v", err))
		return
	}

	log.Printf("Got %d posts", len(posts))

	utils.RespondWithJSON(w, 200, utils.DatabasePostsToPosts(posts))
}
