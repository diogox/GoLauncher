package api

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
