package main

import (

	"github.com/VojtechVitek/go-trello"
	"net/url"
	"encoding/json"
	"log"
	"errors"
	"github.com/jason0x43/go-toggl"
)

type TrelloApi struct {
	client *trello.Client
	config Config
	Data struct {
			 TimeEntries []trello.Card `json:"time_entries"`
		 } `json:"data"`
}

//var ta TrelloApi

//func main() {
//	err := ta.setup()
//	if err != nil {
//		return
//	}
//	list, err := ta.findList(NewList)
//	if err != nil {
//		return
//	}
//	card, err := ta.addCard(list, "newcard")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	moveList, err := ta.findList(ProgressList)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	movedCard, err := ta.moveCard(card, moveList)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println("add card : ", movedCard.IdList)
//
//	 //findCard
//	cardName := "test3"
//	card, err := ta.findCard(cardName)
//	if err == nil {
//		fmt.Println("list ", card.IdList)
//		fmt.Println("is? ", ta.isExist(cardName))
//	} else {
//		fmt.Println("err", err)
//		fmt.Println("is? ", ta.isExist(cardName))
//	}
//}

func (ta *TrelloApi) setup() (err error){
	appKey := config.TrelloAppKey
	token := config.TrelloToken
	client, err := trello.NewAuthClient(appKey, &token)
	if err != nil {
		log.Fatal(err)
	}
	ta.client = client
	return
}

func (ta *TrelloApi) findList(name string) (list trello.List, err error) {
	board, err := ta.client.Board(config.BoardId)
	lists, err := board.Lists()
	if err != nil{
		return
	}

	for _, list := range lists  {
		if list.Name == name {
			return list, err
		}
	}
	return
}

func (ta *TrelloApi) addCard(list trello.List, name string) (card trello.Card, err error) {
	payload := url.Values{}
	payload.Set("idList", list.Id)
	payload.Set("name", name)

	body, err := ta.client.Post("/cards/", payload)

	err = json.Unmarshal(body, &card)
	return
}

func (ta *TrelloApi) moveCard(card trello.Card, list trello.List) (movedCard trello.Card, err error) {
	payload := url.Values{}
	payload.Set("idList", list.Id)

	body, err := ta.client.Put("/cards/"+card.Id, payload)

	err = json.Unmarshal(body, &movedCard)
	return
}

func (ta *TrelloApi) findCardById(cardId string) (card *trello.Card, err error) {
	return ta.client.Card(cardId)
}

func (ta *TrelloApi) findCard(cardName string) (card trello.Card, err error) {
	board, err := ta.client.Board(config.BoardId)
	lists, err := board.Lists()
	if err != nil {
		return
	}

	for _, list := range lists {
		card, err = ta.findCardInList(list, cardName)
		if err == nil {
			return
		}
	}
	err = errors.New("no cards")
	return
}

func (ta *TrelloApi) findCardInList(list trello.List, cardName string) (card trello.Card, err error) {
	cards, err := list.Cards()
	if err != nil {
		return
	}

	for _, card := range cards {
		if card.Name == cardName {
			return card, err
		}
	}
	err = errors.New("no cards")
	return
}

func (ta *TrelloApi) IsRunning(card trello.Card) bool {
	timer, ok := findTimerByName(card.Name)
	if ok {
		return timer.IsRunning()
	}
	return false
}

func (ta *TrelloApi) getEntry(card trello.Card) (toggl.TimeEntry, bool) {
	return findTimerByName(card.Name)
}

func (ta *TrelloApi) isExist(cardName string) bool {
	_, err := ta.findCard(cardName)
	return err == nil
}

func (ta *TrelloApi) isExistInList(list trello.List, cardName string) bool {
	_, err := ta.findCardInList(list, cardName)
	return err == nil
}

func (ta *TrelloApi) getTogglFromCard(card trello.Card)  {
}

func (ta *TrelloApi) GetAccount() error {
	board, err := ta.client.Board(config.BoardId)
	cards, err := board.Cards()
	if err != nil {
		return err
	}
	ta.Data.TimeEntries = cards
	return err
}

