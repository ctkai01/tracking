package main

import (
	"fmt"
	// "time"
	"tracking-go/gui/assets"
	"tracking-go/gui/screens"
	"tracking-go/gui/storage"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	// "fyne.io/fyne/v2/widget"
	// "fyne.io/fyne/v2/app"
	// "fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/widget"
)

func main() {
	appStorage := storage.GetAppStorage()

	application := appStorage.GetApplication()
	mainWindow := application.MainWindow
	fc2Icon := assets.ResourceFc2Png

	mainWindow.SetIcon(fc2Icon)
	mainWindow.SetTitle("Tracking FC2")
	if desk, ok := application.AppCore.(desktop.App); ok {
		m := fyne.NewMenu("App", fyne.NewMenuItem("Show", func() {
			fmt.Println("Tapped Show System tray")
			mainWindow.Show()
		}))
		desk.SetSystemTrayMenu(m)
		desk.SetSystemTrayIcon(fc2Icon)
	}

	mainWindow.SetCloseIntercept(func() {
		mainWindow.Hide()
	})


	go func() {
		for tab := range application.CurrentTab {
			fmt.Println("Tab: ", tab)
			mainWindow.SetTitle(tab.Title)
			mainWindow.SetContent(screens.LoadPage(tab.Active))
		}
	}()

	application.MainWindow = mainWindow
	mainWindow.ShowAndRun()
}
