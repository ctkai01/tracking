package storage

import (
	"sync"
	"tracking-go/gui/utils"
	"tracking-go/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type Application struct {
	AppCore    fyne.App
	MainWindow fyne.Window
	CurrentTab chan utils.Tab
	Counter int
	TaskEvent chan utils.TaskEvent
	ActiveTask *models.Task
}

type AppStorage struct {
	mu          sync.RWMutex
	application *Application
}

var state *AppStorage
var once sync.Once

func GetAppStorage() *AppStorage {
	once.Do(func() {
		if state == nil {
			state = &AppStorage{}

			state.setDefaultApplication(("Tracking"))
		}
		
	})

	return state
}

func (s *AppStorage) setDefaultApplication(window string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	a := app.New()
	mainWindow := a.NewWindow(window)
	mainWindow.Resize(fyne.NewSize(900, 500))
	s.application = &Application{
		AppCore:    a,
		MainWindow: mainWindow,
		CurrentTab: make(chan utils.Tab, 1),
		TaskEvent: make(chan utils.TaskEvent, 1),
	}
	
	s.application.CurrentTab <- utils.HomeTab

	// s.application.CurrentTab <- utils.LoginTab

}

func (s *AppStorage) GetApplication() *Application {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.application
}
