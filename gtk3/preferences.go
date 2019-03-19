package gtk3

import (
	"errors"
	"github.com/diogox/GoLauncher/api"
	"github.com/diogox/GoLauncher/autostart"
)

func PreferencesNew(db *api.DB) Preferences {
	// Load default preferences if they aren't already set
	_ = (*db).LoadDefaultPreferences()
	
	return Preferences{
		db: db,
	}
}

type Preferences struct {
	db *api.DB
	bindHotkeyCallback func(string)
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
	}
	return (*p.db).SetPreference(preference, value)
}

func (p *Preferences) BindHotkeyCallBack(callback func(string)) {
	p.bindHotkeyCallback = callback
}