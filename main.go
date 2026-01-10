package main

import (
	"os"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

func main() {
	go func() { //creating go routing func,routing func in GO is a thread
		w := new(app.Window)                           //creating a window
		w.Option(app.Title("App"))                     //giving a title
		w.Option(app.Size(unit.Dp(500), unit.Dp(500))) //setting up the start size of the window
		th := material.NewTheme()                      //creating a theme th is a variable
		var ops op.Ops                                 //holds the operations,it is a recording buffer
		for {                                          //this loops listens to the window`s events
			switch e := w.Event().(type) { //based on the vents types the code will dicide what to do using switch/cae
			case app.DestroyEvent: //same as distructor in C++
				os.Exit(0)
			case app.FrameEvent: //is being called when the redraw is needed,updates the window
				gtx := app.NewContext(&ops, e)         //e holds the window size f examp. ops is a list of operations
				material.Label(th, 16, "").Layout(gtx) //th - theme,16 - text size,"" - text in the label
				//.Layout(gtx) is drawing the text
				e.Frame(gtx.Ops) //draws called operations
			}
			//============================
		}
	}()
	app.Main() //keeps the Event loop alive
}
