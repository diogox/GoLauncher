package websockets

import (
	"encoding/json"
	"fmt"
	"github.com/diogox/GoLauncher/api"
	"github.com/diogox/GoLauncher/api/actions"
	"github.com/diogox/GoLauncher/api/events"
	"github.com/diogox/GoLauncher/api/models"
	json2 "github.com/diogox/GoLauncher/websockets/json"
	"github.com/gorilla/websocket"
)

func ControllerNew(extension *models.Extension, conn *websocket.Conn) *Controller {
	return &Controller{
		Extension: extension,
		conn:      conn,
	}
}

type Controller struct {
	Extension *models.Extension
	conn      *websocket.Conn
}

func (c *Controller) HandleInput(args string) api.Action {

	// Send Event
	event := api.Event(events.KeywordQueryNew(args))
	err := c.Send(&event)
	if err != nil {
		panic(err)
	}

	// Read Response
	res, err := c.Receive()
	if err == nil {
		return res.Action
	}

	return actions.NewDoNothing()
}

func (c *Controller) Send(res *api.Event) error {

	msg, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}

	// Write message back to browser
	if err := c.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
		return err
	}

	return nil
}

func (c *Controller) Receive() (*api.Response, error) {


	_, message, err := c.conn.ReadMessage()
	if err != nil {
		return nil, err
	}

	// Infer action type
	var res map[string]*json.RawMessage
	fmt.Println(string(message))
	err = json.Unmarshal(message, &res)
	fmt.Println(res)
	if err != nil {
		return nil, err
	}

	// Unmarshal action
	actionJson, err := res["Action"].MarshalJSON()
	if err != nil {
		return nil, err
	}
	action, err := json2.InferActionType(actionJson)
	if err != nil {
		return nil, err
	}

	// Unmarshal event
	eventJson, err := res["Event"].MarshalJSON()
	if err != nil {
		return nil, err
	}
	event, err := json2.InferEventType(eventJson)
	if err != nil {
		return nil, err
	}

	// TODO: !!!! CAN'T UNMARSHAL RESULT_LIST INSIDE OF ACTION!!!! Shows as empty (Problem is in the client!!)

	response := api.Response{
		Action: *action,
		Event:  *event,
	}

	return &response, nil
}