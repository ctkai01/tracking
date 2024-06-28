package home

import (
	"fmt"
	"image/color"

	// "net/url"
	"time"

	// "time"
	"tracking-go/gui/assets"
	"tracking-go/gui/layout"
	"tracking-go/gui/screens/home/timmer"
	"tracking-go/gui/storage"
	"tracking-go/gui/utils"
	"tracking-go/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	fyneLayout "fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	// "fyne.io/fyne/v2/widget"
)

var (
	app                                  = storage.GetAppStorage().GetApplication()
	hoursSection, minSection, secSection utils.Time
	hours, min, sec                      string
)

func HomePage() *fyne.Container {
	go resfreshTimeWorking()

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

var tasks = models.Tasks

func taskListContainer() *widget.List {
	list := widget.NewList(
		func() int {
			fmt.Println(len(tasks))
			return len(tasks)
		},
		func() fyne.CanvasObject {
			// return widget.NewLabel("template")
			// return TaskContainer
			return container.New(fyneLayout.NewVBoxLayout())
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			// o.(*widget.Label).SetText(fmt.Sprintf("%s: %s", tasks[i].ProjectName, tasks[i].TaskName))
			// o.(*widget.Label).Refresh()
			
			taskObj := TaskContainer(&tasks[i])
        // if o.(*fyne.Container) == nil {
        //     o = container.New(fyneLayout.NewVBoxLayout(), taskObj)
        // } else {
        //     o.(*fyne.Container).Add(taskObj)
        // }
        // o.(*fyne.Container).Refresh()
		o.(*fyne.Container).Objects = []fyne.CanvasObject{taskObj}

            // o.(*fyne.Container).Refresh()

			//  taskContainer := TaskContainer(&data[i])
            // o.(*fyne.Container).Objects = append(o.(*fyne.Container).Objects, taskContainer)
            // o.(*fyne.Container).Refresh()
		})
	return list
}

func resfreshTimeWorking() {
	for {

		hours, min, sec = utils.ConvertSecToHourMinSec(app.Counter)
		hoursSection.Refresh()
		minSection.Refresh()
		secSection.Refresh()
		time.Sleep(time.Second)
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

func TimeShowSection() fyne.CanvasObject {
	numberInit := 0
	sectionTime := timmer.NewTimeCounter(numberInit).Container

	// for {

	// 	numberInit++
	// 	time.Sleep(time.Second)
	// }

	return sectionTime

}

func infoUserContainer() *fyne.Container {
	name := canvas.NewText("harvet richard", color.Black)

	name.TextStyle.Bold = true

	return container.NewHBox(container.New(layout.NewPaddingLayout(0, 0, 10, 0), avatarContainer()), container.NewCenter(name))
}

func utilSection() *fyne.Container {
	reportIcon := assets.ResourceReportPng
	// reportImage := canvas.NewImageFromResource(reportIcon)
	// reportImage.FillMode = canvas.ImageFillContain

	// reportImage.SetMinSize(fyne.NewSize(float32(40), float32(40)))
	// text2 := canvas.NewText("2", color.Black)
	// reportText := canvas.NewText("Report", color.Black)

	// urlReport, _ := url.Parse("https://hustdemt.click")
	// reportLink := widget.NewHyperlinkWithStyle("Report", urlReport, fyne.TextAlignLeading, fyne.TextStyle{Bold: true})

	// reportSection := container.NewHBox(reportImage, reportLink)

	signoutIcon := assets.ResourceSignoutPng
	// signoutImage := canvas.NewImageFromResource(signoutIcon)
	// signoutImage.FillMode = canvas.ImageFillContain
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
