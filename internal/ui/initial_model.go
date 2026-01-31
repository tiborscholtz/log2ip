package ui

import (
	"tiborscholtz/log2ip/internal/components"
	"tiborscholtz/log2ip/internal/parsers"

	"github.com/charmbracelet/bubbles/table"
)

func InitialModel(filter string, rowLimit int) model {
	if rowLimit == 0 {
		rowLimit = 5
	}
	ti := components.CreateTextInput(filter)
	entries, _ := parsers.ParseFile("OpenSSH_2k", "openssh")
	auth_logs, _ := parsers.ParseFile("Linux_2k", "auth")
	entries = append(entries, auth_logs...)
	totalPage := (len(entries) - (len(entries) % rowLimit)) / rowLimit
	t := table.New()
	m := model{
		Table:             t,
		Data:              entries,
		Filter:            filter,
		TextInput:         ti,
		ShowSimpleMessage: true,
		RowLimit:          rowLimit,
		Page:              1,
		Limit:             rowLimit,
		TotalPage:         totalPage,
		SelectedRow:       0,
		Mode:              "default",
		ColumnIndex:       0,
	}
	m.Table = CreateCurrentTable(m)
	return m
}
