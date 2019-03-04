package glade

import "github.com/gotk3/gotk3/gtk"

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

func GetCheckButton(app *gtk.Builder, id string) (*gtk.CheckButton, error) {
	obj, err := app.GetObject(id)
	if err != nil {
		return nil, err
	}

	checkButton, ok := obj.(*gtk.CheckButton)
	if !ok {
		return nil, err
	}

	return checkButton, nil
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

func GetImage(app *gtk.Builder, id string) (*gtk.Image, error) {
	obj, err := app.GetObject(id)
	if err != nil {
		return nil, err
	}

	icon, ok := obj.(*gtk.Image)
	if !ok {
		return nil, err
	}

	return icon, nil
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

func GetNotebook(app *gtk.Builder, id string) (*gtk.Notebook, error) {
	obj, err := app.GetObject(id)
	if err != nil {
		return nil, err
	}

	notebook, ok := obj.(*gtk.Notebook)
	if !ok {
		return nil, err
	}

	return notebook, nil
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
