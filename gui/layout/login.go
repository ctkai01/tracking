package layout

import (
	"fyne.io/fyne/v2"
)

const SIDE_WIDTH_LOGIN = 20

type LoginLayout struct {
	top, content fyne.CanvasObject
}

func NewLoginLayout(top, content fyne.CanvasObject) fyne.Layout {
	return &LoginLayout { top: top, content: content }
}

func (l *LoginLayout) Layout(_ []fyne.CanvasObject, size fyne.Size) {
	l.content.Move(fyne.NewPos(size.Width / 2 - (l.content.MinSize().Width / 2), float32(60) + float32(60)))
	l.content.Resize(fyne.NewSize(l.content.MinSize().Width, l.content.MinSize().Height))
}

func (l *LoginLayout) MinSize(object []fyne.CanvasObject) fyne.Size {
	borders := fyne.NewSize(SIDE_WIDTH_LOGIN * 2, l.top.MinSize().Height)
	return borders
}