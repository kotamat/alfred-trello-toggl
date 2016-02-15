package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/jason0x43/go-alfred"
	"github.com/jason0x43/go-toggl"
)

var cacheFile string
var configFile string
var config Config
var cache Cache
var workflow alfred.Workflow
var ta TrelloApi

type Config struct {
	ApiKey        string `json:"api_key"`
	TrelloAppKey  string `json:"trello_app_key"`
	TrelloToken   string `json:"trello_token"`
	TrelloBoardId string `json:"trello_board_id"`
	DurationOnly  bool   `desc:"Extend time entries instead of creating new ones."`
	Rounding      int    `desc:"Minutes to round to, 0 to disable rounding." help:"%v minute increments"`
	NewList      string
	ProgressList string
	DoneList     string
	BoardId      string
	TestMode      bool
}

func configtest(config *Config){
	config.NewList = "new"
	config.ProgressList = "progress"
	config.DoneList = "done"

	config.TrelloAppKey = "2dcc00794b413b4c51e03e71fe7897a7"
	config.TrelloToken = "446f130ff3aac387cb1607b451ed3a37e45f407d6ac38c41415a09d5f6f78bbb"
	config.BoardId = "FgLF9iac"
	return
}

type Cache struct {
	Workspace int
	Toggl     toggl.Account
	Trello    TrelloApi
	Time      time.Time
}

func main() {
	var err error

	workflow, err = alfred.OpenWorkflow(".", true)
	if err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(1)
	}

	toggl.AppName = "alfred-toggl"

	configFile = path.Join(workflow.DataDir(), "config.json")
	cacheFile = path.Join(workflow.CacheDir(), "cache.json")

	log.Println("Using config file", configFile)
	log.Println("Using cache file", cacheFile)

	configtest(&config)
	ta.setup()
	err = alfred.LoadJson(configFile, &config)
	if err != nil {
		log.Println("Error loading config:", err)
	}
	log.Println("loaded config:", config)

	err = alfred.LoadJson(cacheFile, &cache)
	log.Println("loaded cache")

	commands := []alfred.Command{
		//LoginCommand{},
		TokenCommand{},
		TaskFilter{},
		CreateAction{},
		ToggleAction{},
		SyncFilter{},
//		TimeEntryFilter{},
//		ProjectFilter{},
//		TagFilter{},
//		ReportFilter{},
//		OptionsCommand{},
//		LogoutCommand{},
//		ResetCommand{},
//		StartAction{},
//		UpdateTimeEntryAction{},
//		UpdateProjectAction{},
//		CreateProjectAction{},
//		UpdateTagAction{},
//		CreateTagAction{},
//		DeleteAction{},
	}

	workflow.Run(commands)
}
