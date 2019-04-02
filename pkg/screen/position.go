package screen

import "github.com/gotk3/gotk3/gtk"

// 	This whole cache thing is a workaround for the fact that, sometimes, when toggling the launcher,
//  it'll show up at different places due to the changes in its height. This way, it is ensured that
//  the launcher will surface always at the same point in a screen.
//  TODO: Problems with this may include:
// 		* It not actually getting the screen the user is using, through the window method `GetScreen`
// 		* Changing the preferences to show/hide the most frequent apps and the launcher not showing up at
// 		1/5 of the screens height because of the cached position. Maybe I should provide a way to reset
// 		the cache on preference changes?

func CenterAtTopOfScreen(window *gtk.Window) error {

	// Get screen the user is currently using
	screen, err := window.GetScreen()
	if err != nil {
		return err
	}

	// Check for a cached position for that screen
	xPos, yPos, err := getCachedPositionForScreen(screen)
	if err == nil {

		// Set window position
		window.Move(xPos, yPos)
		return nil
	}

	// Get new position
	newX, newY := getNewPosition(window)

	// Set window position
	window.Move(newX, newY)

	// Cache position on this screen
	cachePositionForScreen(screen, newX, newY)

	return nil
}

func getNewPosition(window *gtk.Window) (int, int) {

	// Get current window position
	oldX, oldY := window.GetPosition()

	// Get dimensions
	screenHeight := oldY * 2
	_, winHeight := window.GetSize()

	// By default, the launcher always centers in the monitor the
	// mouse is resting at. Therefore, we can just move it along the y axis
	// to get it to be centered at the top half of the monitor.
	newX := oldX
	newY := screenHeight/5 + winHeight/4 - 14

	return newX, newY
}