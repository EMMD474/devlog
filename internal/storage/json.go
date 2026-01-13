package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"

	"github.com/emmd474/devlog/internal/model"
)

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

	entry := model.Entry{
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
