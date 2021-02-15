package main

//Tutorial: https://developer.fyne.io/tour/basics/windows.html

import (
	//"fmt"
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	//"fyne.io/fyne/v2/layout"
)

func main() {
	//mainWindow()
	secWindows()

	tidyUp()
}

func mainWindow() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello Canvas!")
	//myCanvas := myWindow.Canvas()

	//start with define the canvas
	text1 := canvas.NewText("Hello World!", color.White)
	text1.TextStyle.Italic = true
	text2 := canvas.NewText("Hello Moon?!", color.Opaque)
	text2.Move(fyne.NewPos(20, 20))
	text2.TextStyle.Bold = true
	text3 := canvas.NewText("Heyyyy Pepsi!", color.Black)
	text3.TextStyle.Monospace = true
	text3.Move((fyne.NewPos(40, 50)))

	content := container.NewWithoutLayout(text1, text2, text3)
	//content := container.New(layout.NewGridLayout(2), text1, text2, text3)
	//myCanvas.SetContent(text)
	//myWindow.SetContent(widget.NewLabel("Hello World!"))

	// show and run the application
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
	//ShowANdRUn is short version
	//myApp.Run()

}

func tidyUp() {
	fmt.Println("Exited")
}

//func first_window()

func secWindows() {

	myApp := app.New()
	myWindow := myApp.NewWindow("Widget")

	myWindow.SetContent(widget.NewEntry())
	myWindow.ShowAndRun()

}
