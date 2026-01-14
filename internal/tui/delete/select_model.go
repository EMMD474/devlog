package delete

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/emmd474/devlog/internal/model"
	"github.com/emmd474/devlog/internal/storage"
)

var (
	listStyle    = lipgloss.NewStyle().Margin(1, 2)
	confirmStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#F38BA8")).
			Bold(true).
			Margin(1, 2)
	successStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#A6E3A1")).
			Margin(1, 2)
)

type selectModel struct {
	list       list.Model
	entries    []model.Entry
	width      int
	height     int
	confirming bool
	selected   model.Entry
	deleted    bool
	err        error
}

func NewSelectModel(entries []model.Entry) selectModel {
	items := make([]list.Item, len(entries))
	for i, e := range entries {
		items[i] = item{entry: e}
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Select a log to delete"
	l.SetShowStatusBar(true)
	l.SetFilteringEnabled(true)

	return selectModel{
		list:    l,
		entries: entries,
	}
}

func (m selectModel) Init() tea.Cmd {
	return nil
}

func (m selectModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		h, v := listStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
		return m, nil

	case tea.KeyMsg:
		if m.deleted {
			return m, tea.Quit
		}

		if m.confirming {
			switch msg.String() {
			case "y", "Y":
				if err := m.deleteEntry(); err != nil {
					m.err = err
					return m, tea.Quit
				}
				m.deleted = true
				return m, nil
			case "n", "N", "esc":
				m.confirming = false
				return m, nil
			}
			return m, nil
		}

		switch msg.String() {
		case "enter":
			if m.list.FilterState() == list.Filtering {
				break
			}
			m.selected = m.list.SelectedItem().(item).entry
			m.confirming = true
			return m, nil
		case "q", "ctrl+c":
			if m.list.FilterState() == list.Filtering {
				break
			}
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m *selectModel) deleteEntry() error {
	var remaining []model.Entry
	for _, e := range m.entries {
		if e.ID != m.selected.ID {
			remaining = append(remaining, e)
		}
	}
	return storage.SaveEntries(remaining)
}

func (m selectModel) View() string {
	if m.deleted {
		return successStyle.Render("Entry deleted successfully. Press any key to exit.")
	}

	if m.confirming {
		return confirmStyle.Render(
			fmt.Sprintf("Delete this entry?\n\n  %s\n  %s\n\nPress y to confirm, n to cancel",
				m.selected.Date.Format("2006-01-02 15:04"),
				m.selected.Message,
			),
		)
	}

	return listStyle.Render(m.list.View())
}
