package navigation

import (
	"github.com/diogox/GoLauncher/api"
)

func NewNavigation(items []*api.Result) Navigation {
	return Navigation{
		onItemEnter: func(api.Action) {},
		currentIndex: -1,
		items:        items,
	}
}

type Navigation struct {
	onItemEnter func(api.Action)
	currentIndex int
	items        []*api.Result
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
	if len(n.items) == 0 {
		return nil, nil
	}

	var prevItem *api.Result
	if n.currentIndex != -1 {
		prevItem = n.items[n.currentIndex]
	}

	index := n.currentIndex - 1
	if index >= 0 {
		n.currentIndex = index
		return n.items[index], prevItem
	}

	lastIndex := len(n.items) - 1
	n.currentIndex = lastIndex
	return n.items[lastIndex], prevItem
}

func (n *Navigation) Down() (*api.Result, *api.Result) {
	if len(n.items) == 0 {
		return nil, nil
	}

	var prevItem *api.Result
	if n.currentIndex != -1 {
		prevItem = n.items[n.currentIndex]
	}

	index := n.currentIndex + 1
	if index < len(n.items) {
		n.currentIndex = index
		return n.items[index], prevItem
	}

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
			n.currentIndex = i
			break
		}
	}

	return prevSelected
}