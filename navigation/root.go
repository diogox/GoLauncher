package navigation

import (
	"errors"
	"github.com/diogox/GoLauncher/api"
)

func NewNavigation(scrollController *ScrollController) Navigation {
	return Navigation{
		onItemEnter:  func(api.Action) {},
		currentIndex: -1,
		items:        make([]*api.Result, 0),
		ScrollController: scrollController,
	}
}

type Navigation struct {
	onItemEnter      func(api.Action)
	currentIndex     int
	items            []*api.Result
	ScrollController *ScrollController
}

func (n *Navigation) SetOnItemEnter(onItemEnter func(api.Action)) {
	n.onItemEnter = onItemEnter
}

func (n *Navigation) SetItems(items []*api.Result) {
	n.items = items

	if len(items) == 0 {
		n.currentIndex = -1
	} else {
		n.currentIndex = 0
	}
}

func (n *Navigation) Up() (*api.Result, *api.Result) {
	// No items to navigate through
	if len(n.items) == 0 {
		return nil, nil
	}

	// Update ScrolledWindow
	n.ScrollController.SignalMoveUp()

	// There's only a previous item if a current index exists
	var prevItem *api.Result
	if n.currentIndex != -1 {
		prevItem = n.items[n.currentIndex]
	}

	// Get next index
	index := n.currentIndex - 1
	if index >= 0 {

		// Set current index
		n.currentIndex = index
		return n.items[index], prevItem
	}

	// Overflowed, skipping to the last index (end of the list)
	lastIndex := len(n.items) - 1
	n.currentIndex = lastIndex
	return n.items[lastIndex], prevItem
}

func (n *Navigation) Down() (*api.Result, *api.Result) {
	// No items to navigate through
	if len(n.items) == 0 {
		return nil, nil
	}

	// Update ScrolledWindow
	n.ScrollController.SignalMoveDown()

	// There's only a previous item if a current index exists
	var prevItem *api.Result
	if n.currentIndex != -1 {
		prevItem = n.items[n.currentIndex]
	}

	// Get next index
	index := n.currentIndex + 1
	if index < len(n.items) {

		// Set current index
		n.currentIndex = index
		return n.items[index], prevItem
	}

	// Overflowed, skipping to the first index (beginning of the list)
	firstIndex := 0
	n.currentIndex = firstIndex
	return n.items[firstIndex], prevItem
}

func (n *Navigation) Enter() {
	if n.currentIndex != -1 {
		item := n.items[n.currentIndex]
		n.onItemEnter((*item).OnEnterAction())
	}
}

func (n *Navigation) AltEnter() {
	if n.currentIndex != -1 {
		item := n.items[n.currentIndex]
		n.onItemEnter((*item).OnAltEnterAction())
	}
}

func (n *Navigation) SetSelected(item *api.Result) *api.Result {
	prevSelected := n.items[n.currentIndex]

	for i, it := range n.items {
		if *it == *item {
			// Save current item's index
			n.currentIndex = i
			break
		}
	}

	// Update scroller
	n.ScrollController.SetSelectedIndex(n.currentIndex)

	return prevSelected
}

func (n *Navigation) At(index int) (*api.Result, error) {
	if index >= 0 && index < len(n.items) {
		return n.items[index], nil
	}

	return nil, errors.New("index not in range")
}
