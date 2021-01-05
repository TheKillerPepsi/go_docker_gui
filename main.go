package main

//Tutorial: https://developer.fyne.io/tour/basics/windows.html

import (
	"fmt"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello")
	myWindow.SetContent(widget.NewLabel("Hello"))

	secWindow := myApp.NewWindow("Hello again!!")
	secWindow.SetContent(widget.NewLabel("HAIII"))

	myWindow.Show()
	secWindow.Show()
	myApp.Run()
	tidyUp()
}

func tidyUp() {
	fmt.Println("Exited")
}

func first_window()
