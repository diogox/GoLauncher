package gtk3

import (
	"fmt"
	"github.com/diogox/GoLauncher/gtk3/glade"
	"github.com/gotk3/gotk3/gtk"
)

const GladeResultFile = "/home/diogox/go/src/github.com/diogox/GoLauncher/gtk3/assets/result_item.glade"

const ResultFrameID = "item-frame"
const ResultBoxID = "item-box"
const IconID = "item-icon"
const NameLabelID = "item-name"
const DescriptionLabelID = "item-descr"
const ShortcutLabelID = "item-shortcut"

func NewResultItem(title string, description string, iconName string, position int) ResultItem {
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

	// TODO: Explore this option! (Scaling looks much better!)
	//p, _ := gdk.PixbufNewFromFileAtScale("/usr/share/icons/hicolor/48x48/apps/ulauncher.svg", 35, 35, true)
	//iconImg.SetFromPixbuf(p)

	return ResultItem{
		frame:       resultEventFrame,
		box:         resultEventBox,
		icon:        iconImg,
		label:       nameLabel,
		description: descrLabel,
		shortcut:    shortcutLabel,
	}
}

type ResultItem struct {
	frame       *gtk.EventBox
	box         *gtk.EventBox
	icon        *gtk.Image
	label       *gtk.Label
	description *gtk.Label
	shortcut    *gtk.Label
}

func (r *ResultItem) Select() {
	setStyleClass(&r.box.Widget, "selected")
}

func (r *ResultItem) Unselect() {
	removeStyleClass(&r.box.Widget, "selected")
}
