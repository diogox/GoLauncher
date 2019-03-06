package events

import "github.com/diogox/GoLauncher/api"

func PreferencesNew(preferences map[string]string) Preferences {
	return Preferences{
		Type:        api.PREFERENCES_EVENT,
		Preferences: preferences,
	}
}

type Preferences struct {
	Type        string
	Preferences map[string]string
}

func (p Preferences) GetEventType() string {
	return p.Type
}
