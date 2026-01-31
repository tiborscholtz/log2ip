package parsers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"tiborscholtz/log2ip/internal/config"
)

func getSimpleText(original string) string{
	processedText := original
	for key, val := range config.SIMPLIFIED_MESSAGES{
		if idx := strings.Index(original, key); idx != -1 {
			strVal := string(val)
			processedText = strVal[1 : len(strVal)-1]
			break
		}
	}
	return processedText
}

func GetPossibleSimpleTextList(){
	pwd, err := os.Getwd()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
	jsonFile, err := os.Open(pwd+"/internal/resources/simplified_log_pattern.json")
	if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &config.SIMPLIFIED_MESSAGES)
}
