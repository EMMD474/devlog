package storage

import (
	"github.com/emmd474/devlog/internal/model"
)

func UpdateEntry(updated model.Entry) error {
	entries, err := LoadEntries()
	if err != nil {
		return err
	}

	for i, e := range entries {
		if e.ID == updated.ID {
			entries[i] = updated
			break
		}
	}

	return SaveEntries(entries)
}