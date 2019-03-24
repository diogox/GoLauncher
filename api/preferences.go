package api

// Preferences
const PreferenceHotkey = "hotkey"
const PreferenceKeepInputOnHide = "keep_input_on_hide"
const PreferenceLaunchAtStartUp = "launch_at_startup"
const PreferenceNResultsToShow = "number_of_results_to_show"

// For Booleans values
const PreferenceTRUE = "true"
const PreferenceFALSE = "false"

func AssertPreferenceBool(value string) bool {
	if value == PreferenceTRUE {
		return true
	} else if value == PreferenceFALSE {
		return false
	}
	panic("invalid preference boolean")
}

type Preferences interface {
	GetPreference(preference string) (string, error)
	SetPreference(preference string, value string) error
	BindCallback(preference string, callback func(arg interface{}))
}
