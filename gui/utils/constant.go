package utils

import "tracking-go/models"

type ActiveTab string

const (
	LOGIN ActiveTab = "LOGIN"
	HOME  ActiveTab = "HOME"

	HOUR = 1
	MINUTE = 2
	SECOND = 3

	PAUSE = 1
	PLAYING = 2

	TASK_STATUS_CHANGE = "TASK_STATUS_CHANGE"

)

type Tab struct {
	Title string
	Active ActiveTab
}

type TaskEvent struct {
	Task *models.Task
	Type string
	// PreActiveTask *models.Task
}

var (
	LoginTab  = Tab {
		Title: "Login",
		Active: LOGIN,
	}
	HomeTab = Tab {
		Title: "Home",
		Active: HOME,
	}

)
