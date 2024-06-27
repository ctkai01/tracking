package home

import (
	"image/color"
	"tracking-go/gui/screens/home/timmer"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	fyneLayout "fyne.io/fyne/v2/layout"
	// "fyne.io/fyne/v2/widget"
)

func HomePage() *fyne.Container {
	workingTimeTitle := canvas.NewText("Working Time", color.Black)

	

	timeSection := container.NewVBox(workingTimeTitle, TimeShowSection())
	text2 := canvas.NewText("2", color.Black)
	// 
	header := container.New(fyneLayout.NewGridLayout(2), timeSection, text2)
    // Create a container for the button
    content := container.NewVBox(
        header,
    )

    return content
}

func TimeShowSection() fyne.CanvasObject {
	sectionTime := timmer.NewTimeCounter(50).Container

	return sectionTime

}