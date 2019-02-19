package glade

import "github.com/gotk3/gotk3/gtk"

const WindowName = "window"
const BodyName = "body"
const InputBoxName = "input-box"
const InputName = "input"
const PrefsBtnName = "prefs_btn"
const ResultsBoxName = "result_box"
const ResultFrameName = "item-frame"

func BuildFromFile(fileName string) (*gtk.Builder, error) {
	bldr, err := gtk.BuilderNew()
	if err != nil {
		return nil, err
	}

	err = bldr.AddFromFile(fileName)
	if err != nil {
		return nil, err
	}

	return bldr, nil
}

func GetWindow(app *gtk.Builder, id string) (*gtk.Window, error) {
	obj, err := app.GetObject(id)
	if err != nil {
		return nil, err
	}

	win, ok := obj.(*gtk.Window)
	if !ok {
		return nil, err
	}

	return win, nil
}

func GetEntry(app *gtk.Builder, id string) (*gtk.Entry, error) {
	obj, err := app.GetObject(id)
	if err != nil {
		return nil, err
	}

	entry, ok := obj.(*gtk.Entry)
	if !ok {
		return nil, err
	}

	return entry, nil
}

func GetButton(app *gtk.Builder, id string) (*gtk.Button, error) {
	obj, err := app.GetObject(id)
	if err != nil {
		return nil, err
	}

	btn, ok := obj.(*gtk.Button)
	if !ok {
		return nil, err
	}

	return btn, nil
}

func GetBox(app *gtk.Builder, id string) (*gtk.Box, error) {
	obj, err := app.GetObject(id)
	if err != nil {
		return nil, err
	}

	box, ok := obj.(*gtk.Box)
	if !ok {
		return nil, err
	}

	return box, nil
}

func GetEventBox(app *gtk.Builder, id string) (*gtk.EventBox, error) {
	obj, err := app.GetObject(id)
	if err != nil {
		return nil, err
	}

	eventBox, ok := obj.(*gtk.EventBox)
	if !ok {
		return nil, err
	}

	return eventBox, nil
}

func GetLabel(app *gtk.Builder, id string) (*gtk.Label, error) {
	obj, err := app.GetObject(id)
	if err != nil {
		return nil, err
	}

	label, ok := obj.(*gtk.Label)
	if !ok {
		return nil, err
	}

	return label, nil
}
