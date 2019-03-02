package gtk3

import (
	"fmt"
	"github.com/diogox/GoLauncher/api"
	"github.com/diogox/GoLauncher/gtk3/glade"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

const GladeResultFile = "/home/diogox/go/src/github.com/diogox/GoLauncher/gtk3/assets/result_item.glade"

const ResultFrameID = "item-frame"
const ResultBoxID = "item-box"
const IconID = "item-icon"
const NameLabelID = "item-name"
const DescriptionLabelID = "item-descr"
const ShortcutLabelID = "item-shortcut"

func NewResultItem(title string, description string, iconName string, position int, onEnterAction api.Action, onAltEnterAction api.Action) ResultItem {
	bldr, err := glade.BuildFromFile(GladeResultFile)
	if err != nil {
		panic(err)
	}

	resultEventFrame, err := glade.GetEventBox(bldr, ResultFrameID)
	if err != nil {
		panic(err)
	}

	resultEventBox, err := glade.GetEventBox(bldr, ResultBoxID)
	if err != nil {
		panic(err)
	}

	nameLabel, err := glade.GetLabel(bldr, NameLabelID)
	if err != nil {
		panic(err)
	}

	descrLabel, err := glade.GetLabel(bldr, DescriptionLabelID)
	if err != nil {
		panic(err)
	}

	shortcutLabel, err := glade.GetLabel(bldr, ShortcutLabelID)
	if err != nil {
		panic(err)
	}

	iconImg, err := glade.GetImage(bldr, IconID)
	if err != nil {
		panic(err)
	}

	nameLabel.SetText(title)
	descrLabel.SetText(description)

	shortcut := fmt.Sprintf("Alt+%d", position)
	shortcutLabel.SetText(shortcut)

	iconImg.SetFromIconName(iconName, gtk.ICON_SIZE_DND)
	iconImg.SetPixelSize(40)

	resultItem := ResultItem{
		onEnterAction: onEnterAction,
		onAltEnterAction: onAltEnterAction,

		frame:       resultEventFrame,
		box:         resultEventBox,
		icon:        iconImg,
		label:       nameLabel,
		description: descrLabel,
		shortcut:    shortcutLabel,
	}

	return resultItem
}

type ResultItem struct {
	onEnterAction    api.Action
	onAltEnterAction api.Action

	frame       *gtk.EventBox
	box         *gtk.EventBox
	icon        *gtk.Image
	label       *gtk.Label
	description *gtk.Label
	shortcut    *gtk.Label
}

func (r *ResultItem) Title() string {
	title, _ := r.label.GetText()
	return title
}

func (r *ResultItem) Description() string {
	description, _ := r.description.GetText()
	return description
}

func (r *ResultItem) IconPath() string {
	iconName, _ := r.icon.GetIconName()
	return iconName
}

func (r *ResultItem) OnEnterAction() api.Action {
	return r.onEnterAction
}

func (r *ResultItem) OnAltEnterAction() api.Action {
	fmt.Println("With Alt modifier!")
	return r.onAltEnterAction
}

func (r *ResultItem) Select() {
	setStyleClass(&r.box.Widget, "selected")
}

func (r *ResultItem) Unselect() {
	removeStyleClass(&r.box.Widget, "selected")
}

func (r *ResultItem) BindMouseHover(callback func()) {
	_, _ = r.frame.Connect("enter-notify-event", func(eventBox *gtk.EventBox, event *gdk.Event) {
		callback()
	})
}

func (r *ResultItem) BindMouseClick(callback func()) {
	_, _ = r.frame.Connect("button_press_event", func(eventBox *gtk.EventBox, event *gdk.Event) {
		callback()
	})
}
