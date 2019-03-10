package api

import "encoding/json"

func ResponseNew(event Event, action Action) Response {
	return Response{
		Event:  event,
		Action: action,
	}
}

type Response struct {
	Event  Event
	Action Action
}

func (r *Response) MarshalJSON() ([]byte, error) {

	action, err := json.Marshal(r.Action)
	if err != nil {
		panic(err)
	}

	return json.Marshal(map[string]string {
		"Action": string(action),
	})
}
