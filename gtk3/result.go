package gtk3

import (
	"github.com/diogox/GoLauncher/gtk3/glade"
	"github.com/gotk3/gotk3/gtk"
)

const GladeResultFile = "/home/diogox/go/src/github.com/diogox/GoLauncher/assets/result_item.glade"

const ResultFrameName = "item-frame"
const NameLabel = "item-name"
const DescriptionLabel = "item-descr"

func NewResultItem(title string, description string) ResultItem {
	bldr, err := glade.BuildFromFile(GladeResultFile)
	if err != nil {
		panic(err)
	}

	resultEventBox, err := glade.GetEventBox(bldr, ResultFrameName)
	if err != nil {
		panic(err)
	}

	nameLabel, err := glade.GetLabel(bldr, NameLabel)
	if err != nil {
		panic(err)
	}


	descrLabel, err := glade.GetLabel(bldr, DescriptionLabel)
	if err != nil {
		panic(err)
	}

	nameLabel.SetText(title)
	descrLabel.SetText(description)

	return ResultItem {
		frame: resultEventBox,
		label: nameLabel,
		description: descrLabel,
	}
}

type ResultItem struct {
	frame *gtk.EventBox
	label *gtk.Label
	description *gtk.Label
}
