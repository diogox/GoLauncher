package events

import (
	"github.com/diogox/GoLauncher/api"
)

func ItemEnterNew(data interface{}) ItemEnter {
	return ItemEnter{
		data: data,
	}
}

type ItemEnter struct {
	api.BaseEvent
	data interface{}
}

func (kqe ItemEnter) Data() interface{} {
	return kqe.data
}