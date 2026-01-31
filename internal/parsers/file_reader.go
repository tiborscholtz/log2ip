package parsers

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"tiborscholtz/log2ip/internal/config"
	"tiborscholtz/log2ip/internal/domain"
)

func ParseFile(path string, filetype string) ([]domain.LogEntry, error) {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	cacheName := pwd + "/internal/cache/" + path + ".json"
	entries := make([]domain.LogEntry, 0, 0)
	cacheData, cacheErr := os.ReadFile(cacheName)
	if cacheErr == nil {
		var objmap map[string]interface{}
		json.Unmarshal(cacheData, &objmap)
		objMapData := objmap["data"].([]interface{})
		for i := 0; i < len(objMapData); i++ {
			currentEntry := objMapData[i].(map[string]interface{})
			logEntry := domain.LogEntry{Hostname: currentEntry["Hostname"].(string), Address: currentEntry["Address"].(string), Date: currentEntry["Date"].(string), ServiceName: currentEntry["ServiceName"].(string), SimpleMessage: currentEntry["SimpleMessage"].(string), Message: currentEntry["Message"].(string)}
			entries = append(entries, logEntry)
		}
	} else {
		file, err := os.Open(pwd + "/" + config.TEST_DATA_DIRECTORY + "/" + path + ".log")
		if err != nil {
			return nil, errors.New("error_in_opening_file")
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		amount := 0
		limit := 500
		for scanner.Scan() {
			if filetype == "openssh" {
				entry, err := ParseSSHLogLine(scanner.Text())
				if err == nil {
					amount = amount + 1
					if amount == limit {
						break
					}
					logEntry := domain.LogEntry{Hostname: entry.Hostname, Address: entry.Address, Date: entry.Date, ServiceName: entry.ServiceName, SimpleMessage: entry.SimpleMessage, Message: entry.Message}
					entries = append(entries, logEntry)
				}
			} else if filetype == "auth" {
				entry, err := ParseAuthLogLine(scanner.Text())
				if err == nil {
					amount = amount + 1
					if amount == limit {
						break
					}
					logEntry := domain.LogEntry{Hostname: entry.Hostname, Address: entry.Address, Date: entry.Date, ServiceName: entry.ServiceName, SimpleMessage: entry.SimpleMessage, Message: entry.Message}
					entries = append(entries, logEntry)
				}
			} else {
				amount = 0
				break
			}
		}
	}
	if cacheErr != nil {
		cachedData := make(map[string]interface{})
		cachedData["version"] = "1"
		cachedData["data"] = entries
		jsonBytes, error := json.Marshal(cachedData)
		if error == nil {
			filename := pwd + "/internal/cache/" + path + ".json"
			os.WriteFile(filename, jsonBytes, 0644)
		}
	}
	if len(entries) == 0 {
		return nil, errors.New("no_entry_in_the_log_file")
	}
	return entries, nil
}
