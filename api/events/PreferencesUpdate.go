package events

import (
	"github.com/diogox/GoLauncher/api"
)

func PreferencesUpdateNew(preference string, oldValue string, newValue string) PreferencesUpdate {
	return PreferencesUpdate{
		Preference: preference,
		OldValue:   oldValue,
		NewValue:   newValue,
	}
}

type PreferencesUpdate struct {
	api.BaseEvent
	Preference string
	OldValue   string
	NewValue   string
}
