package websockets

import (
	"encoding/json"
	"fmt"
	"github.com/diogox/GoLauncher/api"
	json2 "github.com/diogox/GoLauncher/websockets/json"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func StartExtensionsServer(responseChannel chan *api.Response) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

		for {
			// Read message from browser
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}

			// Infer action type
			var res map[string]*json.RawMessage
			fmt.Println(string(msg))
			err = json.Unmarshal(msg, &res)
			fmt.Println(res)
			if err == nil {
			}

			// Unmarshal action
			actionJson, err := res["Action"].MarshalJSON()
			if err != nil {
				panic(err)
			}
			action, err := json2.InferActionType(actionJson)
			if err != nil {
				panic(err)
			}

			// Unmarshal event
			eventJson, err := res["Event"].MarshalJSON()
			if err != nil {
				panic(err)
			}
			event, err := json2.InferEventType(eventJson)
			if err != nil {
				panic(err)
			}

			response := api.Response{
				Action: *action,
				Event:  *event,
			}

			// Send response
			responseChannel <- &response

			// Print the message to the console
			//fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), (*action).GetType())

			// Write message back to browser
			if err = conn.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
	})

	http.ListenAndServe(":8080", nil)
}
