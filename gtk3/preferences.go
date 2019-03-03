package gtk3

import (
	"github.com/gotk3/gotk3/gtk"
	"sync"
)

var preferencesWindowInstance *gtk.Window
var mtx sync.Mutex

func ShowPreferencesWindow() {
	mtx.Lock()
	defer mtx.Unlock()

	if preferencesWindowInstance == nil {
		preferencesWindowInstance = buildPreferencesWindow()
	}

	preferencesWindowInstance.ShowAll()
}

type PreferencesWindow struct {

}

func buildPreferencesWindow() *gtk.Window {
	prefsWin, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	label, _ := gtk.LabelNew("Preferences Window!")
	prefsWin.Add(label)

	_, _ = prefsWin.Connect("destroy", func() {
		preferencesWindowInstance = nil
	})
	return prefsWin
}