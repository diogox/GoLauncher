package gtk3

import (
	"github.com/diogox/GoLauncher/api"
	"github.com/diogox/GoLauncher/gtk3/glade"
	"github.com/gotk3/gotk3/gtk"
	"sync"
)

const GladeSettingsFile = "/home/diogox/go/src/github.com/diogox/GoLauncher/gtk3/assets/settings.glade"

const SettingsWindowID = "settings_window"
const SettingsHotkeyInputID = "hotkey_input"
const SettingsClearInputOnHideID = "keep_input_on_hide"
const SettingsSaveButtonID = "save"
const SettingsCancelButtonID = "cancel"

var settingsWindowInstance *gtk.Window
var mtx sync.Mutex

func ShowSettingsWindow(preferences *api.Preferences) {
	mtx.Lock()
	defer mtx.Unlock()

	if settingsWindowInstance == nil {
		settingsWindowInstance = buildSettingsWindow(preferences)
	}

	settingsWindowInstance.ShowAll()
}

func buildSettingsWindow(preferences *api.Preferences) *gtk.Window {
	bldr, err := glade.BuildFromFile(GladeSettingsFile)
	if err != nil {
		panic(err)
	}

	win, err :=  glade.GetWindow(bldr, SettingsWindowID)
	if err != nil {
		panic(err)
	}

	hotkeyInput, err :=  glade.GetEntry(bldr, SettingsHotkeyInputID)
	if err != nil {
		panic(err)
	}

	keepInputOnHideCheckButton, err :=  glade.GetCheckButton(bldr, SettingsClearInputOnHideID)
	if err != nil {
		panic(err)
	}

	saveButton, err :=  glade.GetButton(bldr, SettingsSaveButtonID)
	if err != nil {
		panic(err)
	}

	cancelButton, err :=  glade.GetButton(bldr, SettingsCancelButtonID)
	if err != nil {
		panic(err)
	}

	// Load with current preference values
	isKeepInputOnHide, err := (*preferences).GetPreference(api.PreferenceKeepInputOnHide)
	if err != nil {
		panic(err)
	}
	if isKeepInputOnHide == api.PreferenceTRUE {
		keepInputOnHideCheckButton.SetActive(true)
	} else {
		keepInputOnHideCheckButton.SetActive(false)
	}

	currentHotkey, err := (*preferences).GetPreference(api.PreferenceHotkey)
	if err != nil {
		panic(err)
	}
	hotkeyInput.SetText(currentHotkey)

	// Save new preferences
	_, _ = saveButton.Connect("clicked", func(btn *gtk.Button) {
		newHotkey, err := hotkeyInput.GetText()
		if err != nil {
			panic(err)
		}
		err = (*preferences).SetPreference(api.PreferenceHotkey, newHotkey)
		if err != nil {
			panic(err)
		}

		isKeepInputOnHide := keepInputOnHideCheckButton.GetActive()
		if isKeepInputOnHide {
			err = (*preferences).SetPreference(api.PreferenceKeepInputOnHide, api.PreferenceTRUE)
		} else {
			err = (*preferences).SetPreference(api.PreferenceKeepInputOnHide, api.PreferenceFALSE)
		}
		if err != nil {
			panic(err)
		}

		win.Destroy()
	})

	_, _ = cancelButton.Connect("clicked", func() {
		win.Destroy()
	})

	_, _ = win.Connect("destroy", func() {
		settingsWindowInstance = nil
	})
	return win
}