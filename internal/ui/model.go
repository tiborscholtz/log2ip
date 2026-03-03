package ui

import (
	"strconv"
	"strings"
	"tiborscholtz/log2ip/internal/domain"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	Table             table.Model
	Data              []domain.LogEntry
	Filter            string
	TextInput         textinput.Model
	ShowSimpleMessage bool
	RowLimit          int
	Page              int
	Limit             int
	TotalPage         int
	SelectedRow       int
	Mode              string
	ColumnIndex       int
}

func matchesSearch(e domain.LogEntry, search string, simple bool) bool {
	if search == "" {
		return true
	}

	search = strings.ToLower(search)

	message := e.Message
	if simple {
		message = e.SimpleMessage.(string)
	}

	return strings.Contains(strings.ToLower(e.Date), search) ||
		strings.Contains(strings.ToLower(e.Address), search) ||
		strings.Contains(strings.ToLower(e.ServiceName), search) ||
		strings.Contains(strings.ToLower(message), search)
}

func currentColumn(defaultText string, i int, m model) string {
	if i == m.ColumnIndex && m.Mode == "table" {
		return defaultText + "*"
	}
	return defaultText
}

func CreateCurrentTable(m model) table.Model {
	columns := []table.Column{
		{Title: currentColumn("#", 0, m), Width: 3},
		{Title: currentColumn("Time", 1, m), Width: 20},
		{Title: currentColumn("Address", 2, m), Width: 20},
		{Title: currentColumn("Service", 3, m), Width: 10},
		{Title: currentColumn("Text", 4, m), Width: 200},
	}
	rows := make([]domain.LogEntry, 0, 0)
	for i := 0; i < len(m.Data); i++{
		e := m.Data[i]
		if matchesSearch(e, m.TextInput.Value(), m.ShowSimpleMessage) {
			rows = append(rows, e)
		}
	}
	finalRows := make([]table.Row, 0, 0)
	for i := (m.Limit * m.Page) - m.Limit; i < ((m.Limit*m.Page)+m.Limit)-m.Limit; i++ {
		if i > (len(rows) - 1) {
			break
		}
		e := rows[i]
		current_message := e.Message
		if m.ShowSimpleMessage == true {
			current_message = e.SimpleMessage.(string)
		}
		finalRows = append(finalRows, table.Row{strconv.Itoa(i + 1), e.Date, e.Address, e.ServiceName, current_message})
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(finalRows),
		table.WithFocused(true),
		table.WithHeight(m.RowLimit+1),
	)
	s := table.DefaultStyles()
	s.Header = s.Header.BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)
	return t
}
