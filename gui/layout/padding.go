package layout

import (
	"fyne.io/fyne/v2"
)

type PaddingLayout struct {
	top    float32
	bottom float32
	right  float32
	left   float32
}

func NewPaddingLayout(top, bottom, right, left float32) fyne.Layout {
	return &PaddingLayout{top, bottom, right, left}
}

func (l *PaddingLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	pos := fyne.NewPos(l.right, l.top)

	// siz := fyne.NewSize(size.Width-2*padding, size.Height-2*padding)
	for _, child := range objects {
		siz := fyne.NewSize(size.Width-l.left-l.right, child.MinSize().Height)

		child.Resize(siz)
		child.Move(pos)
	}
}

func (l *PaddingLayout) MinSize(objects []fyne.CanvasObject) (min fyne.Size) {
	for _, child := range objects {
		if !child.Visible() {
			continue
		}

		min = min.Max(child.MinSize())
	}
	min = min.Add(fyne.NewSize(l.left+l.right, l.top+l.bottom))
	return
}
