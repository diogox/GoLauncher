package gtk3

import (
	"github.com/gotk3/gotk3/gtk"
	"sync"
)

// TODO: Save an instance of the window with the launcher for performance? (Maybe we can get rid of the singleton)

var preferencesWindowInstance *gtk.Window
var mtx sync.Mutex

func ShowSettingsWindow() {
	mtx.Lock()
	defer mtx.Unlock()

	if preferencesWindowInstance == nil {
		preferencesWindowInstance = buildSettingsWindow()
	}

	preferencesWindowInstance.ShowAll()
}

func buildSettingsWindow() *gtk.Window {
	prefsWin, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	label, _ := gtk.LabelNew("Preferences Window!")
	prefsWin.Add(label)

	_, _ = prefsWin.Connect("destroy", func() {
		preferencesWindowInstance = nil
	})
	return prefsWin
}