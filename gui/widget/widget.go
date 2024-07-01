package mywidget

import (
	// "fmt"
	"tracking-go/gui/assets"
	"tracking-go/gui/storage"
	"tracking-go/gui/utils"
	"tracking-go/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type MyToolActionWidget struct {
	widget.BaseWidget
	ActionButton *widget.Button
	WorkingTime  *canvas.Text
	task       *models.Task
}

func NewMyToolActionWidget(workingTime string) *MyToolActionWidget {

	item := &MyToolActionWidget{
		ActionButton: widget.NewButtonWithIcon("", nil, nil),
		WorkingTime:  canvas.NewText(workingTime, utils.TransformColorFromHex("#800080")),
	}

	return item
}

func (item *MyToolActionWidget) CreateRenderer() fyne.WidgetRenderer {
	c := container.NewHBox(item.ActionButton, item.WorkingTime)
	return widget.NewSimpleRenderer(c)
}

func (item *MyToolActionWidget) SetTask(task *models.Task) {
	playIcon := assets.ResourcePlayPng
	pauseIcon := assets.ResourceSuspendPng

	// Init
	if task.Status == utils.PAUSE {
		item.ActionButton.SetIcon(playIcon)
	}

	if task.Status == utils.PLAYING {
		item.ActionButton.SetIcon(pauseIcon)
	}
	item.task = task

	item.ActionButton.OnTapped =
		func() {
			app := storage.GetAppStorage().GetApplication()
			// fmt.Printf("Play: %d\n", item.task.Id)
			if task.Status == utils.PAUSE {
				// fmt.Println("PAUSE OK")
				if app.ActiveTask != nil {
					if app.ActiveTask.Status == utils.PLAYING {
						return
					}
				}
				item.ActionButton.SetIcon(pauseIcon)
				item.task.Status = utils.PLAYING

				app.TaskEvent <- utils.TaskEvent{
					Task:  item.task,
					Type: utils.TASK_STATUS_CHANGE,
				}
				
			} else if task.Status == utils.PLAYING {
				item.ActionButton.SetIcon(playIcon)
				item.task.Status = utils.PAUSE
			}
			// fmt.Printf("Task: %v\n", task)
		}
	
		
}
