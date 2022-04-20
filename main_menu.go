package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func GetMainMenu(app *fyne.App, w *fyne.Window, fn func(string)) *fyne.MainMenu {
	// Main menu
	fileMenu := fyne.NewMenu("File",
		fyne.NewMenuItem("Quit", func() { topWindow.Close() }),
	)

	helpMenu := fyne.NewMenu("Help",
		fyne.NewMenuItem("About", func() {
			dialog.ShowCustom("About", "Close", container.NewVBox(
				widget.NewLabel("Button Quest for Go"),
				widget.NewLabel("Version: v0.1"),
				widget.NewLabel("Author: Jarvis"),
			), *w)
		}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Other", func() {
			fn("other")
		}),
		fyne.NewMenuItem("debug", func() {
			fn("debug")
		}),
		fyne.NewMenuItem("home", func() {
			fn("home")
		}),
		fyne.NewMenuItem("save", func() {
			fn("save")
		}),
	)
	mainMenu := fyne.NewMainMenu(
		fileMenu,
		helpMenu,
	)

	return mainMenu
}
