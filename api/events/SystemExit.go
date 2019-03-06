package events

import "github.com/diogox/GoLauncher/api"

func SystemExitNew() SystemExit {
	return SystemExit{
		Type: api.SYSTEM_EXIT_EVENT,
	}
}

type SystemExit struct {
	Type string
}

func (se SystemExit) GetEventType() string {
	return se.Type
}
