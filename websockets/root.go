package websockets

import (
	"github.com/diogox/GoLauncher/api"
	"github.com/diogox/GoLauncher/websockets/json"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func StartExtensionsServer(actionChannel chan *api.Action) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

		for {
			// Read message from browser
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}

			// Infer action type
			action, err := json.InferActionType(msg)
			if err != nil {
				panic(err)
			}

			// Send action
			actionChannel <- action

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
