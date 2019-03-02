package gtk3

import (
	kb "github.com/Isolus/go-keybinder"
	"github.com/diogox/GoLauncher/common"
	"github.com/diogox/GoLauncher/common/actions"
	"github.com/diogox/GoLauncher/gtk3/glade"
	"github.com/diogox/GoLauncher/navigation"
	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"unsafe"
)

const GladeFile = "/home/diogox/go/src/github.com/diogox/GoLauncher/gtk3/assets/launcher.glade"
const CssFile = "/home/diogox/go/src/github.com/diogox/GoLauncher/gtk3/assets/theme.css"

const WindowID = "window"
const BodyID = "body"
const InputBoxID = "input-box"
const InputID = "input"
const PrefsBtnID = "prefs_btn"
const ResultsBoxID = "result_box"

func NewLauncher() *Launcher {

	// Initiate gtk (Must be here, so that it occurs before the hotkey binding)
	gtk.Init(nil)

	// Build from glade file
	bldr, err := glade.BuildFromFile(GladeFile)
	if err != nil {
		panic(err)
	}

	// Get CSS provider
	_, err = setCssProvider()
	if err != nil {
		panic(err)
	}

	// Get window
	win, err := glade.GetWindow(bldr, WindowID)
	if err != nil {
		panic("Failed to get Window: " + err.Error())
	}

	// Get body
	body, err := glade.GetBox(bldr, BodyID)
	if err != nil {
		panic("Failed to get Body: " + err.Error())
	}

	// Get input
	input, err := glade.GetEntry(bldr, InputID)
	if err != nil {
		panic("Failed to get Input: " + err.Error())
	}

	// Get preferences button
	prefsBtn, err := glade.GetButton(bldr, PrefsBtnID)
	if err != nil {
		panic("Failed to get Input: " + err.Error())
	}

	resultsBox, err := glade.GetBox(bldr, ResultsBoxID)
	if err != nil {
		panic("Failed to get Input: " + err.Error())
	}

	nav := navigation.NewNavigation(make([]*common.Result, 0))

	return &Launcher{
		window:     win,
		body:       body,
		input:      input,
		prefsBtn:   prefsBtn,
		resultsBox: resultsBox,
		navigation: &nav,
		isVisible:  true,
	}
}

type Launcher struct {
	window     *gtk.Window
	body       *gtk.Box
	input      *gtk.Entry
	prefsBtn   *gtk.Button
	resultsBox *gtk.Box
	navigation *navigation.Navigation
	isVisible  bool
}

func (l *Launcher) HandleInput(callback func(string)) {
	_, _ = l.input.Connect("changed", func(entry *gtk.Entry) {
		_, _ = glib.IdleAdd(func() {
			input, err := entry.GetText()
			if err != nil {
				panic(err)
			}

			// TODO: Move this to external logic? (To main.go?)
			if input == "" {
				l.clearResults()
				return
			}

			callback(input)
		})
	})
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

	// Detect navigation ('Up', 'Down', 'Enter')
	_, _ = l.window.Connect("key-press-event", func(widget *gtk.Window, event *gdk.Event) {
		keyEvent := &gdk.EventKey{
			Event: event,
		}

		var result, prevResult *common.Result

		// Resolve action
		const KEY_Enter = 65293
		switch keyEvent.KeyVal() {
		case gdk.KEY_Up:
			result, prevResult = l.navigation.Up()
			if result == nil {
				return
			}
		case gdk.KEY_Down:
			result, prevResult = l.navigation.Down()
			if result == nil {
				return
			}
		case KEY_Enter:
			if keyEvent.State() == gdk.GDK_MOD1_MASK {
				l.navigation.AltEnter()
				return
			}
			l.navigation.Enter()
			return
		default:
			return
		}

		res, ok := (*result).(*ResultItem)
		if !ok {
			panic("Error in navigation logic!")
		}
		prevRes, ok := (*prevResult).(*ResultItem)
		if !ok {
			panic("Error in navigation logic!")
		}

		prevRes.Unselect()
		res.Select()
	})

	l.navigation.SetOnItemEnter(func(action common.Action) {
		action.Run()
		if !action.KeepAppOpen() {
			l.hide()
		}
	})

	// Setup actions
	actions.SetupCopyToClipboard(func(text string) {
		_, _ = glib.IdleAdd(func(text string) {
			clipboard, err := gtk.ClipboardGet(gdk.SELECTION_CLIPBOARD)
			if err != nil {
				panic("Failed to get clipboard!")
			}

			clipboard.SetText(text)
			clipboard.Store()
		}, text)
	})
	actions.SetupSetUserQuery(func(query string) {
		_, _ = glib.IdleAdd(func(query string) {
			l.input.SetText(query)
		}, query)
	})
	actions.SetupRenderResultList(func(results []common.Result) {
		_, _ = glib.IdleAdd(func(results []common.Result) {
			l.ShowResults(results)
		}, results)
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

func (l *Launcher) ShowResults(searchResults []common.Result) {

	results := make([]*ResultItem, 0)

	// Convert results
	for i, r := range searchResults {
		result := NewResultItem(r.Title(), r.Description(), r.IconPath(), i+1, r.OnEnterAction(), r.OnAltEnterAction())
		result.BindMouseHover(func() {
			_, _ = glib.IdleAdd(func() {
				res := common.Result(&result)
				prevSelected := l.navigation.SetSelected(&res)
				prevRes, _ := (*prevSelected).(*ResultItem)
				prevRes.Unselect()
				result.Select()
			})
		})
		result.BindMouseClick(func() {
			_, _ = glib.IdleAdd(func() {
				l.navigation.Enter()
			})
		})
		results = append(results, &result)
	}

	l.clearResults()

	// Set Margins
	if len(results) == 0 {
		l.resultsBox.SetMarginTop(0)
		l.resultsBox.SetMarginBottom(0)

	} else {
		l.resultsBox.SetMarginTop(3)

		// Select first one automatically
		results[0].Select()
	}

	// Show New Results
	for _, result := range results {
		l.resultsBox.Add(result.frame)
	}

	// Update Launcher
	resultItems := make([]*common.Result, 0)
	for _, r := range results {
		res := common.Result(r)
		resultItems = append(resultItems, &res)
	}
	l.navigation.SetItems(resultItems)
}

func (l *Launcher) clearResults() {
	// Clear navigation
	l.navigation.SetItems(make([]*common.Result, 0))

	// Get Children
	previousResults := l.resultsBox.GetChildren()

	// Remove Each From The Results
	previousResults.Foreach(func(prev interface{}) {
		item, ok := prev.(gtk.IWidget)
		if ok {
			l.resultsBox.Remove(item)
		}
	})
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
	l.clearResults()
}
