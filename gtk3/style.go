package gtk3

import (
	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

func setCssProvider() (*gtk.CssProvider, error) {
	cssProvider, err := gtk.CssProviderNew()
	if err != nil {
		return nil, err
	}

	// Load styles onto the provider
	err = cssProvider.LoadFromPath(CssFile)
	if err != nil {
		return nil, err
	}

	screen, err := gdk.ScreenGetDefault()
	if err != nil {
		panic(err)
	}
	gtk.AddProviderForScreen(screen, cssProvider, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)

	return cssProvider, nil
}

func setStyleClass(obj *gtk.Widget, className string) {

	// Get object style context
	styleCtx, err := obj.GetStyleContext()
	if err != nil {
		panic(err)
	}

	// Add desired class to the context
	styleCtx.AddClass(className)
}

func removeStyleClass(obj *gtk.Widget, className string) {

	// Get object style context
	styleCtx, err := obj.GetStyleContext()
	if err != nil {
		panic(err)
	}

	// Add desired class to the context
	styleCtx.RemoveClass(className)
}

/* TRANSPARENCY */

var alphaSupported bool

// TODO: Sometimes, the screen will flicker when we type too fast because this is quite heavy
func setTransparent(w *gtk.Window, ctx *cairo.Context) {

	if alphaSupported {

		// Alpha - being 0.0 - sets the background of the app as transparent
		ctx.SetSourceRGBA(0.0, 0.0, 0.0, 0.0)
	} else {
		ctx.SetSourceRGB(0.0, 0.0, 0.0)
	}

	ctx.SetOperator(cairo.OPERATOR_SOURCE)
	ctx.Paint()
	ctx.SetOperator(cairo.OPERATOR_OVER)
}

func screenChanged(window *gtk.Window) {
	screen, _ := window.GetScreen()
	visual, _ := screen.GetRGBAVisual()

	if visual != nil {
		alphaSupported = true
	} else {
		println("Alpha not supported")
		alphaSupported = false
	}

	window.SetVisual(visual)
}
