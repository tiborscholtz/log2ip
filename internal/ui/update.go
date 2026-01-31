package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	previousTextFocus := m.Mode == "search"
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab":
			if m.Mode == "default" {
				m.ColumnIndex = 0
				m.Mode = "table"
				m.Table = CreateCurrentTable(m)
			}
		case "ctrl+c", "q":
			if msg.String() == "q" && m.Mode == "search" {
				break
			}
			return m, tea.Quit
		case "enter":
			m.Mode = "default"
			m.Table = CreateCurrentTable(m)
		case "esc":
			if m.Mode == "search" {
				m.Mode = "default"
				m.TextInput.Blur()
			}
		case ":":
			if previousTextFocus == false {
				m.Mode = "search"
				m.TextInput.Focus()
				m.Table = CreateCurrentTable(m)
			}
		case "x":
			if m.Mode == "default" {
				index := m.Table.Cursor()
				m.Data = append(m.Data[:index], m.Data[index+1:]...)
				m.Table = CreateCurrentTable(m)
			}
		case "left":
			if m.Page > 1 {
				m.Page -= 1
				m.Table = CreateCurrentTable(m)
			}
		case "right":
			if m.Page < (m.TotalPage) {
				m.Page += 1
				m.Table = CreateCurrentTable(m)
			}
		case "up":
			if m.SelectedRow > 0 {
				m.SelectedRow -= 1
				m.Table.SetCursor(m.SelectedRow)
			}
		case "down":
			if m.SelectedRow < (m.RowLimit - 1) {
				m.SelectedRow += 1
				m.Table.SetCursor(m.SelectedRow)
			}
		case "s":
			if m.Mode == "default" {
				m.ShowSimpleMessage = !m.ShowSimpleMessage
				m.Table = CreateCurrentTable(m)
			}
		case "pgup":
			m.SelectedRow = 0
			m.Table.SetCursor(m.SelectedRow)
		case "pgdown":
			m.SelectedRow = m.RowLimit - 1
			m.Table.SetCursor(m.SelectedRow)
		}
		if m.Mode == "search" {
			if (previousTextFocus == false && msg.String() != ":") || (previousTextFocus == true) {
				m.TextInput, cmd = m.TextInput.Update(msg)
			}
		}
		return m, nil
	}

	m.Table, cmd = m.Table.Update(msg)
	return m, cmd
}
