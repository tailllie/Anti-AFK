package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/go-vgo/robotgo"
	"math/rand"
	"time"
)

func main() {
	gui()
}

func gui() {
	App := app.New()
	Window := App.NewWindow("AntiAFK by taillie")

	sleep := 5.0

	Slider := widget.NewSlider(1, 60)
	Slider.Step = 1
	Slider.Value = 5

	Label := widget.NewLabelWithStyle("5s", fyne.TextAlignLeading, fyne.TextStyle{})

	Slider.OnChanged = func(value float64) {
		sleep = value
		Label.SetText(fmt.Sprintf("%.0fs", value))
	}

	var Button *widget.Button // Объявление переменной Button внутри функции gui()

	Button = widget.NewButton("Start", func() {
		if Button.Text == "Start" {
			Button.SetText("Stop")
			go mouse(sleep)
		} else {
			Button.SetText("Start")
		}
	})

	content := container.NewVBox(
		container.NewGridWithColumns(2,
			widget.NewLabel("Duration"),
			Label,
			Slider,
		),
		Button,
	)
	Window.SetContent(content)
	Window.ShowAndRun()
}

func mouse(sleep float64) {
	rand.Seed(time.Now().UnixNano())
	scrW, scrH := robotgo.GetScreenSize()

	for {
		x := rand.Intn(scrW)
		y := rand.Intn(scrH)

		robotgo.MoveMouseSmooth(x, y)

		time.Sleep(time.Duration(sleep) * time.Second)
	}
}

/** func keyboard(sleep float64) {
	rand.Seed(time.Now().UnixNano())

	for {
		keyIndex := rand.Intn(4)
		var key string
		switch keyIndex {
		case 0:
			key = "w"
		case 1:
			key = "a"
		case 2:
			key = "s"
		case 3:
			key = "d"
		}

		robotgo.KeyTap(key)

		time.Sleep(time.Duration(sleep) * time.Second)
	}
} **/
