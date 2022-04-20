package main

import (
	"encoding/json"
	"log"

	"fyne.io/fyne/v2"
)

func BeforeAppDestroy() {

	log.Println("BeforeAppDestoryed")

	rErr := fyne.CurrentApp().Storage().Remove(fyne.CurrentApp().Preferences().StringWithFallback("PlayerPath", DefaultPlayerPath))

	if rErr != nil {
		log.Println("Unable to remove player list", rErr)
	}

	file, cError := fyne.CurrentApp().Storage().Create(fyne.CurrentApp().Preferences().StringWithFallback("PlayerPath", DefaultPlayerPath))

	if cError != nil {
		log.Println("Error", cError)
		return
	}
	defer file.Close()

	data, err := json.Marshal(playerList)
	if err != nil {
		fyne.LogError("problem converting to JSON", err)
	}
	_, err = file.Write(data)
	if err != nil {
		log.Println("Failed to write temporary file")
	}
}
