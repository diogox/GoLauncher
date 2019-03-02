package actions

var copyToClipboardInstance *CopyToClipboard

// Copying text to the clipboard may rely on the GUI framework.
// To keep the action platform-agnostic, we need to set it up before using it.
func SetupCopyToClipboard(copyToClipboardCallback func(string)) {
	copyToClipboardInstance = &CopyToClipboard {
		copyToClipboardCallback: copyToClipboardCallback,
	}
}

func NewCopyToClipboard(text string) CopyToClipboard {
	if copyToClipboardInstance == nil {
		panic("You need to setup this action before you can use it!")
	}

	newInstance := *copyToClipboardInstance
	newInstance.text = text
	return newInstance
}

type CopyToClipboard struct {
	copyToClipboardCallback func(string)
	text string
}

func (CopyToClipboard) KeepAppOpen() bool {
	return false
}

func (c CopyToClipboard) Run() {
	c.copyToClipboardCallback(c.text)
}

