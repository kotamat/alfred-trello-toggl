package main
import (
	"log"
	"encoding/json"
)


type WaitAction struct{}


func (c WaitAction) IsEnabled() bool {
	return config.ApiKey == ""
}

func(c WaitAction) Keyword() string {
	return "wait"
}

func(c WaitAction) Do(query string) (string, error) {
	log.Printf("wait '%s'", query)

	var data startMessage
	err := json.Unmarshal([]byte(query), &data)
	if err != nil {
		return "", err
	}
	return "", err
}