package screen

import (
	"errors"
	"github.com/gotk3/gotk3/gdk"
)

type position struct {
	X int
	Y int
}

var screenPositionsCache = make(map[int]position, 0)

func cachePositionForScreen(screen *gdk.Screen, xPos int, yPos int) {
	screenPositionsCache[screen.GetScreenNumber()] = position{
		X: xPos,
		Y: yPos,
	}
}

func getCachedPositionForScreen(screen *gdk.Screen) (int, int, error) {
	position, ok := screenPositionsCache[screen.GetScreenNumber()]
	if !ok {
		return -1, -1, errors.New("no position in cache for this screen")
	}

	return position.X, position.Y, nil
}