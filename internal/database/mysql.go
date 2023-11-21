package database

import (
	"database/sql"
	"fmt"
	"time"

	"check24/internal/models"
)

type Storage struct {
	db *sql.DB
}

func (s *Storage) Close() error {
	return s.db.Close()
}

func setDB(path string) (*sql.DB, error) {
	db, err := sql.Open("mysql", path)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (s *Storage) GetDB() *sql.DB {
	return s.db
}

func New(path string) (*Storage, error) {
	const op = "database.New"

	db, err := setDB(path)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) CreateBlogEntry(entry *models.BlogEntry) error {
	query := "INSERT INTO blog_entries (title, creation_date, author, text, image_link) VALUES (?, ?, ?, ?, ?)"
	_, err := s.db.Exec(query, entry.Title, entry.CreationDate, entry.Author, entry.Text, entry.ImageLink)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) GetBlogEntries() ([]*models.BlogEntry, error) {
	// TODO realize some filters.
	query := "SELECT * FROM blog_entries"

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []*models.BlogEntry
	for rows.Next() {
		var entry models.BlogEntry
		var creationDateStr string
		var imageLink sql.NullString
		err = rows.Scan(&entry.ID, &entry.Title, &creationDateStr, &entry.Author, &entry.Text, &imageLink)
		if err != nil {
			return nil, err
		}
		entry.CreationDate, err = time.Parse("2006-01-02 15:04:05", creationDateStr)
		if err != nil {
			return nil, err
		}
		if imageLink.Valid {
			entry.ImageLink = imageLink.String
		} else {
			entry.ImageLink = "" // or default value.
		}

		entries = append(entries, &entry)
	}

	return entries, nil
}

func (s *Storage) GetBlogEntry(id string) (*models.BlogEntry, error) {
	query := "SELECT * FROM blog_entries WHERE entry_id = ?"

	var entry models.BlogEntry
	var creationDateStr string
	var imageLink sql.NullString

	err := s.db.QueryRow(query, id).Scan(&entry.ID, &entry.Title, &creationDateStr, &entry.Author, &entry.Text, &imageLink)
	if err != nil {
		return nil, err
	}
	entry.CreationDate, err = time.Parse("2006-01-02 15:04:05", creationDateStr)
	if err != nil {
		return nil, err
	}
	if imageLink.Valid {
		entry.ImageLink = imageLink.String
	} else {
		entry.ImageLink = "" // or default value.
	}

	return &entry, nil
}
