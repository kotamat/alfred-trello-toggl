package main

import (
	"log"

	"github.com/jason0x43/go-alfred"
)

type TokenCommand struct{}

func (c TokenCommand) Keyword() string {
	return "token"
}

func (c TokenCommand) IsEnabled() bool {
//	return config.ApiKey == ""
	return true
}

func (c TokenCommand) MenuItem() alfred.Item {
	return alfred.Item{
		Title:        c.Keyword(),
		Autocomplete: c.Keyword(),
		Arg:          "token",
		SubtitleAll:  "Manually enter toggl.com API token",
	}
}

func (c TokenCommand) Items(prefix, query string) ([]alfred.Item, error) {
	return []alfred.Item{c.MenuItem()}, nil
}

func (c TokenCommand) Do(query string) (string, error) {

	//toggl
	btn, tgtoken, err := workflow.GetInput("Toggl API token", "", false)
	if err != nil {
		return "", err
	}

	if btn != "Ok" {
		log.Println("User didn't click OK")
		return "", nil
	}
	log.Printf("token: %s", tgtoken)

	config.ApiKey = tgtoken

	// trello

	btn, trappkey, err := workflow.GetInput("Trello App key", "", false)
	if err != nil {
		return "", err
	}

	if btn != "Ok" {
		log.Println("User didn't click OK")
		return "", nil
	}
	log.Printf("appkey: %s", trappkey)

	config.TrelloAppKey = trappkey

	btn, trtoken, err := workflow.GetInput("Trello Token key", "", false)
	if err != nil {
		return "", err
	}

	if btn != "Ok" {
		log.Println("User didn't click OK")
		return "", nil
	}
	log.Printf("appkey: %s", trtoken)

	config.TrelloAppKey = trtoken

	btn, trboardid, err := workflow.GetInput("Trello Board Id", "", false)
	if err != nil {
		return "", err
	}

	if btn != "Ok" {
		log.Println("User didn't click OK")
		return "", nil
	}
	log.Printf("appkey: %s", trboardid)

	config.TrelloBoardId = trboardid

	err = alfred.SaveJson(configFile, &config)
	if err != nil {
		return "", err
	}

	workflow.ShowMessage("Token saved!")
	return "", nil
}
