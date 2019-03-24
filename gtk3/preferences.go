package gtk3

import (
	"errors"
	"github.com/diogox/GoLauncher/api"
	"github.com/diogox/GoLauncher/autostart"
	"strconv"
)

func PreferencesNew(db *api.DB) Preferences {
	// Load default preferences if they aren't already set
	_ = (*db).LoadDefaultPreferences()

	return Preferences{
		db: db,
		callbacks: make(map[string][]func(interface{}), 0),
	}
}

type Preferences struct {
	db                 *api.DB
	bindHotkeyCallback func(string)
	callbacks          map[string][]func(interface{})
}

func (p *Preferences) GetPreference(preference string) (string, error) {
	preferenceValue := (*p.db).GetPreference(preference)
	if preferenceValue == "" {
		return "", errors.New("preference doesn't exist")
	}

	return preferenceValue, nil
}

func (p *Preferences) SetPreference(preference string, value string) error {
	switch preference {
	case api.PreferenceHotkey:
		p.bindHotkeyCallback(value)
	case api.PreferenceKeepInputOnHide:
		// TODO: Maybe not needed?
	case api.PreferenceLaunchAtStartUp:
		isStart := api.AssertPreferenceBool(value)
		autostart.SetAppStart(isStart)
	case api.PreferenceNResultsToShow:
		// Make sure it's an int
		_, err := strconv.Atoi(value)
		if err != nil {
			// Keep old value
			value = (*p.db).GetPreference(preference)
		}

		// Run all associated callbacks
		for _, c := range p.callbacks[api.PreferenceNResultsToShow] {
			nOfResults, _ := strconv.Atoi(value)
			c(nOfResults)
		}
	}
	return (*p.db).SetPreference(preference, value)
}

func (p *Preferences) BindHotkeyCallBack(callback func(string)) {
	p.bindHotkeyCallback = callback
}

func (p *Preferences) BindCallback(preference string, callback func(arg interface{})) {
	p.callbacks[preference] = append(p.callbacks[preference], callback)
}
