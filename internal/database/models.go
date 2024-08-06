// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

import (
	"database/sql"
	"time"
)

type Feed struct {
	ID            interface{}
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Name          string
	Url           string
	UserID        interface{}
	LastFetchedAt sql.NullTime
}

type FeedFollow struct {
	ID        interface{}
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    interface{}
	FeedID    interface{}
}

type Post struct {
	ID          interface{}
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Title       string
	Description sql.NullString
	PublishedAt time.Time
	Url         string
	FeedID      interface{}
}

type User struct {
	ID        interface{}
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	ApiKey    string
}
