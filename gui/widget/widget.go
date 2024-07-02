package mywidget

import (
	// "fmt"
	"tracking-go/gui/assets"
	"tracking-go/gui/storage"
	"tracking-go/gui/utils"
	"tracking-go/models"

	"fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type MyToolActionWidget struct {
	widget.BaseWidget
	ActionButton *widget.Button
	Hours  *widget.Label
	Minutes  *widget.Label
	Seconds  *widget.Label
	task       *models.Task
}

func NewMyToolActionWidget(hours, minutes, seconds string) *MyToolActionWidget {
	hoursLabel := widget.NewLabel(hours)
	hoursLabel.Importance = widget.HighImportance

	minutesLabel := widget.NewLabel(minutes)
	minutesLabel.Importance = widget.HighImportance

	secondsLabel := widget.NewLabel(seconds)
	secondsLabel.Importance = widget.HighImportance
	item := &MyToolActionWidget{
		ActionButton: widget.NewButtonWithIcon("", nil, nil),
		Hours:  hoursLabel,
		Minutes: minutesLabel,
		Seconds: secondsLabel,
	}

	return item
}

func (item *MyToolActionWidget) CreateRenderer() fyne.WidgetRenderer {
	divideHour := widget.NewLabel(":")
	divideMinute := widget.NewLabel(":")
	divideHour.Importance = widget.HighImportance
	divideMinute.Importance = widget.HighImportance

	c := container.NewHBox(item.ActionButton, container.NewHBox(item.Hours, divideHour, item.Minutes, divideMinute, item.Seconds))
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
						// dialog.NewCustom()
						popup := dialog.NewCustom("Remind", "Close", container.NewVBox(
							widget.NewLabel("Only 1 task can run at the same time."),
							// widget.NewLabel("You can add any content here."),
						), app.MainWindow)
				
						// Show the popup dialog
						popup.Show()
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
