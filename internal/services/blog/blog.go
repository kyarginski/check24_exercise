package blog

import (
	"log/slog"

	"check24/internal/models"
)

type Entry struct {
	log          *slog.Logger
	blogSaver    Saver
	blogProvider Provider
}

type Saver interface {
	CreateBlogEntry(entry *models.BlogEntry) error
}

type Provider interface {
	GetBlogEntries() ([]*models.BlogEntry, error)
	GetBlogEntry(id int64) (*models.BlogEntry, error)
}

func New(log *slog.Logger, blogSaver Saver, blogProvider Provider) *Entry {
	return &Entry{
		log:          log,
		blogSaver:    blogSaver,
		blogProvider: blogProvider,
	}
}
