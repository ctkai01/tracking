package utils

import (
	"fyne.io/fyne/v2"
)
type Time interface {
	Refresh()
	UI()  *fyne.Container
}