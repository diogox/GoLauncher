package events

import "github.com/diogox/GoLauncher/api"

func PreferencesNew(preferences map[string]string) Preferences {
	return Preferences{
		Preferences: preferences,
	}
}

type Preferences struct {
	api.BaseEvent
	Preferences map[string]string
}

