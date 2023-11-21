package api

import (
	"encoding/json"
	"net/http"

	"check24/internal/database"
	"check24/internal/models"

	"github.com/gorilla/mux"
)

func GetBlogEntries(storage *database.Storage) http.HandlerFunc {
	// swagger: meta
	return func(w http.ResponseWriter, r *http.Request) {
		entries, err := storage.GetBlogEntries()
		if err != nil {
			http.Error(w, "error in GetBlogEntries", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(entries)
		if err != nil {
			http.Error(w, "error in GetBlogEntries", http.StatusInternalServerError)
			return
		}
	}
}

func GetBlogEntry(storage *database.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		entryID := vars["entry_id"]

		entry, err := storage.GetBlogEntry(entryID)
		if err != nil {
			http.Error(w, "error in GetBlogEntry", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(entry)
		if err != nil {
			http.Error(w, "error in GetBlogEntry", http.StatusInternalServerError)
			return
		}
	}
}

func CreateBlogEntry(storage *database.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var entry models.BlogEntry

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&entry); err != nil {
			http.Error(w, "Failed to decode request body: "+err.Error(), http.StatusBadRequest)
			return
		}

		defer r.Body.Close()

		err := storage.CreateBlogEntry(&entry)
		if err != nil {
			http.Error(w, "error in CreateBlogEntry", http.StatusInternalServerError)
			return
		}

		// Return a success response.
		w.WriteHeader(http.StatusCreated)
	}
}
