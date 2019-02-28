package actions

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

func NewCopyToClipboardAction(text string) CopyToClipboard {
	return CopyToClipboard{
		text: text,
	}
}

type CopyToClipboard struct {
	text string
}

func (CopyToClipboard) keepAppOpen() bool {
	return false
}

func (c *CopyToClipboard) run() {
	clipboard, err := gtk.ClipboardGet(gdk.SELECTION_CLIPBOARD)
	if err != nil {
		panic("Failed to get clipboard!")
	}

	clipboard.SetText(c.text)
	clipboard.Store()
}

