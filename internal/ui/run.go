package ui

import (
	"flag"
	"fmt"
	"os"
	"tiborscholtz/log2ip/internal/parsers"

	tea "github.com/charmbracelet/bubbletea"
)

func Run(){
	parsers.GetPossibleSimpleTextList()
	var filter string
	flag.StringVar(&filter,"filter", "", "Filter expression")
	var rowLimit int
	flag.IntVar(&rowLimit,"rows", 30, "Rows per page")
	flag.Parse()
	p := tea.NewProgram(InitialModel(filter,rowLimit))
	if _,err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
