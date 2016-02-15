package main
import (
	"github.com/jason0x43/go-alfred"
	"strings"
	"fmt"
)

type TaskFilter struct {}

//var ta TrelloApi
//var config Config
//
//type Config struct {
//	NewList      string
//	ProgressList string
//	DoneList     string
//	TrelloAppKey string
//	TrelloToken  string
//	BoardId      string
//}
//
//func (ta *TrelloApi) configtest(config *Config){
//	config.NewList = "new"
//	config.ProgressList = "progress"
//	config.DoneList = "done"
//
//	config.TrelloAppKey = "2dcc00794b413b4c51e03e71fe7897a7"
//	config.TrelloToken = "446f130ff3aac387cb1607b451ed3a37e45f407d6ac38c41415a09d5f6f78bbb"
//	config.BoardId = "FgLF9iac"
//	return
//}

func (c TaskFilter) Keyword() string {
	return "tasks"
}

func (c TaskFilter) IsEnabled() bool {
	return config.TrelloAppKey != ""
}

func (c TaskFilter) MenuItem() alfred.Item {
	return alfred.Item{
		Title:        c.Keyword(),
		Autocomplete: c.Keyword() + " ",
		Valid:        alfred.Invalid,
		SubtitleAll:  "List and modify recent time entries, add new ones",
	}
}

func (c TaskFilter) Items(prefix, query string) (items []alfred.Item, err error) {
	parts := alfred.TrimAllLeft(strings.Split(query, alfred.Separator))
//	addItem := func(title, subtitle, keyword string, hasNext, showSubtitle bool) {
//		item := alfred.Item{
//			Title:        title,
//			Autocomplete: prefix + parts[0] + alfred.Separator + " " + keyword,
//		}
//
//		if showSubtitle {
//			item.Subtitle = subtitle
//		}
//
//		if hasNext {
//			item.Autocomplete += alfred.Separator + " "
//		}
//
//		if len(parts) > 2 {
//			item.Arg = parts[2]
//		} else {
//			item.Valid = alfred.Invalid
//		}
//
//		items = append(items, item)
//	}
	cards := getCardsForQuery(parts[0])
	if len(cards) == 0 {
		items = append(items, alfred.Item{
			Title:       parts[0],
			SubtitleAll: "New entry",
			Arg:         "create " + parts[0],
		})
	} else {
		for _, card := range cards {
			subtitle := "toggle timer"
			item := alfred.Item{
				Title:        card.Name,
				SubtitleAll:  subtitle,
				Arg:          fmt.Sprintf("toggle %v", card.Id),
				Autocomplete: prefix + fmt.Sprintf("%d%s ",card.Id, alfred.Separator),
			}

			if ta.IsRunning(card) {
				item.Icon = "running.png"
			}

			items = append(items, item)
		}
	}

	return
}
