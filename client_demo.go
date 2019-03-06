package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/diogox/GoLauncher/api"
	"github.com/diogox/GoLauncher/api/actions"
	"github.com/diogox/GoLauncher/api/events"
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

	/*
	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()
	*/

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	err = c.WriteMessage(websocket.TextMessage, jsonObj)
	if err != nil {
		log.Println("write:", err)
		return
	}
	/*
	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
	*/
}
