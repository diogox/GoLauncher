package gtk3

import (
	kb "github.com/Isolus/go-keybinder"
	"github.com/diogox/GoLauncher/api"
	"github.com/diogox/GoLauncher/api/actions"
	"github.com/diogox/GoLauncher/gtk3/glade"
	"github.com/diogox/GoLauncher/navigation"
	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"strconv"
	"unsafe"
)

const GladeFile = "/home/diogox/go/src/github.com/diogox/GoLauncher/gtk3/assets/launcher.glade"
const CssFile = "/home/diogox/go/src/github.com/diogox/GoLauncher/gtk3/assets/theme.css"

const WindowID = "window"
const BodyID = "body"

//const InputBoxID = "input-box"
const InputID = "input"
const PrefsBtnID = "prefs_btn"
const ResultsBoxScrollableID = "result_box_scrollable"
const ResultsBoxID = "result_box"

func NewLauncher(preferences *api.Preferences) *Launcher {

	// Initiate gtk (Must be here, so that it occurs before the hotkey binding)
	gtk.Init(nil)

	// Build from glade file
	bldr, err := glade.BuildFromFile(GladeFile)
	if err != nil {
		panic(err)
	}

	// Get CSS provider
	_, err = setCssProvider(CssFile)
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
		panic("Failed to get Entry: " + err.Error())
	}

	// Get preferences button
	prefsBtn, err := glade.GetButton(bldr, PrefsBtnID)
	if err != nil {
		panic("Failed to get Button: " + err.Error())
	}

	resultsScrollableBox, err := glade.GetScrolledWindow(bldr, ResultsBoxScrollableID)
	if err != nil {
		panic("Failed to get ScrolledWindow: " + err.Error())
	}

	// Create Results container
	resultsBox, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	st, _ := resultsBox.GetStyleContext()
	st.AddClass("result-box")

	// Add it to the scrolled window
	resultsScrollableBox.Add(resultsBox)

	scrollController := navigation.NewScrollController(resultsScrollableBox, 4) // TODO: Change '4' to get value from preferences

	nav := navigation.NewNavigation(scrollController)

	return &Launcher{
		hotkey:               nil,
		preferences:          preferences,
		window:               win,
		body:                 body,
		input:                input,
		prefsBtn:             prefsBtn,
		resultsBox:           resultsBox,
		resultsScrollableBox: resultsScrollableBox,
		navigation:           &nav,
		isVisible:            true,
	}
}

// TODO: Probably better to make a LauncherOptions object
type Launcher struct {
	hotkey               *string
	preferences          *api.Preferences
	window               *gtk.Window
	body                 *gtk.Box
	input                *gtk.Entry
	prefsBtn             *gtk.Button
	resultsBox           *gtk.Box
	resultsScrollableBox *gtk.ScrolledWindow
	navigation           *navigation.Navigation
	isVisible            bool
}

func (l *Launcher) HandleInput(callback func(string)) {
	_, _ = l.input.Connect("changed", func(entry *gtk.Entry) {
		input, err := entry.GetText()
		if err != nil {
			panic(err)
		}

		if input == "" {
			l.clearResults()
			return
		}

		callback(input)
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

	/* TODO: Uncomment this!
	// When the window loses focus, hide it
	_, _ = l.window.Connect("focus-out-event", func(widget *gtk.Window, event *gdk.Event) {
		_, _ = glib.IdleAdd(l.hide)
	})
	*/

	// Detect navigation ('Up', 'Down', 'Enter')
	_, _ = l.window.Connect("key-press-event", func(widget *gtk.Window, event *gdk.Event) {
		keyEvent := &gdk.EventKey{
			Event: event,
		}

		var result, prevResult *api.Result

		// Resolve action
		const KEY_Enter = 65293
		key := keyEvent.KeyVal()
		switch key {
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
			if keyEvent.State() == gdk.GDK_MOD1_MASK {
				if (key >= 48) && (key <= 57) { // info: (48 == '0')(57 == '9')
					index, _ := strconv.Atoi(string(key))
					r, err := l.navigation.At(index - 1)
					if err == nil {
						// Select new item
						currentItem := (*r).(*ResultItem)
						currentItem.Select()

						// Unselect previous item
						prev := l.navigation.SetSelected(r)
						prevItem, _ := (*prev).(*ResultItem)
						prevItem.Unselect()

						// Run Action
						currentItem.OnEnterAction().Run()
					}
				}
			}
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

	_, _ = l.prefsBtn.Connect("clicked", func(btn *gtk.Button) {
		_, _ = glib.IdleAdd(ShowSettingsWindow, l.preferences)
	})

	l.navigation.SetOnItemEnter(func(action api.Action) {
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
		_, _ = glib.IdleAdd(l.input.SetText, query)
		_, _ = glib.IdleAdd(l.input.SetPosition, len(query))
	})
	actions.SetupRenderResultList(func(results []api.Result) {
		_, _ = glib.IdleAdd(l.ShowResults, results)
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

	// Unbind previous hotkey, if exists
	if l.hotkey != nil {
		kb.Unbind(*l.hotkey, toggle)
	}

	// Bind hotkey
	kb.Init()
	kb.Bind(hotkey, toggle, nil)
	l.hotkey = &hotkey
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

func (l *Launcher) ShowResults(searchResults []api.Result) {

	results := make([]*ResultItem, 0)

	// Convert results
	for i, r := range searchResults {
		var result ResultItem
		if i <= 9 {
			result = NewResultItem(r.Title(), r.Description(), r.IconPath(), i+1, r.OnEnterAction(), r.OnAltEnterAction())
		} else {
			result = NewResultItem(r.Title(), r.Description(), r.IconPath(), -1, r.OnEnterAction(), r.OnAltEnterAction())
		}
		result.BindMouseHover(func() {
			_, _ = glib.IdleAdd(func() {
				res := api.Result(&result)
				prevSelected := l.navigation.SetSelected(&res)
				prevRes, _ := (*prevSelected).(*ResultItem)
				prevRes.Unselect()
				result.Select()
			})
		})
		result.BindMouseClick(func() {
			_, _ = glib.IdleAdd(l.navigation.Enter)
		})
		results = append(results, &result)
	}

	l.clearResults()

	// Set Margins
	if len(results) != 0 {
		l.resultsScrollableBox.SetMarginTop(3)
		l.resultsScrollableBox.SetMarginBottom(10)

		// Select first one automatically
		results[0].Select()
	}

	// Show New Results
	for _, result := range results {
		l.resultsBox.Add(result.frame)
	}

	// Update Launcher
	resultItems := make([]*api.Result, 0)
	for _, r := range results {
		res := api.Result(r)
		resultItems = append(resultItems, &res)
	}
	l.navigation.SetItems(resultItems)

	// Set ScrolledWindow height
	resultItemHeight := float64(0)
	if len(results) != 0 {
		_, height := results[0].frame.GetPreferredHeight()
		resultItemHeight = float64(height)
	}

	newScrolledHeight := resultItemHeight * float64(len(results))
	if len(results) > 4 { // TODO: Get '4' from the preferences. This is the maximum number of results to show
		newScrolledHeight = float64(resultItemHeight * 4)
	}
	l.navigation.ScrollController.SetHeight(int(newScrolledHeight))

	// Set 'scroll interval', so that when the user scrolls, it will skip, exactly, one item's height

	newAdjustment, err := gtk.AdjustmentNew(-1, float64(0), newScrolledHeight, resultItemHeight, resultItemHeight, newScrolledHeight)
	if err != nil {
		panic(err)
	}
	l.navigation.ScrollController.SetAdjustment(newAdjustment)
	l.navigation.ScrollController.SetNewResults(len(results))
}

func (l *Launcher) clearResults() {
	// Clear navigation
	l.navigation.SetItems(make([]*api.Result, 0))

	// Get Children
	previousResults := l.resultsBox.GetChildren()

	// Remove Each From The Results
	previousResults.Foreach(func(prev interface{}) {
		item, ok := prev.(gtk.IWidget)
		if ok {
			l.resultsBox.Remove(item)
		}
	})

	// Set ScrolledWindow height
	l.navigation.ScrollController.SetHeight(0)

	// Remove margins
	l.resultsScrollableBox.SetMarginTop(0)
	l.resultsScrollableBox.SetMarginBottom(0)
}

func (l *Launcher) show() {

	// Keep input or not
	keepInput, err := (*l.preferences).GetPreference(api.PreferenceKeepInputOnHide)
	if err != nil {
		panic(err)
	}

	if keepInput == api.PreferenceFALSE {
		l.ClearInput()
	}

	// Position (after clearing results - otherwise it won't center properly)
	centerAtTopOfScreen(l.window)

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
