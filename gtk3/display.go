package gtk3

import (
	"github.com/gotk3/gotk3/gtk"
)

func centerAtTopOfScreen(window *gtk.Window) {
	oldX, oldY := window.GetPosition()

	screenHeight := oldY *2
	_, winHeight := window.GetSize()

	// By default, the launcher always centers in the monitor the
	// mouse is resting at. Therefore, we can just move it along the y axis
	// to get it to be centered at the top half of the monitor.
	newX := oldX
	newY := screenHeight/5 + winHeight/4 - 6

	window.Move(newX, newY)
}
