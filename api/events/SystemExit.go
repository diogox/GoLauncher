package events

import "github.com/diogox/GoLauncher/api"

func SystemExitNew() SystemExit {
	return SystemExit {}
}

type SystemExit struct {
	api.BaseEvent
}
