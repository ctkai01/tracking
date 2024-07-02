package login

import (
	"fmt"
	"net/url"
	"tracking-go/gui/assets"
	"tracking-go/gui/layout"
	"tracking-go/gui/storage"
	"tracking-go/gui/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	fyneLayout "fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func LoginPage() *fyne.Container {
	appStorage := storage.GetAppStorage()

	application := appStorage.GetApplication()

	userName := widget.NewEntry()
	userName.SetPlaceHolder("UserName")
	password := widget.NewEntry()
	password.SetPlaceHolder("Password")
	password.Password = true

	urlForgot, _ := url.Parse("https://github.com")
	hyperLinkForgotPassword := widget.NewHyperlinkWithStyle("Forgot password?", urlForgot, fyne.TextAlignTrailing, fyne.TextStyle{Bold: true})

	urlSignUp, _ := url.Parse("https://google.com")
	hyperLinkSignUp := widget.NewHyperlinkWithStyle("SignUp", urlSignUp, fyne.TextAlignLeading, fyne.TextStyle{Bold: true})

	loginBtn := widget.NewButton("Login", func() {
		fmt.Println("basic auth : userName : ", userName.Text, " | password : ", password.Text)
		application.CurrentTab <- utils.HomeTab	})

	
	loginBtn.Alignment = widget.ButtonAlignCenter
	loginBtn.Importance = widget.HighImportance
	
	image := imageContainer()

	centerContainer :=  container.NewVBox(image,container.New(layout.NewPaddingLayout(10, 0, 0, 0), userName), container.New(layout.NewPaddingLayout(10, 10, 0, 0), password) , container.New(layout.NewPaddingLayout(0, 20, 0, 0), container.NewGridWithColumns(2, hyperLinkSignUp, hyperLinkForgotPassword) ), container.NewCenter(fyneLayout.NewSpacer(), loginBtn, fyneLayout.NewSpacer()))
	// return container.NewCenter(container.New(layout.NewPaddingLayout(0, 0, 100, 100), centerContainer)) 
	return container.New(layout.NewPaddingLayout(100, 0, 300, 300), centerContainer)
	// return container.New(layout.NewLoginLayout(fyneLayout.NewSpacer(), content), objs...)
}

// func centerContainer() fyne.CanvasObject {
// 	userName := widget.NewEntry()
// 	userName.SetPlaceHolder("UserName")
// 	password := widget.NewEntry()
// 	password.SetPlaceHolder("Password")
// 	password.Password = true

// 	forgotPasswordText := canvas.NewText("Forgot password?", utils.TransformColorFromHex("#136acb"))
// 	forgotPasswordText.TextStyle.Bold = true
// 	forgotPasswordText.Alignment = fyne.TextAlignTrailing

// 	loginBtn := widget.NewButton("Login", func() {
// 		fmt.Println("basic auth : userName : ", userName.Text, " | password : ", password.Text)
// 	})
// 	loginBtn.Alignment = widget.ButtonAlignCenter
// 	loginBtn.Importance = widget.HighImportance

// 	layoutContent := container.NewVBox(
// 		userName,
// 		// container.New(layout.NewPaddingLayout(0, 10, 0, 0), userName),

// 		// container.New(layout.NewPaddingLayout(0, 20, 0, 0), password),
		
// 		// container.New(layout.NewPaddingLayout(0, 10, 0, 0), forgotPasswordText),
		
// 		// loginBtn,

// 	)
// 	return layoutContent
// }

func imageContainer() fyne.CanvasObject {
	// https://static-e.live.fc2.com/img/logo/main_logo_new.png?20220722
	fc2Icon := assets.ResourceFc2Png
	// image := canvas.NewImageFromURI(fyne.URI {
		 
	// })

	image := canvas.NewImageFromResource(fc2Icon)

	image.FillMode = canvas.ImageFillOriginal
		// image.Resize(fyne.NewSize(float32(300), float32(300)))
		// image.Move(fyne.NewPos(float32(150), float32(150)))
	return image
}