package actions

import "github.com/diogox/GoLauncher/api"

var copyToClipboardInstance *CopyToClipboard

// Copying text to the clipboard may rely on the GUI framework.
// To keep the action platform-agnostic, we need to set it up before using it.
func SetupCopyToClipboard(copyToClipboardCallback func(string)) {
	copyToClipboardInstance = &CopyToClipboard {
		Type: api.COPY_TO_CLIPBOARD_ACTION,
		copyToClipboardCallback: copyToClipboardCallback,
	}
}

func NewCopyToClipboard(text string) CopyToClipboard {
	if copyToClipboardInstance == nil {
		panic("You need to setup this action before you can use it!")
	}

	newInstance := *copyToClipboardInstance
	newInstance.Text = text
	return newInstance
}

type CopyToClipboard struct {
	Type string
	copyToClipboardCallback func(string)
	Text string
}

func (c CopyToClipboard) GetType() string {
	return c.Type
}

func (CopyToClipboard) KeepAppOpen() bool {
	return false
}

func (c CopyToClipboard) Run() {
	c.copyToClipboardCallback(c.Text)
}

