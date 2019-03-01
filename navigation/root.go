package navigation

import (
	"github.com/diogox/GoLauncher/common"
)

func NewNavigation(items []*common.Result) Navigation {
	return Navigation{
		onItemEnter: func(common.Action) {},
		currentIndex: -1,
		items:        items,
	}
}

type Navigation struct {
	onItemEnter func(common.Action)
	currentIndex int
	items        []*common.Result
}

func (n *Navigation) SetOnItemEnter(onItemEnter func(common.Action)) {
	n.onItemEnter = onItemEnter
}

func (n *Navigation) SetItems(items []*common.Result) {
	n.items = items

	if len(items) == 0 {
		n.currentIndex = -1
	} else {
		n.currentIndex = 0
	}
}

func (n *Navigation) Up() (*common.Result, *common.Result) {
	if len(n.items) == 0 {
		return nil, nil
	}

	var prevItem *common.Result
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

func (n *Navigation) Down() (*common.Result, *common.Result) {
	if len(n.items) == 0 {
		return nil, nil
	}

	var prevItem *common.Result
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

func (n *Navigation) SetSelected(item *common.Result) *common.Result {
	prevSelected := n.items[n.currentIndex]

	for i, it := range n.items {
		if *it == *item {
			n.currentIndex = i
			break
		}
	}

	return prevSelected
}