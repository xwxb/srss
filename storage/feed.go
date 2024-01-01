package storage

import (
	// "database/sql"
	"log"
)

type Feed struct {
	Id          int64   `json:"id"`
	FolderId    *int64  `json:"folder_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Link        string  `json:"link"`
	FeedLink    string  `json:"feed_link"`
	Icon        *[]byte `json:"icon,omitempty"`
	HasIcon     bool    `json:"has_icon"`
}
