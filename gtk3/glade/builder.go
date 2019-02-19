package glade

import "github.com/gotk3/gotk3/gtk"

const WindowName = "window"
const BodyName = "body"
const InputBoxName = "input-box"
const InputName = "input"
const PrefsBtnName = "prefs_btn"

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

func GetWindow(app *gtk.Builder) (*gtk.Window, error) {
	obj, err := app.GetObject(WindowName)
	if err != nil {
		return nil, err
	}

	win, ok := obj.(*gtk.Window)
	if !ok {
		return nil, err
	}

	return win, nil
}

func GetBody(app *gtk.Builder) (*gtk.Box, error) {
	obj, err := app.GetObject(BodyName)
	if err != nil {
		return nil, err
	}

	body, ok := obj.(*gtk.Box)
	if !ok {
		return nil, err
	}

	return body, nil
}

func GetInput(app *gtk.Builder) (*gtk.Entry, error) {
	obj, err := app.GetObject(InputName)
	if err != nil {
		return nil, err
	}

	entry, ok := obj.(*gtk.Entry)
	if !ok {
		return nil, err
	}

	return entry, nil
}

func GetBtn(app *gtk.Builder) (*gtk.Button, error) {
	obj, err := app.GetObject(PrefsBtnName)
	if err != nil {
		return nil, err
	}

	btn, ok := obj.(*gtk.Button)
	if !ok {
		return nil, err
	}

	return btn, nil
}
