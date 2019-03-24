package navigation

import (
	"fmt"
	"github.com/gotk3/gotk3/gtk"
)

func NewScrollController(scrolledWindow *gtk.ScrolledWindow, itemsToShow int) *ScrollController {
	return &ScrollController{
		window:       scrolledWindow,
		nItemsToShow: itemsToShow,
	}
}

type ScrollController struct {
	window              *gtk.ScrolledWindow
	increment           float64
	currentIndex        int
	nItemsToShow        int
	nItemsAvailable     int
	itemVisibilityRange map[int]ViewingRange
}

func (sc *ScrollController) SetNewResults(nOfResults int) {
	sc.nItemsAvailable = nOfResults

	viewingArea := float64(sc.nItemsToShow) * sc.increment

	visibility := make(map[int]ViewingRange)
	for i := 0; i < nOfResults; i++ {
		min := float64(i) * sc.increment
		max := min + viewingArea

		visibility[i] = ViewingRange{
			Min: min,
			Max: max,
		}
	}

	sc.itemVisibilityRange = visibility
}
func (sc *ScrollController) SetAdjustment(adj *gtk.Adjustment) {
	sc.setIncrement(adj.GetStepIncrement())
	sc.window.SetVAdjustment(adj)
}

func (sc *ScrollController) setIncrement(inc float64) {
	sc.increment = inc
	sc.currentIndex = 0
}

func (sc *ScrollController) SetHeight(height int) {
	sc.window.SetSizeRequest(-1, height)
}

func (sc *ScrollController) SignalMoveUp() {

	// Assert new indexes
	previousIndex := sc.currentIndex
	sc.currentIndex = previousIndex - 1

	// Make sure we didn't try to go up from index 0
	if sc.currentIndex < 0 {

		// Assign index to the index of the last item
		sc.currentIndex = sc.nItemsAvailable - 1

		// Move scroll to first item
		lastMin := sc.itemVisibilityRange[sc.currentIndex].Min
		sc.window.GetVAdjustment().SetValue(lastMin)
		return
	}

	// Check if the scroll's lower level was not set to match the end of the previous item
	currentValue := sc.window.GetVAdjustment().GetValue()
	if currentValue < sc.itemVisibilityRange[sc.currentIndex].Min {
		return
	}

	// Set the scroll's upper level to match the beginning of the next item
	sc.window.GetVAdjustment().SetValue(sc.itemVisibilityRange[sc.currentIndex].Min)
}

func (sc *ScrollController) SignalMoveDown() {
	fmt.Println(sc.currentIndex)
	// Assert new indexes
	previousIndex := sc.currentIndex
	sc.currentIndex = previousIndex + 1

	// Make sure we didn't try to go down from the last index
	if sc.currentIndex >= sc.nItemsAvailable {

		// Assign index to the index of the first item
		sc.currentIndex = 0

		// Move scroll to first item
		sc.window.GetVAdjustment().SetValue(0)
		return
	}

	// Check if the scroll's lower level was not set to match the end of the previous item
	currentValue := sc.window.GetVAdjustment().GetValue()
	lowerBound := currentValue + (sc.increment * float64(sc.nItemsToShow))
	if lowerBound > sc.itemVisibilityRange[sc.currentIndex].Min {
		return
	}

	// Set the scroll's upper level to match the beginning of the next item
	itemInLowerBound := sc.itemVisibilityRange[sc.currentIndex].Min - (sc.increment * 3) // TODO: Get this value from `ViewingRange`
	sc.window.GetVAdjustment().SetValue(itemInLowerBound)
}

func (sc *ScrollController) SetSelectedIndex(index int) {
	sc.currentIndex = index
}

type ViewingRange struct {
	Min float64
	Max float64
}
