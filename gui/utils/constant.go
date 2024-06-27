package utils

type ActiveTab string

const (
	LOGIN ActiveTab = "LOGIN"
	HOME  ActiveTab = "HOME"
)

type Tab struct {
	Title string
	Active ActiveTab
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
