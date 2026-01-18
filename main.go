package main

import (
	"fmt"
	"image"
	"os"

	"database/sql"

	_ "modernc.org/sqlite"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// ==================================
// GPU structure
type GPU struct {
	Name string
	Vram string
	MHz  string
	VMHz string
}

func main() {
	//==================================
	database, err := sql.Open("sqlite", "./database.db") //creating db
	if err != nil {
		fmt.Println(err)
		return
	}

	//===================
	stmt, _ := database.Prepare(
		"CREATE TABLE IF NOT EXISTS gpus (id INTEGER PRIMARY KEY, gpuname TEXT, vram TEXT, MHz TEXT, vmhz TEXT)",
	)
	stmt.Exec() //creating table

	//filling db
	stmt, _ = database.Prepare("INSERT INTO gpus(gpuname,vram,MHz,vmhz) VALUES (?,?,?,?)")
	stmt.Exec("RTX 3060", "12GB", "1867 MHz chip", "15000 MHz VRAM")
	stmt.Exec("RTX 3070", "8GB", "1815 MHz chip", "14000 MHz VRAM")
	stmt.Exec("RTX 3080", "10GB", "1710 MHz chip", "19000 MHz VRAM")
	stmt.Exec("RTX 4060", "8GB", "2535 MHz chip", "17000 MHz VRAM")
	stmt.Exec("RTX 4070", "12GB", "2550 MHz chip", "21000 MHz VRAM")
	stmt.Exec("RX 7600", "8GB", "2755 MHz chip", "18000 MHz VRAM")
	stmt.Exec("RX 9070 XT", "16GB", "3030 MHz chip", "20000 MHz VRAM")
	stmt.Exec("Arc B570", "10GB", "2660 MHz chip", "19000 MHz VRAM")

	//==================================
	go func() { //creating go routing func,routing func in GO is a thread
		w := new(app.Window)                           //creating a window
		w.Option(app.Title("App"))                     //giving a title
		w.Option(app.Size(unit.Dp(500), unit.Dp(500))) //setting up the start size of the window
		th := material.NewTheme()                      //creating a theme th is a variable

		//==================================
		var ops op.Ops
		var ed widget.Editor
		var searchBtn widget.Clickable // button

		var results []GPU // data for List

		//==================================
		var list widget.List //creating list analog of Listbox in C# winforms
		list.Axis = layout.Vertical

		//==================================
		for {
			switch e := w.Event().(type) {
			case app.DestroyEvent:
				os.Exit(0)

			case app.FrameEvent:
				gtx := app.NewContext(&ops, e)

				//==================================
				// Button click handling
				if searchBtn.Clicked(gtx) {
					query := ed.Text()
					results = results[:0]

					rows, err := database.Query(
						"SELECT gpuname, vram, MHz, vmhz FROM gpus WHERE gpuname LIKE ?",
						"%"+query+"%",
					)
					if err == nil {
						for rows.Next() {
							var gpu GPU
							rows.Scan(&gpu.Name, &gpu.Vram, &gpu.MHz, &gpu.VMHz)
							results = append(results, gpu)
						}
						rows.Close()
					}
				}

				//==================================
				// creating on Flex to group the items
				layout.Flex{
					Axis:      layout.Vertical,
					Alignment: layout.Middle,
				}.Layout(gtx,

					//==================================
					//Textbox
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							gtx.Constraints = layout.Exact(image.Pt(gtx.Dp(300), gtx.Dp(80)))
							e := material.Editor(th, &ed, "Type GPU name")
							e.TextSize = unit.Sp(24)
							e.Editor.Filter = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ "
							e.Editor.Alignment = text.Middle
							e.Editor.MaxLen = 20
							return e.Layout(gtx)
						})
					}),

					//==================================
					//Button
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							btn := material.Button(th, &searchBtn, "Search")
							return btn.Layout(gtx)
						})
					}),

					//==================================
					//List
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						gtx.Constraints.Max.Y = gtx.Dp(220)
						gtx.Constraints.Max.X = gtx.Dp(350)

						return material.List(th, &list).Layout(gtx, len(results), func(gtx layout.Context, i int) layout.Dimensions {
							gpu := results[i]

							return layout.UniformInset(unit.Dp(8)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								return layout.Flex{
									Axis: layout.Vertical,
								}.Layout(gtx,
									layout.Rigid(material.Body1(th, gpu.Name).Layout),
									layout.Rigid(material.Body2(th, "VRAM: "+gpu.Vram).Layout),
									layout.Rigid(material.Body2(th, gpu.MHz).Layout),
									layout.Rigid(material.Body2(th, gpu.VMHz).Layout),
								)
							})
						})
					}),
				)

				//==================================
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}
