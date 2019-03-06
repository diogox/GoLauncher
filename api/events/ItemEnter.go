package events

import "github.com/diogox/GoLauncher/api"

func ItemEnterNew(data interface{}) ItemEnter {
	return ItemEnter{
		Type: api.ITEM_ENTER_EVENT,
		data: data,
	}
}

type ItemEnter struct {
	Type string
	data interface{}
}

func (ie ItemEnter) GetEventType() string {
	return ie.Type
}

func (ie ItemEnter) Data() interface{} {
	return ie.data
}
