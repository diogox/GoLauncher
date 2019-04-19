package gtk3

import (
	"errors"
	"github.com/diogox/GoLauncher/api"
	"github.com/diogox/GoLauncher/pkg/autostart"
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
		p.runAllCallbacks(api.PreferenceHotkey, value)

	case api.PreferenceKeepInputOnHide:
		p.runAllCallbacks(api.PreferenceKeepInputOnHide, nil)

	case api.PreferenceLaunchAtStartUp:
		isStart := api.AssertPreferenceBool(value)
		autostart.SetAppStart(isStart)

		// Run all associated callbacks
		p.runAllCallbacks(api.PreferenceLaunchAtStartUp, nil)

	case api.PreferenceNResultsToShow:
		// Make sure it's an interface
		_, err := strconv.Atoi(value)
		if err != nil {
			// Keep old value
			value = (*p.db).GetPreference(preference)
		}

		// Run all associated callbacks
		p.runAllCallbacks(api.PreferenceNResultsToShow, value)

	case api.PreferenceNAppResults:
		// Make sure it's an interface
		_, err := strconv.Atoi(value)
		if err != nil {
			// Keep old value
			value = (*p.db).GetPreference(preference)
		}

		// Run all associated callbacks
		p.runAllCallbacks(api.PreferenceNAppResults, value)

	case api.PreferenceShowFrequentApps:
		shouldShow := api.AssertPreferenceBool(value)

		// Run all associated callbacks
		p.runAllCallbacks(api.PreferenceShowFrequentApps, shouldShow)
	}
	return (*p.db).SetPreference(preference, value)
}

func (p *Preferences) BindCallback(preference string, callback func(arg interface{})) {
	p.callbacks[preference] = append(p.callbacks[preference], callback)
}

func (p *Preferences) runAllCallbacks(preference string, arg interface{}) {
	for _, c := range p.callbacks[preference] {
		c(arg)
	}
}