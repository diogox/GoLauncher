package gtk3

import (
	kb "github.com/Isolus/go-keybinder"
	"github.com/diogox/GoLauncher/gtk3/glade"
	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"unsafe"
)

const GladeFile = "/home/diogox/go/src/github.com/diogox/GoLauncher/assets/launcher.glade"
const CssFile = "/home/diogox/go/src/github.com/diogox/GoLauncher/assets/theme.css"

const WindowName = "window"
const BodyName = "body"
const InputBoxName = "input-box"
const InputName = "input"
const PrefsBtnName = "prefs_btn"
const ResultsBoxName = "result_box"

func NewLauncher() Launcher {

	// Initiate gtk (Must be here, so that it occurs before the hotkey binding)
	gtk.Init(nil)

	// Build from glade file
	bldr, err := glade.BuildFromFile(GladeFile)
	if err != nil {
		panic(err)
	}

	// Get CSS provider
	cssProvider, err := setCssProvider()
	if err != nil {
		panic(err)
	}

	// Get window
	win, err := glade.GetWindow(bldr, WindowName)
	if err != nil {
		panic("Failed to get Window: " + err.Error())
	}

	// Get body
	body, err := glade.GetBox(bldr, BodyName)
	if err != nil {
		panic("Failed to get Body: " + err.Error())
	}

	// Set body style
	setStyleClass(cssProvider, &body.Widget, "app")

	// Get input
	input, err := glade.GetEntry(bldr, InputName)
	if err != nil {
		panic("Failed to get Input: " + err.Error())
	}

	// Set input style
	setStyleClass(cssProvider, &input.Widget, "input")

	// Get preferences button
	prefsBtn, err := glade.GetButton(bldr, PrefsBtnName)
	if err != nil {
		panic("Failed to get Input: " + err.Error())
	}

	// Set preferences button style
	setStyleClass(cssProvider, &prefsBtn.Widget, "prefs-btn")


	// TODO: Remove this!
	resultsBox, err := glade.GetBox(bldr, ResultsBoxName)
	if err != nil {
		panic("Failed to get Input: " + err.Error())
	}

	res1 := NewResultItem("Item1", "Description for item 1.")
	res2 := NewResultItem("Item2", "Description for item 2.")
	res3 := NewResultItem("Item3", "Description for item 3.")
	resultsBox.PackEnd(res1.frame, true, true, 2)
	resultsBox.PackEnd(res2.frame, true, true, 2)
	resultsBox.PackEnd(res3.frame, true, true, 2)

	return Launcher{
		window:    win,
		body:      body,
		input:     input,
		prefsBtn:  prefsBtn,
		isVisible: true,
	}
}

type Launcher struct {
	window    *gtk.Window
	body      *gtk.Box
	input     *gtk.Entry
	prefsBtn  *gtk.Button
	isVisible bool
}

func (l *Launcher) Start() {

	// Keep the launcher above everything
	l.window.SetKeepAbove(true)

	// Set the groundwork for transparency
	screenChanged(l.window)

	// When the monitor/screen changes, check for transparency support
	_, _ = l.window.Connect("screen-changed", func(window *gtk.Window, oldScreen *gdk.Screen, userData ...interface{}) {
		screenChanged(window)
	})

	// Set transparency on draw
	_, _ = l.window.Connect("draw", func(window *gtk.Window, context *cairo.Context) {
		setTransparent(window, context)
	})

	// When the window loses focus, hide it
	_, _ = l.window.Connect("focus-out-event", func(widget *gtk.Window, event *gdk.Event) {
		_, _ = glib.IdleAdd(l.hide)
	})

	// Show everything in the app
	l.show()

	// Start loop
	go gtk.Main()
}

func (l *Launcher) Stop() {

	gtk.MainQuit()
	// TODO: Will it exit? WaitGroup is still waiting probably...
}

func (l *Launcher) BindHotkey(hotkey string) {

	// Create handler
	toggle := func(keystring string, data unsafe.Pointer) {
		l.ToggleVisibility()
	}

	// Bind hotkey
	kb.Init()
	kb.Bind(hotkey, toggle, nil)
}

func (l *Launcher) ToggleVisibility() {

	if l.isVisible {

		l.hide()
	} else {

		l.show()
	}
}

func (l *Launcher) ClearInput() {
	l.input.SetText("")
}

func (l *Launcher) ShowResults([]ResultItem) {

	// TODO
}

func (l *Launcher) show() {

	// Position
	centerAtTopOfScreen(l.window)

	// Clear
	l.ClearInput()

	// Show
	l.window.ShowAll()
	l.isVisible = true

	// Focus
	l.window.PresentWithTime(kb.GetCurrentEventTime())
}

func (l *Launcher) hide() {

	// Hide
	l.window.Hide()
	l.isVisible = false
}
