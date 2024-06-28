package home

import (
	"fmt"
	"tracking-go/gui/utils"
	"tracking-go/models"

	"fyne.io/fyne/v2"
	fyneLayout "fyne.io/fyne/v2/layout"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)


type CustomList struct {
	fyne.Widget
	
}

func TaskContainer(task *models.Task) *fyne.Container {
	name := canvas.NewText(fmt.Sprintf("%s: %s", task.ProjectName, task.TaskName), utils.TransformColorFromHex("#800080"))

	start := canvas.NewText("OK", utils.TransformColorFromHex("#800080"))

	return container.New(fyneLayout.NewGridLayout(2), name, start)
}