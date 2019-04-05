package navigation

import (
	"errors"
	"github.com/diogox/GoLauncher/api"
)

func NewNavigation(scrollController *ScrollController) Navigation {
	return Navigation{
		onItemEnter:  func(api.Action) {},
		currentIndex: -1,
		items:        make([]*NavigationItem, 0),
		ScrollController: scrollController,
	}
}

type Navigation struct {
	onItemEnter      func(api.Action)
	currentIndex     int
	items            []*NavigationItem
	ScrollController *ScrollController
}

func (n *Navigation) SetOnItemEnter(onItemEnter func(api.Action)) {
	n.onItemEnter = onItemEnter
}

func (n *Navigation) SetItems(searchResults []api.SearchResult) {
	items := make([]*NavigationItem, 0)

	// Convert to entries
	for _, searchResult := range searchResults {
		items = append(items, newNavigationItem(searchResult))
	}

	// Set current items
	n.items = items

	// Set current index
	if len(items) == 0 {
		n.currentIndex = -1
	} else {
		n.currentIndex = 0
	}
}

func (n *Navigation) Up() (*NavigationItem, *NavigationItem) {
	// No items to navigate through
	if len(n.items) == 0 {
		return nil, nil
	}

	// Update ScrolledWindow
	n.ScrollController.SignalMoveUp()

	// There's only a previous item if a current index exists
	var prevItem *NavigationItem
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

func (n *Navigation) Down() (*NavigationItem, *NavigationItem) {
	// No items to navigate through
	if len(n.items) == 0 {
		return nil, nil
	}

	// Update ScrolledWindow
	n.ScrollController.SignalMoveDown()

	// There's only a previous item if a current index exists
	var prevItem *NavigationItem
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
		item := n.items[n.currentIndex].SearchResult
		n.onItemEnter(item.OnEnterAction())
	}
}

func (n *Navigation) AltEnter() {
	if n.currentIndex != -1 {
		item := n.items[n.currentIndex].SearchResult
		n.onItemEnter(item.OnAltEnterAction())
	}
}

func (n *Navigation) SetSelected(item *api.SearchResult) *NavigationItem {
	prevSelected := n.items[n.currentIndex]

	for i, it := range n.items {
		if it.SearchResult == *item {
			// Save current item's index
			n.currentIndex = i
			break
		}
	}

	// Update scroller
	n.ScrollController.SetSelectedIndex(n.currentIndex)

	return prevSelected
}

func (n *Navigation) At(index int) (*NavigationItem, error) {
	if index >= 0 && index < len(n.items) {
		return n.items[index], nil
	}

	return nil, errors.New("index not in range")
}

func (n *Navigation) GetItems() []NavigationItem {
	items := make([]NavigationItem, 0)

	for _, item := range n.items {
		items = append(items, *item)
	}

	return items
}