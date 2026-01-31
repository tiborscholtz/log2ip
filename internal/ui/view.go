package ui

import (
	"strconv"
	"tiborscholtz/log2ip/internal/components"

	"github.com/charmbracelet/lipgloss"
	"github.com/common-nighthawk/go-figure"
)

var baseStyle = lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("240"))
var header = figure.NewFigure("LOG2IP", "", true)
var TableOptions = components.GetOptions()

func (m model) View() string {
	returnData := ""
	returnData += header.String()
	returnData += "\n"
	returnData += m.TextInput.View()
	returnData += "\n"
	returnData += baseStyle.Render(m.Table.View())
	returnData += "\n\n"
	returnData += strconv.Itoa(m.Page)
	returnData += "/"
	returnData += strconv.Itoa(m.TotalPage)
	returnData += "\n\n"
	for i := 0; i < len(TableOptions); i++ {
		returnData += "(" + TableOptions[i].Key + ") - " + TableOptions[i].Text + "\n"
	}
	return returnData
}
