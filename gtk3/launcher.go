package gtk3

import (
	"fmt"
	kb "github.com/Isolus/go-keybinder"
	"github.com/diogox/GoLauncher/api"
	"github.com/diogox/GoLauncher/api/actions"
	"github.com/diogox/GoLauncher/gtk3/glade"
	"github.com/diogox/GoLauncher/gtk3/utils"
	"github.com/diogox/GoLauncher/navigation"
	"github.com/diogox/GoLauncher/pkg/screen"
	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"strconv"
	"strings"
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
	_, err = utils.SetCssProvider(CssFile)
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

	// Get number of results to show
	nOfResultsToShowStr, _ := (*preferences).GetPreference(api.PreferenceNResultsToShow)
	nOfResultsToShow, _ := strconv.Atoi(nOfResultsToShowStr)

	scrollController := navigation.NewScrollController(resultsScrollableBox, nOfResultsToShow)

	// Update ScrollController every time a change is made to the preference 'PreferenceNResultsToShow'
	(*preferences).BindCallback(api.PreferenceNResultsToShow, func(arg interface{}) {
		argString, _ := arg.(string)
		nOfResultsToShow, _ := strconv.Atoi(argString)
		scrollController.SetNOfItemsToShow(nOfResultsToShow)
	})

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
	isVisible            bool
	preferences          *api.Preferences
	window               *gtk.Window
	body                 *gtk.Box
	input                *gtk.Entry
	prefsBtn             *gtk.Button
	resultsBox           *gtk.Box
	resultsScrollableBox *gtk.ScrolledWindow
	navigation           *navigation.Navigation
}

func (l *Launcher) HandleInput(callback func(string), onEmptyCallback func()) {
	_, _ = l.input.Connect("changed", func(entry *gtk.Entry) {

		input := getTrimmedInput(entry)

		if input == "" {
			l.clearResults()
			_, _ = glib.IdleAdd(onEmptyCallback)
			return
		}

		_, _ = glib.IdleAdd(callback, input)
	})
}

func (l *Launcher) Start() error {

	// Keep the launcher above everything
	l.window.SetKeepAbove(true)

	// Set the groundwork for transparency
	utils.ScreenChanged(l.window)

	// When the monitor/screen changes, check for transparency support
	_, err := l.window.Connect("screen-changed", func(window *gtk.Window, oldScreen *gdk.Screen, userData ...interface{}) {
		utils.ScreenChanged(window)
	})
	if err != nil {
		return err
	}

	// Set transparency on draw
	_, err = l.window.Connect("draw", func(window *gtk.Window, context *cairo.Context) {
		utils.SetTransparent(window, context)
	})
	if err != nil {
		return err
	}

	// When the window loses focus, hide it
	_, err = l.window.Connect("focus-out-event", func(widget *gtk.Window, event *gdk.Event) {
		_, _ = glib.IdleAdd(l.hide)
	})
	if err != nil {
		return err
	}

	// Detect navigation ('Up', 'Down', 'Enter')
	_, err = l.window.Connect("key-press-event", func(widget *gtk.Window, event *gdk.Event) {
		keyEvent := &gdk.EventKey{
			Event: event,
		}

		var item, prevItem *navigation.NavigationItem

		// Resolve action
		const KEY_Enter = 65293
		key := keyEvent.KeyVal()
		switch key {
		case gdk.KEY_Up:
			item, prevItem = l.navigation.Up()
			if item == nil {
				return
			}
		case gdk.KEY_Down:
			item, prevItem = l.navigation.Down()
			if item == nil {
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

				// Get result index
				index := rune(key - 97 + 10) // Magic ascii transformation
				if (key >= 48) && (key <= 57) { // info: (48 == '0')(57 == '9')
					indexInt, _ := strconv.Atoi(string(key))
					index = rune(indexInt)
				}

				// Get result at that index
				current, err := l.navigation.At(int(index) - 1)
				if err == nil {

					currentSearchItem := current.SearchResult
					currentResultItem := current.ResulItem

					// Select new item
					currentResultItem.Select()

					// Unselect previous item
					prev := l.navigation.SetSelected(&currentSearchItem)
					prevItem := prev.ResulItem
					prevItem.Unselect()

					// Run Action
					err = currentSearchItem.OnEnterAction().Run()
					if err != nil {
						panic(err)
					}
				}
			}
			return
		}

		prevItem.ResulItem.Unselect()
		item.ResulItem.Select()
	})
	if err != nil {
		return err
	}

	_, err = l.prefsBtn.Connect("clicked", func(btn *gtk.Button) {
		_, _ = glib.IdleAdd(ShowSettingsWindow, l.preferences)
	})
	if err != nil {
		return err
	}

	l.navigation.SetOnItemEnter(func(action api.Action) {
		err := action.Run()
		if err != nil {
			panic(err)
		}

		if !action.KeepAppOpen() {
			l.hide()
		}
	})

	// Setup actions
	actions.SetupCopyToClipboard(func(text string) error {
		var err error

		_, _ = glib.IdleAdd(func(text string) {
			clipboard, e := gtk.ClipboardGet(gdk.SELECTION_CLIPBOARD)
			if e != nil {
				// TODO: Log error
				err = e
				return
			}

			clipboard.SetText(text)
			clipboard.Store()
		}, text)

		return err
	})
	actions.SetupSetUserQuery(func(query string) error {

		// Set the user query
		_, err := glib.IdleAdd(l.input.SetText, query)
		if err != nil {
			return err
		}

		// Set the cursor's position to the end of the query
		_, err = glib.IdleAdd(l.input.SetPosition, len(query))
		if err != nil {
			return err
		}

		return nil
	})
	actions.SetupRenderResultList(func(results []api.SearchResult) error {
		_, err := glib.IdleAdd(l.ShowResults, results)
		if err != nil {
			return err
		}

		return nil
	})

	// Show everything in the app
	l.show()

	// Start loop
	go gtk.Main()

	return nil
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
	l.input.DeleteText(0, -1)
}

func (l *Launcher) ShowResults(searchResults []api.SearchResult) {

	l.clearResults()

	// Update Navigation
	l.navigation.SetItems(searchResults)

	items := l.navigation.GetItems()
	results := make([]api.ResultItem, 0)

	// Convert results
	for i, item := range items {
		// Get results
		searchResult := item.SearchResult
		resultItem := item.ResulItem

		position := fmt.Sprintf("%d", i+1)
		if i > 9 {
			position = fmt.Sprintf("%s", string(rune(97+i-9)))
		}

		resultItem.SetPosition(position)

		resultItem.BindMouseHover(func() {
			_, _ = glib.IdleAdd(func() {
				prevSelected := l.navigation.SetSelected(&searchResult)
				prevSelected.ResulItem.Unselect()
				resultItem.Select()
			})
		})
		resultItem.BindMouseClick(func() {
			_, _ = glib.IdleAdd(l.navigation.Enter)
		})
		results = append(results, resultItem)
	}

	// Set Margins
	if len(results) != 0 {
		l.resultsScrollableBox.SetMarginTop(3)
		l.resultsScrollableBox.SetMarginBottom(10)

		// Select first one automatically
		results[0].Select()

		// Check if any of the results should be the automatic default
		for i, searchResult := range searchResults {
			if searchResult.IsDefaultSelect() {

				// Unselect the first item that was automatically selected
				results[0].Unselect()

				// Select default item
				results[i].Select()
				break
			}
		}
	}

	// Show New Results
	for _, r := range results {
		r.AccessInternals(func(args... interface{}) {
			// get frame
			frame, _ := args[0].(*gtk.EventBox)

			// Add it to the results box
			l.resultsBox.Add(frame)
		})
	}

	// Show ScrolledWindow here (Had to hide it, initially, to keep from showing an awkward whitespace. Couldn't find another way to fix that...)
	l.resultsScrollableBox.Show()

	// Set ScrolledWindow height
	resultItemHeight := float64(0)
	if len(results) != 0 {
		results[0].AccessInternals(func(args... interface{}) {
			// Get frame
			frame, _ := args[0].(*gtk.EventBox)

			// Get height of the item
			_, height := frame.GetPreferredHeight()

			// Return it
			resultItemHeight = float64(height)
		})
	}

	newScrolledHeight := resultItemHeight * float64(len(results))

	// Get number of results to display
	nOfResultsToShowStr, err := (*l.preferences).GetPreference(api.PreferenceNResultsToShow)
	nOfResultsToShow, _ := strconv.Atoi(nOfResultsToShowStr)

	if len(results) > nOfResultsToShow {
		newScrolledHeight = resultItemHeight * float64(nOfResultsToShow)
	}

	// Set adjustment
	newAdjustment, err := gtk.AdjustmentNew(-1, float64(0), -1, resultItemHeight, resultItemHeight, newScrolledHeight)
	if err != nil {
		panic(err)
	}

	// Update ScrollController
	l.navigation.ScrollController.SetHeight(int(newScrolledHeight))
	l.navigation.ScrollController.SetAdjustment(newAdjustment)
	l.navigation.ScrollController.SetNewResults(len(results))
}

func (l *Launcher) clearResults() {
	// Need to hide ScrolledWindow, otherwise it shows an awkward whitespace...
	l.resultsScrollableBox.Hide()

	// Clear navigation
	l.navigation.SetItems(make([]api.SearchResult, 0))

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
		// TODO: This is a hack, since `l.input.Emit("changed", l.input)` is not working as it should.
		// It is here so that the input realizes it's been changed and procs the 'onChange' closure that brings up the most frequent apps.
		// If this wasn't here, every time the launcher is hidden with "" as the query, it doesn't recognize the change in input when it's brought back up.
		_, _ = glib.IdleAdd(l.input.SetText, " ")
		l.ClearInput()
	}

	// Position (after clearing results - otherwise it won't center properly)
	err = screen.CenterAtTopOfScreen(l.window)
	if err != nil {
		panic(err)
	}

	// Show
	l.window.ShowAll()
	l.isVisible = true

	// Need to hide, otherwise it shows whitespace (Couldn't figure out why...)
	// TODO: Fix this mess by removing the need to hide the ScrolledWindow!!
	if isInputEmpty := len(getTrimmedInput(l.input)) == 0; isInputEmpty {
		l.resultsScrollableBox.Hide()
		_, _ = glib.IdleAdd(l.input.SetText, " ")
		l.ClearInput()
	}

	// Focus
	l.window.PresentWithTime(kb.GetCurrentEventTime())
}

func (l *Launcher) hide() {

	// Hide
	l.window.Hide()
	l.isVisible = false
}

func getTrimmedInput(entry *gtk.Entry) string {
	// Get input
	input, err := entry.GetText()
	if err != nil {
		panic(err)
	}

	// Trim whitespace
	input = strings.TrimSpace(input)

	return input
}
