package main

import (
	"log"

	"github.com/jason0x43/go-alfred"
)

type CreateAction struct{}

func (c CreateAction) IsEnabled() bool {
	return config.ApiKey != ""
}

func (c CreateAction) Keyword() string {
	return "create"
}

func (c CreateAction) Do(query string) (string, error) {
	log.Printf("create '%s'", query)

	// trello に追加のみ
	list, err := ta.findList(config.NewList)
	if err != nil {
		return "", err
	}
	card, err := ta.addCard(list, query)
	if err != nil {
		return "", err
	}
	if err == nil {
		log.Printf("Got entry: %#v\n", card)
		cache.Trello.Data.TimeEntries = append(cache.Trello.Data.TimeEntries, card)
		err := alfred.SaveJson(cacheFile, &cache)
		if err != nil {
			log.Printf("Error saving cache: %s\n", err)
		}
	}

	return "Created time entry", err
}
