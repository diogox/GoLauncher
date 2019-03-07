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

var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {

	time.Sleep(3000)
	action := api.Action(actions.NewOpenUrl("http://google.com"))
	event := api.Event(events.KeywordQueryNew("query"))
	response := api.ResponseNew(event, action)
	jsonObj, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonObj))

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
			fmt.Println(queryEvent.Query)

			results := make([]api.Result, 0)
			act := actions.NewOpenUrl("http://news.ycombinator.com/")
			results = append(results, result.NewSearchResult("Story 1", "Click here to read more about story 1!", "google", act, act))
			event := events.KeywordQueryNew(string(message))
			action := actions.RenderResultList{
				Type: api.RENDER_RESULT_LIST_ACTION,
				ResultList: results,
			}
			res := api.ResponseNew(event, action)
			resJson, err := json.Marshal(res)
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
	err = c.WriteMessage(websocket.TextMessage, jsonObj)
	if err != nil {
		log.Println("write:", err)
		return
	}
}
