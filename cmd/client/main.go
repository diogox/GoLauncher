package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/diogox/GoLauncher/api"
	"github.com/diogox/GoLauncher/api/actions"
	"github.com/diogox/GoLauncher/api/events"
	"github.com/diogox/GoLauncher/search/result"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"
)

var addr = flag.String("addr", "localhost:8081", "http service address")

func main() {

	time.Sleep(3000)
	/*
	action := api.Action(actions.NewOpenUrl("http://google.com"))
	event := api.Event(events.KeywordQueryNew("query"))
	response := api.ResponseNew(event, action)
	jsonObj, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	*/

	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}

			var queryEvent events.KeywordQuery
			err = json.Unmarshal(message, &queryEvent)
			if err != nil {
				panic(err)
			}
			//	fmt.Println(queryEvent.Query)

			actions.SetupRenderResultList(func(results []api.SearchResult) error {
				fmt.Println("Sending: ", results)
				return nil
			})

			results := make([]api.SearchResult, 0)

			opts := result.SearchResultOptions{
				Title: "Story 1",
				Description: "Click here to read more about story 1!",
				IconPath: "google",
				IsDefaultSelect: false,
				OnEnterAction: actions.NewOpenUrl("http://news.ycombinator.com/"),
				OnAltEnterAction: actions.NewOpenUrl("http://news.ycombinator.com/"),
			}

			results = append(results, result.NewSearchResult(opts))
			event := events.KeywordQueryNew(string(message))
			action := actions.NewRenderResultList(results)
			res := api.ResponseNew(event, action)
			resJson, err := json.Marshal(&res)
			fmt.Println(string(resJson))
			if err != nil {
				panic(err)
			}
			err = c.WriteMessage(websocket.TextMessage, resJson)
			if err != nil {
				panic(err)
			}
		}
	}()

	for true {}
	//err = c.WriteMessage(websocket.TextMessage, jsonObj)
	if err != nil {
		log.Println("write:", err)
		return
	}
}
