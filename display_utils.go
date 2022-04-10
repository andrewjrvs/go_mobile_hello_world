package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func New_Color_Button(label string, tapped func(), bColor color.Color) *fyne.Container {
	btn := widget.NewButton(label, tapped)
	btn_color := canvas.NewRectangle(bColor)
	rContainer := container.New(
		layout.NewMaxLayout(),
		btn_color,
		btn,
	)
	btn.Importance = widget.HighImportance
	return rContainer

}
