package main

import (

	"log"

	"github.com/jason0x43/go-alfred"
	"github.com/jason0x43/go-toggl"
)

// ToggleAction toggles a time entry's running state.
type ToggleAction struct{}

// Keyword return's the action's keyword.
func (c ToggleAction) Keyword() string {
	return "toggle"
}

func (c ToggleAction) IsEnabled() bool {
	return config.ApiKey != ""
}

func (c ToggleAction) Do(query string) (string, error) {
	log.Printf("doToggle(%s)", query)

	card, err := ta.findCardById(query)
	if err != nil {
		return "", err
	}

	adata := &cache.Toggl.Data
	session := toggl.OpenSession(config.ApiKey)
	running, isRunning := getRunningTimer()
	isNew := true

	entry, ok := ta.getEntry(*card)
	var operation string
	var updatedEntry toggl.TimeEntry
	var moveToListName string

	// toggle
	if !ok {
		operation = "Start"
		updatedEntry, err = session.StartTimeEntry(card.Name)
		moveToListName = config.ProgressList
	} else if entry.IsRunning() {
		// two p's so we get "Stopped"
		operation = "Stopp"
		updatedEntry, err = session.StopTimeEntry(entry)
		moveToListName = config.DoneList
	} else {
		operation = "Restart"
		updatedEntry, err = session.ContinueTimeEntry(entry, config.DurationOnly)
		moveToListName = config.ProgressList
	}
	// trello 移動
	moveList, err := ta.findList(moveToListName)
	if err != nil {
		return "", err
	}
	_, err = ta.moveCard(*card, moveList)
	if err != nil {
		return "", err
	}

	for i := 0; i < len(adata.TimeEntries); i++ {
		if adata.TimeEntries[i].Id == updatedEntry.Id{
			adata.TimeEntries[i] = updatedEntry
			isNew = false;
		}
	}

	if isNew {
		adata.TimeEntries = append(adata.TimeEntries, updatedEntry)
	}

	if isRunning && running.Id != updatedEntry.Id {
		// If a different timer was previously running, refresh everything
		err = refresh()
	} else {
		err = alfred.SaveJson(cacheFile, &cache)
	}

	if err != nil {
		log.Printf("Error saving cache: %v\n", err)
	}

	return operation + "ed " + entry.Description, nil
}
