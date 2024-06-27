package screens

import (
	"tracking-go/gui/screens/home"
	"tracking-go/gui/screens/login"
	"tracking-go/gui/utils"

	"fyne.io/fyne/v2"
)

func LoadPage(page utils.ActiveTab) *fyne.Container {
	switch page {
	case utils.LOGIN:
		return login.LoginPage()
	case utils.HOME:
		return home.HomePage()
	default:
		return login.LoginPage()
	}
}
