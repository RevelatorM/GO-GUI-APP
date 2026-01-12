package main

import (
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func main() {
	go func() { //creating go routing func,routing func in GO is a thread
		w := new(app.Window)                           //creating a window
		w.Option(app.Title("App"))                     //giving a title
		w.Option(app.Size(unit.Dp(500), unit.Dp(500))) //setting up the start size of the window
		th := material.NewTheme()                      //creating a theme th is a variable
		var ops op.Ops
		var ed widget.Editor
		for { //this loops listens to the window`s events
			switch e := w.Event().(type) { //based on the vents types the code will dicide what to do using switch/case
			case app.DestroyEvent: //same as distructor in C++
				os.Exit(0)
			case app.FrameEvent: //is being called when the redraw is needed,updates the window
				gtx := app.NewContext(&ops, e) //e holds the window size f examp. ops is a list of operations
				//==================================
				layout.Flex{Axis: layout.Vertical}.Layout( //creating Textbox
					gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						e := material.Editor(th, &ed, "Type here")                                         //e := creating var to change Editor`s parameters
						e.TextSize = unit.Sp(30)                                                           //font size
						e.Editor.Filter = "0123456789abcdefghigklmnopqrstuvwxyzABCDEFGHIGKLMONPQRSTUVWXYZ" //filters the chracters that will be entered
						return e.Layout(gtx)
					}),
				)
				//==================================
				//.Layout(gtx) is renders
				e.Frame(gtx.Ops) //draws called operations
			}
			//============================
		}
	}()
	app.Main() //keeps the Event loop alive
}
