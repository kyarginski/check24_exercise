package models

import (
	"time"
)

// BlogEntry represents a blog entry.
type BlogEntry struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	CreationDate time.Time `json:"creation_date"`
	Author       string    `json:"author"`
	Text         string    `json:"text"`
	ImageLink    string    `json:"image_link"`
}

// Comment represents a comment on a blog entry.
type Comment struct {
	ID          string    `json:"id"`
	EntryID     string    `json:"entry_id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	URL         string    `json:"url"`
	CommentText string    `json:"comment_text"`
	CommentDate time.Time `json:"comment_date"`
}
