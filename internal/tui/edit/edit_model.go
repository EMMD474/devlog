package edit

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/emmd474/devlog/internal/model"
	"github.com/emmd474/devlog/internal/storage"
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#7D56F4")).
			MarginBottom(1)

	dateStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#6C7086")).
			MarginBottom(1)

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#6C7086")).
			MarginTop(1)

	successStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#A6E3A1")).
			Bold(true)

	errorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#F38BA8")).
			Bold(true)

	containerStyle = lipgloss.NewStyle().
			Margin(1, 2)
)

type editModel struct {
	entry    model.Entry
	textarea textarea.Model
	width    int
	height   int
	err      error
	saved    bool
}

func NewEditModel(entry model.Entry, width, height int) editModel {
	ta := textarea.New()
	ta.SetValue(entry.Message)
	ta.Focus()
	ta.CharLimit = 500

	if width > 0 {
		ta.SetWidth(width - 8)
	}
	if height > 0 {
		ta.SetHeight(height - 10)
	}

	return editModel{
		entry:    entry,
		textarea: ta,
		width:    width,
		height:   height,
	}
}

func (m editModel) Init() tea.Cmd {
	return textarea.Blink
}

func (m editModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.textarea.SetWidth(msg.Width - 8)
		m.textarea.SetHeight(msg.Height - 10)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+s":
			m.entry.Message = m.textarea.Value()
			if err := storage.UpdateEntry(m.entry); err != nil {
				m.err = err
				return m, nil
			}
			m.saved = true
			return m, tea.Quit

		case "esc", "ctrl+c":
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.textarea, cmd = m.textarea.Update(msg)
	return m, cmd
}

func (m editModel) View() string {
	title := titleStyle.Render("Edit Log Entry")
	date := dateStyle.Render(m.entry.Date.Format("2006-01-02 15:04"))

	var status string
	if m.err != nil {
		status = errorStyle.Render(fmt.Sprintf("Error: %v", m.err))
	} else if m.saved {
		status = successStyle.Render("Saved!")
	}

	help := helpStyle.Render("ctrl+s save • esc cancel")

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		date,
		m.textarea.View(),
		status,
		help,
	)

	return containerStyle.Render(content)
}
