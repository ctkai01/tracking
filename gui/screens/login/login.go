package login

import (
	"fmt"
	"tracking-go/gui/layout"
	"tracking-go/gui/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	fyneLayout "fyne.io/fyne/v2/layout"
)

func LoginPage() *fyne.Container {
	content :=  centerContainer()

	objs := []fyne.CanvasObject{fyneLayout.NewSpacer(), content}

	return container.New(layout.NewLoginLayout(fyneLayout.NewSpacer(), content), objs...)
}

func centerContainer() fyne.CanvasObject {
	userName := widget.NewEntry()
	userName.SetPlaceHolder("UserName")
	// userName.SET(fyne.NewSize(800, userName.MinSize().Height))
	password := widget.NewEntry()
	password.SetPlaceHolder("Password")
	password.Password = true
	// password.Resize(fyne.NewSize(800, password.MinSize().Height))

	forgotPasswordText := canvas.NewText("Forgot password?", utils.TransformColorFromHex("#136acb"))
	forgotPasswordText.TextStyle.Bold = true
	forgotPasswordText.Alignment = fyne.TextAlignTrailing

	loginBtn := widget.NewButton("Login", func() {
		fmt.Println("basic auth : userName : ", userName.Text, " | password : ", password.Text)
	})
	loginBtn.Alignment = widget.ButtonAlignCenter
	loginBtn.Importance = widget.HighImportance

	layoutContent := container.NewVBox(
		container.New(layout.NewPaddingLayout(0, 10, 0, 0), userName),

		container.New(layout.NewPaddingLayout(0, 20, 0, 0), password),
		
		container.New(layout.NewPaddingLayout(0, 10, 0, 0), forgotPasswordText),
		
		loginBtn,

	)
	return layoutContent
}