package gtk3

import (
	"errors"
	"github.com/diogox/GoLauncher/api"
)

func PreferencesNew(db *api.DB) Preferences {
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
	}
	return (*p.db).SetPreference(preference, value)
}

func (p *Preferences) BindHotkeyCallBack(callback func(string)) {
	p.bindHotkeyCallback = callback
}