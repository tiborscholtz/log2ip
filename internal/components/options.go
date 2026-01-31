package components

type TableOption struct {
	Text string
	Key  string
}

func GetOptions() []TableOption {
	return []TableOption{TableOption{Text: "Show simple text", Key: "s"}, TableOption{Text: "Refresh data", Key: "r"}, TableOption{Text: "Quit", Key: "q"}}
}
