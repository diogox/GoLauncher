package api

// Preferences
const PreferenceHotkey = "hotkey"
const PreferenceKeepInputOnHide = "keep_input_on_hide"

// For Booleans values
const PreferenceTRUE = "true"
const PreferenceFALSE = "false"

type Preferences interface {
	GetPreference(preference string) (string, error)
	SetPreference(preference string, value string) error
}
