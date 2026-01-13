package edit

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/emmd474/devlog/internal/model"
)

func Run(entries []model.Entry) error {
	m := NewSelectModel(entries)
	p := tea.NewProgram(m)
	_, err := p.Run()
	return err
}