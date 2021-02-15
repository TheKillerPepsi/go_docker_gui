package main

//Tutorial: https://developer.fyne.io/tour/basics/windows.html

import (
	"fmt"
	//"time"

	//"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	//"fyne.io/fyne/v2/canvas"
	//"fyne.io/fyne/v2/theme"
)

func main() {
	main_window()
	tidyUp()
}

func main_window() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello")
	myWindow.SetContent(widget.NewLabel("Hello World!"))

	// show and run the application
	myWindow.ShowAndRun()
	//ShowANdRUn is short version
	//myApp.Run()

}

func tidyUp() {
	fmt.Println("Exited")
}

//func first_window()

