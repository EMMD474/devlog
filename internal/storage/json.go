package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

type Entry struct {
	Message string    `json:"message"`
	Date    time.Time `json:"date"`
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

func LoadEntries() ([]Entry, error) {
	path, err := dataFilePath()
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return []Entry{}, nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var entries []Entry
	if err := json.Unmarshal(data, &entries); err != nil {
		return nil, err
	}

	return entries, nil
}

func LoadEntriesSince(duration time.Duration) ([]Entry, error) {
	entries, err := LoadEntries()
	if err != nil {
		return nil, err
	}
	cutoffTime := time.Now().Add(-duration)

	var filteredEntries []Entry
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

	entry := Entry{
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
