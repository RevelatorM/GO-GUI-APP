package main

import (
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// ==================================
func run(window *app.Window) error {
	var nameInput widget.Editor //textbox
	nameInput.SingleLine = true
	nameInput.Submit = true //catches Enter button
	theme := material.NewTheme()
	var ops op.Ops
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			// This graphics context is used for managing the rendering state.
			gtx := app.NewContext(&ops, e)

			// Define an large label with an appropriate text:
			title := material.H1(theme, "GPU Finder")

			// Change the color of the label.
			maroon := color.NRGBA{R: 100, G: 50, B: 50, A: 255}
			title.Color = maroon

			// Change the position of the label.
			title.Alignment = text.Middle

			// Draw the label to the graphics context.
			title.Layout(gtx)

			// Pass the drawing operations to the GPU.
			e.Frame(gtx.Ops)
		}
	}
}

// ==================================
func main() {
	go func() {
		window := new(app.Window) //creating window
		err := run(window)        //calling run
		if err != nil {           //if run function returns error
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}
