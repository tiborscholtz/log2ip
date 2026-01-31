package components

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

func CreateTextInput(filter string) textinput.Model {
	ti := textinput.New()
	ti.SetValue(filter)
	ti.Placeholder = "Search in logs..."
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 50
	ti.PromptStyle.AlignHorizontal(lipgloss.Center)
	ti.PromptStyle.Align(lipgloss.Center)
	return ti
}
