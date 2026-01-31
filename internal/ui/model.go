package ui

import (
	"strconv"
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
	rows := make([]table.Row, 0, 0)
	for i := (m.Limit * m.Page) - m.Limit; i < ((m.Limit*m.Page)+m.Limit)-m.Limit; i++ {
		if i > (len(m.Data) - 1) {
			break
		}
		e := m.Data[i]
		current_message := e.Message
		if m.ShowSimpleMessage == true {
			current_message = e.SimpleMessage.(string)
		}
		rows = append(rows, table.Row{strconv.Itoa(i + 1), e.Date, e.Address, e.ServiceName, current_message})
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
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
