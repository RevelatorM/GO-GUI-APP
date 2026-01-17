package main

import (
	"fmt"
	"image"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
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
		//==================================
		var ops op.Ops
		var ed widget.Editor
		//==================================
		var list widget.List //creating list analog of Listbox in C# winforms
		list.Axis = layout.Vertical
		//==================================
		for { //this loops listens to the window`s events
			switch e := w.Event().(type) { //based on the vents types the code will dicide what to do using switch/case
			case app.DestroyEvent: //same as distructor in C++
				os.Exit(0)
			case app.FrameEvent: //is being called when the redraw is needed,updates the window
				gtx := app.NewContext(&ops, e) //e holds the window size for examp. ops is a list of operations

				// creating on Flex to group the items
				layout.Flex{
					Axis:      layout.Vertical,
					Alignment: layout.Middle, // making everything centered
				}.Layout(gtx,
					//==================================
					//Textbox
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions { //making textbox centered
							gtx.Constraints = layout.Exact(image.Pt(gtx.Dp(300), gtx.Dp(100)))
							e := material.Editor(th, &ed, "Type here")                                         //e := creating var to change Editor`s parameters
							e.TextSize = unit.Sp(30)                                                           //font size
							e.Editor.Filter = "0123456789abcdefghigklmnopqrstuvwxyzABCDEFGHIGKLMONPQRSTUVWXYZ" //filters the chracters that will be entered
							e.Editor.Alignment = text.Middle                                                   //making text inside centered
							e.Editor.MaxLen = 20                                                               //max amount of characters
							return e.Layout(gtx)
						})
					}),

					//==================================
					//List
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						
						gtx.Constraints.Max.Y = gtx.Dp(200)
						gtx.Constraints.Max.X = gtx.Dp(300)
					//===========Parameters of List must be above code of List "return"==================
						list.ScrollToEnd = true        //auto scroll
						list.Alignment = layout.Middle //position in middle

						return material.List(th, &list).Layout(gtx, 20, func(gtx layout.Context, i int) layout.Dimensions {
							return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								return material.Body1(th, fmt.Sprintf("Item %c", 'A'+i)).Layout(gtx)
							})
						})
					}),
				)

				//==================================
				//.Layout(gtx) is renders
				e.Frame(gtx.Ops) //draws called operations
			}
		}
	}()
	app.Main() //keeps the Event loop alive
}
