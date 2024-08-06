package utils

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/yuanzix/rss_aggregator/internal/database"
)

func StartScraping(
	db *database.Queries,
	concurrency int,
	timeBetweenRequests time.Duration,
) {
	log.Printf("Starting scraping every %v duration with %v workers.\n", timeBetweenRequests, concurrency)

	ticker := time.NewTicker(timeBetweenRequests)

	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(
			context.Background(),
			int64(concurrency),
		)

		if err != nil {
			log.Printf("Error getting feeds to fetch: %v\n", err)
			continue
		}

		wg := &sync.WaitGroup{}

		for _, feed := range feeds {
			wg.Add(1)

			go scrapeFeed(db, wg, feed)
		}
		wg.Wait()

	}
}

func scrapeFeed(
	db *database.Queries,
	wg *sync.WaitGroup,
	feed database.Feed,
) {
	defer wg.Done()

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)

	if err != nil {
		log.Printf("Error marking feed as fetched: %v\n", err)
		return
	}

	rssFeed, err := urlToFeed(feed.Url)

	if err != nil {
		log.Printf("Error fetching feed: %v\n", err)
		return
	}

	for _, post := range rssFeed.Channel.Item {
		description := sql.NullString{}
		t, err := time.Parse(time.RFC1123Z, post.PubDate)
		if err != nil {
			log.Printf("Error parsing publish date: %v with err %v\n", post.PubDate, err)
			continue
		}

		if post.Description != "" {
			description.String = post.Description
			description.Valid = true
		}

		_, err = db.CreatePost(
			context.Background(),
			database.CreatePostParams{
				ID:          uuid.New(),
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
				Title:       post.Title,
				Description: description,
				PublishedAt: t,
				Url:         post.Link,
				FeedID:      feed.ID,
			})

		if err != nil {
			if strings.Contains(err.Error(), "UNIQUE constraint") {
				continue
			}
			log.Printf("Error creating post: %v\n", err)
			continue
		}
	}

	log.Printf("Feed %s collected, %v posts found.\n", rssFeed.Channel.Title, len(rssFeed.Channel.Item))
}
