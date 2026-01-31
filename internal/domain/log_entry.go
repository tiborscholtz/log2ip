package domain

type LogEntry struct {
	Date        string
	Hostname    string
	ServiceName string
	Message     string
	Address string
	SimpleMessage any
}
