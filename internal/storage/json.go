package storage

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"time"

	"github.com/emmd474/devlog/internal/model"
)

// generateID creates a unique ID using crypto/rand
func generateID() (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func dataFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	dir := filepath.Join(home, ".devlog")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}

	return filepath.Join(dir, "logs.json"), nil
}

func LoadEntries() ([]model.Entry, error) {
	path, err := dataFilePath()
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return []model.Entry{}, nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var entries []model.Entry
	if err := json.Unmarshal(data, &entries); err != nil {
		return nil, err
	}

	// Ensure all entries have IDs
	modified := false
	for i := range entries {
		if entries[i].ID == "" {
			id, err := generateID()
			if err != nil {
				return nil, err
			}
			entries[i].ID = id
			modified = true
		}
	}

	// Save entries if we added IDs
	if modified {
		if err := SaveEntries(entries); err != nil {
			return nil, err
		}
	}

	return entries, nil
}

func LoadEntriesSince(duration time.Duration) ([]model.Entry, error) {
	entries, err := LoadEntries()
	if err != nil {
		return nil, err
	}
	cutoffTime := time.Now().Add(-duration)

	var filteredEntries []model.Entry
	for _, entry := range entries {
		if entry.Date.After(cutoffTime) || entry.Date.Equal(cutoffTime) {
			filteredEntries = append(filteredEntries, entry)
		}
	}

	return filteredEntries, nil

}

func SaveEntry(message string) error {
	entries, err := LoadEntries()
	if err != nil {
		return err
	}

	id, err := generateID()
	if err != nil {
		return err
	}

	entry := model.Entry{
		ID:      id,
		Message: message,
		Date:    time.Now(),
	}

	entries = append(entries, entry)

	data, err := json.MarshalIndent(entries, "", "  ")
	if err != nil {
		return err
	}

	path, err := dataFilePath()
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

func SaveEntries(entries []model.Entry) error {
	data, err := json.MarshalIndent(entries, "", "  ")
	if err != nil {
		return err
	}

	path, err := dataFilePath()
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
