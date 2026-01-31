package parsers

import (
	"strings"
	"regexp"
	"errors"
)

type AuthLogEntry struct {
	Date        string
	Hostname    string
	ServiceName string
	Message     string
	Address string
	SimpleMessage any
}
func ParseAuthLogLine(line string) (AuthLogEntry, error) {
	fields := strings.Fields(line)
	if len(fields) < 5 {
		return AuthLogEntry{},errors.New("not_enough_fields")
	}

	re := regexp.MustCompile(`((25[0-5]|2[0-4]\d|1\d{2}|[1-9]?\d)(\.)){3}(25[0-5]|2[0-4]\d|1\d{2}|[1-9]?\d)`)
	match := re.FindString(line)
	if match == ""{
		return AuthLogEntry{},errors.New("no_ip_address_found")
	}
	address := match
	date := strings.Join(fields[0:3], " ")

	hostname := fields[3]

	rest := strings.Join(fields[4:], " ")

	parts := strings.SplitN(rest, ": ", 2)
	if len(parts) != 2 {
		return AuthLogEntry{},errors.New("not_enough_parts")
	}

	// Service name may include [pid], remove it
	service := parts[0]
	if idx := strings.Index(service, "["); idx != -1 {
		service = service[:idx]
	}
	if idx := strings.Index(service, "("); idx != -1 {
		service = service[:idx]
	}

	message := parts[1]


	return AuthLogEntry{
		Date:        date,
		Hostname:    hostname,
		ServiceName: service,
		Message:     message,
		SimpleMessage: getSimpleText(message),
		Address: address,
	}, nil
}
