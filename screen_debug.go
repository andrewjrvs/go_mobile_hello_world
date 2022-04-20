package main

import (
	"bufio"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func debugScreen(w fyne.Window) fyne.CanvasObject {
	logPath := fyne.CurrentApp().Preferences().StringWithFallback("logFile", "app.log")

	fDt, err := fyne.CurrentApp().Storage().Open(logPath)

	if err != nil {
		return container.NewMax(
			canvas.NewText("Debug Logs", theme.ForegroundColor()),
			widget.NewLabel(string(err.Error())),
		)
	}

	defer fDt.Close()

	debugData := binding.NewStringList()

	go func() {
		r := bufio.NewReader(fDt)
		c, p, e := r.ReadLine()
		for e == nil {
			debugData.Prepend(string(c))
			if p {
				c, p, e = r.ReadLine()
				debugData.Prepend(string(c))
			}
			c, p, e = r.ReadLine()
		}
	}()

	list := widget.NewListWithData(debugData,
		func() fyne.CanvasObject {
			el := widget.NewLabel("template")
			el.Wrapping = fyne.TextWrapBreak
			return el
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		})

	return container.NewMax(
		canvas.NewText("Debug Logs", theme.ForegroundColor()),
		list,
	)
}
