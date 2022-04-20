package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

func FirstLoad() (bool, bool) {
	log.Println("FirstLoad")
	playerCreated := false
	settingsLoaded := false

	// setup application
	fyne.CurrentApp().SetIcon(theme.FyneLogo())

	// load users
	file, opnError := fyne.CurrentApp().Storage().Open(fyne.CurrentApp().Preferences().StringWithFallback("PlayerPath", DefaultPlayerPath))

	if opnError != nil {
		log.Println("Error", opnError)

		newFileFirstLoad()
	} else {
		//fContent []byte = nil
		defer file.Close()

		dat, readErr := ioutil.ReadAll(file)
		if readErr == nil {
			merr := json.Unmarshal(dat, &playerList)

			if merr != nil {
				log.Printf("Unable to parse player Info %#v", merr)
			} else {
				playerCreated = true
			}
		} else {
			log.Println("Read error", readErr)
		}

	}

	return playerCreated, settingsLoaded
}

func newFileFirstLoad() {
	log.Println("newFile")
	file, cError := fyne.CurrentApp().Storage().Create(fyne.CurrentApp().Preferences().StringWithFallback("PlayerPath", DefaultPlayerPath))

	if cError != nil {
		log.Println("Error: C22: ", cError)
		return
	}
	defer file.Close()
	return
}
