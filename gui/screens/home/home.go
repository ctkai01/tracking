package home

import (
	"fmt"
	"image/color"

	// "net/url"
	"time"

	// "time"
	"tracking-go/gui/assets"
	"tracking-go/gui/layout"
	"tracking-go/gui/storage"
	"tracking-go/gui/utils"
	mywidget "tracking-go/gui/widget"
	"tracking-go/gui/widget/timmer"
	"tracking-go/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	fyneLayout "fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	// "fyne.io/fyne/v2/widget"
)

var (
	app                                  = storage.GetAppStorage().GetApplication()
	hoursSection, minSection, secSection utils.Time
	hours, min, sec                      string
	ticker                               = time.NewTicker(1 * time.Second)
)

func HomePage() *fyne.Container {
	go resfreshTimeWorking()
	go func() {
		for e := range app.TaskEvent {
			switch e.Type {
			case utils.TASK_STATUS_CHANGE:
				fmt.Printf("Task %d status changed\n", e.Task.Id)
				app.ActiveTask = e.Task
				// Handle task status change event
				// Example: Update UI or perform other actions based on the event
			}
		}
	}()
	workingTimeTitle := canvas.NewText("Working Time", color.Black)
	workingTimeTitle.TextStyle.Bold = true

	timeSection := container.NewVBox(workingTimeTitle, countingTimerSection())

	infoSection := container.NewVBox(container.New(layout.NewPaddingLayout(0, 10, 0, 0), infoUserContainer()), utilSection())
	//
	header := container.New(fyneLayout.NewGridLayout(2), timeSection, infoSection)
	// Create a container for the button
	content := container.NewBorder(
		container.New(layout.NewPaddingLayout(20, 0, 20, 20), header),
		nil, nil, nil, taskListContainer(),
	)

	return content
}

func renderListLine() fyne.CanvasObject {
	t := mywidget.NewMyToolActionWidget("")
	// fmt.Println("T: ", t)
	// return container.New(fyneLayout.NewBorderLayout(nil, nil, nil, t), t, widget.NewLabel(""))

	return container.New(fyneLayout.NewGridLayout(4), widget.NewLabel(""), fyneLayout.NewSpacer(), fyneLayout.NewSpacer(), t)
}

func bindDataToListLine() func(di binding.DataItem, co fyne.CanvasObject) {
	return func(di binding.DataItem, co fyne.CanvasObject) {
		t := models.NewTaskFromDataItem(di)
		container := co.(*fyne.Container)

		label := container.Objects[0].(*widget.Label)
		toolAction := container.Objects[3].(*mywidget.MyToolActionWidget)

		toolAction.SetTask(t)
		nameTask := t.NameTask()
		label.Truncation = fyne.TextTruncateOff
		label.Wrapping = fyne.TextWrapOff
		label.Bind(binding.BindString(&nameTask))

		_, _, workingTime := utils.ConvertSecToHourMinSec(t.WorkingTime)
		toolAction.WorkingTime.Text = workingTime

	}
}

func taskListContainer() *widget.List {
	tasks := models.NewTaskData(models.Tasks)

	list := widget.NewListWithData(tasks, renderListLine, bindDataToListLine())

	return list
}

func resfreshTimeWorking() {
	for {

		if hoursSection != nil && minSection != nil && secSection != nil {
			hours, min, sec = utils.ConvertSecToHourMinSec(app.Counter)
			hoursSection.Refresh()
			minSection.Refresh()
			secSection.Refresh()
			time.Sleep(time.Second)
		}

	}
}

func countingTimerSection() *fyne.Container {
	hours, min, sec = utils.ConvertSecToHourMinSec(app.Counter)

	hoursSection = timmer.NewTimeSection(&hours)
	minSection = timmer.NewTimeSection(&min)
	secSection = timmer.NewTimeSection(&sec)

	divideHour := canvas.NewText(":", utils.TransformColorFromHex("#800080"))

	divideHour.TextStyle.Bold = true
	divideHour.TextSize = 18

	divideMinute := canvas.NewText(":", utils.TransformColorFromHex("#800080"))
	divideMinute.TextStyle.Bold = true
	divideMinute.TextSize = 18

	return container.NewHBox(hoursSection.UI(), divideHour, minSection.UI(), divideMinute, secSection.UI())
}

func infoUserContainer() *fyne.Container {
	name := canvas.NewText("harvet richard", color.Black)

	name.TextStyle.Bold = true

	return container.NewHBox(container.New(layout.NewPaddingLayout(0, 0, 10, 0), avatarContainer()), container.NewCenter(name))
}

func utilSection() *fyne.Container {
	reportIcon := assets.ResourceReportPng

	signoutIcon := assets.ResourceSignoutPng
	reportBtn := widget.NewButtonWithIcon("Report", reportIcon, func() {
		fmt.Println("Report")
	})

	signOutBtn := widget.NewButtonWithIcon("Sigout", signoutIcon, func() {
		fmt.Println("Signout")
	})
	return container.New(fyneLayout.NewGridLayout(2), reportBtn, signOutBtn)

}

func avatarContainer() fyne.CanvasObject {
	defaultAvatar := assets.ResourceDefaultUserPng

	image := canvas.NewImageFromResource(defaultAvatar)

	// image.FillMode = canvas.ImageFillContain

	image.SetMinSize(fyne.NewSize(float32(40), float32(40)))

	return image
}
