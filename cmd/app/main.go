package main

import (
	"fmt"
	"time"
	"tracking-go/gui/screens"
	"tracking-go/gui/storage"
	// "fyne.io/fyne/v2/widget"
	// "fyne.io/fyne/v2/app"
	// "fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/widget"
)

func main() {
	appStorage := storage.GetAppStorage()

	application := appStorage.GetApplication()
	mainWindow := application.MainWindow

	mainWindow.SetTitle("Tracking")

	go func() {
		for tab := range application.CurrentTab {
			fmt.Println("Tab: ", tab)
			mainWindow.SetTitle(tab.Title)
			mainWindow.SetContent(screens.LoadPage(tab.Active))
		}
	}()

	go func ()  {
		for {
			time.Sleep(time.Second)
			application.Counter++ 
		}
		
	}()
	application.MainWindow = mainWindow
	mainWindow.ShowAndRun()
}
