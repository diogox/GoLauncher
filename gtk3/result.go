package gtk3

import (
	"github.com/diogox/GoLauncher/gtk3/glade"
	"github.com/gotk3/gotk3/gtk"
)

const GladeResultFile = "/home/diogox/go/src/github.com/diogox/GoLauncher/assets/result_item.glade"

const ResultFrameID = "item-frame"
const IconID = "item-icon"
const NameLabelID = "item-name"
const DescriptionLabelID = "item-descr"
const ShortcutLabelID = "item-shortcut"

func NewResultItem(cssProvider *gtk.CssProvider, title string, description string) ResultItem {
	bldr, err := glade.BuildFromFile(GladeResultFile)
	if err != nil {
		panic(err)
	}

	resultEventBox, err := glade.GetEventBox(bldr, ResultFrameID)
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

	shotcutLabel, err := glade.GetLabel(bldr, ShortcutLabelID)
	if err != nil {
		panic(err)
	}

	iconImg, err := glade.GetImage(bldr, IconID)
	if err != nil {
		panic(err)
	}

	nameLabel.SetText(title)
	descrLabel.SetText(description)

	// Set Styles
	setStyleClass(cssProvider, &nameLabel.Widget, "item-name")
	setStyleClass(cssProvider, &descrLabel.Widget, "item-text")
	setStyleClass(cssProvider, &shotcutLabel.Widget, "item-shortcut")
	setStyleClass(cssProvider, &iconImg.Widget, "item-icon")

	return ResultItem {
		frame: resultEventBox,
		icon: iconImg,
		label: nameLabel,
		description: descrLabel,
		shortcut: shotcutLabel,
	}
}

type ResultItem struct {
	frame *gtk.EventBox
	icon *gtk.Image
	label *gtk.Label
	description *gtk.Label
	shortcut *gtk.Label
}
