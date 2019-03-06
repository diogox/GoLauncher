package events

import "github.com/diogox/GoLauncher/api"

func PreferencesUpdateNew(preference string, oldValue string, newValue string) PreferencesUpdate {
	return PreferencesUpdate{
		Type:       api.PREFERENCES_UPDATE_EVENT,
		Preference: preference,
		OldValue:   oldValue,
		NewValue:   newValue,
	}
}

type PreferencesUpdate struct {
	Type       string
	Preference string
	OldValue   string
	NewValue   string
}

func (pu PreferencesUpdate) GetEventType() string {
	return pu.Type
}
