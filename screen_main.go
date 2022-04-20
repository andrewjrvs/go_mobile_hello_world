package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
)

func mainScreen(w fyne.Window, player binding.ExternalUntyped, playerList *PlayerList) fyne.CanvasObject {

	bndPlayerList := binding.BindUntyped(&playerList)

	return container.NewVBox(
		NewHeroList(bndPlayerList, activePlayerIndex),
		outsideScreen(w, player),
	)
}
