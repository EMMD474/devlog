package edit

import "github.com/emmd474/devlog/internal/model"

type item struct {
	entry model.Entry
}

func (i item) Title() string {
	return i.entry.Date.Format("2006-01-02 15:04")
}

func (i item) Description() string {
	return i.entry.Message
}

func (i item) FilterValue() string {
	return i.entry.Message
}