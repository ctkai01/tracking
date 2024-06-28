package timmer

import (
	// "fyne.io/fyne"
	"tracking-go/gui/utils"

	// "fyne.io/fyne/canvas"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	wg "fyne.io/x/fyne/widget"
)
var timeCounterSize = 2
type timeCounter struct {
	*fyne.Container
	digits []*wg.HexWidget
	current int

}

type timeSection struct {
	*fyne.Container
	Value *string
	TimeDigit *canvas.Text
}

func NewTimeSection(value *string) *timeSection {
	result := &timeSection{ Value: value}

	timeDigit := canvas.NewText(*result.Value, utils.TransformColorFromHex("#800080"))
	timeDigit.TextStyle.Bold = true
	timeDigit.TextSize = 20


	result.TimeDigit = timeDigit

	result.Container = container.NewVBox(timeDigit)

	return result
}

func (t *timeSection) Refresh() {
	t.TimeDigit.Text = *t.Value
	t.TimeDigit.Refresh()
}

func (t *timeSection) UI() *fyne.Container {
	return t.Container
}

func NewTimeCounter(init int) *timeCounter {
	numd := utils.GetDigits(init, timeCounterSize)
	d := make([]*wg.HexWidget, timeCounterSize)
	c := container.NewHBox()

	for i := 0; i < timeCounterSize; i++ {
		h := wg.NewHexWidget()
		h.SetSize(fyne.NewSize(15, 15))
		h.SetSlant(0)
		h.Set(numd[i])

		d[i] = h
		c.Add(h)
	}

	tc := &timeCounter{
		Container: c,
		digits:    d,
		current:   init,
	}

	return tc
}
// type timeSection struct {
// 	*fyne.Container
// 	Value *string
// 	TimeDigit *canvas.Text
// }

// func NewTimeSection(value *string) *timeSection {
// 	result := &timeSection{
// 		Value: value,
// 	}

// 	timeDigit := canvas.NewText(*result.Value)
// }