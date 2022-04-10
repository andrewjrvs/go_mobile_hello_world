//go:build darwin || linux || windows
// +build darwin linux windows

package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"path/filepath"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
)

type PlayerList struct {
	Players []Player `json:"players" yaml:"players" bson:"players"`
}

type Player struct {
	Status       int    `json:"status" yaml:"status" bson:"status"`
	Level        int    `json:"level" yaml:"level" bson:"level"`
	Experience   uint64 `json:"exp" yaml:"exp" bson:"exp"`
	Name         string `json:"name" yaml:"name" bson:"name"`
	NextLevelExp uint64 `json:"-" yaml:"-"`
}

var (
	topWindow         fyne.Window
	playerList        PlayerList
	activePlayerIndex binding.Int
	activePlayer      binding.ExternalUntyped
	infoPath          fyne.URI = storage.NewFileURI(filepath.Join(".", "info.json"))
)

func main() {
	t := flag.Bool("t", false, "Dark theme")
	flag.Parse()

	// //if !fyne.CurrentDevice().IsMobile() {
	// 	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	log.SetOutput(file)
	// //}

	prep_init()

	// for i := 1; i < 10000; i += 20 {
	// 	lvl := Find_Level_From_Experience(uint64(i))
	// 	log.Printf("ex %d [%d] %d-%d", i, lvl, Get_Exp_Base_for_Level(lvl), Get_Exp_Base_for_Level(lvl+1))
	// }

	log.Println("Application Starting")

	//activePlayer = binding.BindUntyped(&Player{Status: 0, Level: 1, Experience: 5, Name: "NoMan"})
	activePlayerIndex.AddListener(binding.NewDataListener(func() {
		actIndx, _ := activePlayerIndex.Get()
		activePlayer = binding.BindUntyped(&playerList.Players[actIndx])
	}))

	myApp := app.NewWithID("us.cvbn.button_quest")
	if *t {
		myApp.Settings().SetTheme(theme.DarkTheme())
	} else {
		myApp.Settings().SetTheme(theme.LightTheme())
	}

	myApp.SetIcon(theme.FyneLogo())
	//logLifecycle(myApp)

	topWindow := myApp.NewWindow("Button Quest")

	//topWindow.Resize(fyne.NewSize(400, 400))

	topWindow.SetMainMenu(GetMainMenu(&myApp, &topWindow))

	topWindow.SetMaster()
	content := container.NewMax()

	content.Objects = []fyne.CanvasObject{welcomeScreen(topWindow, activePlayer)}
	content.Refresh()

	go func() {
		time.Sleep(5 * time.Second)
		content.Objects = []fyne.CanvasObject{outsideScreen(topWindow, activePlayer)}
		content.Refresh()
	}()

	// Close the App when Escape key is pressed
	topWindow.Canvas().SetOnTypedKey(func(keyEvent *fyne.KeyEvent) {

		if keyEvent.Name == fyne.KeyEscape {
			myApp.Quit()
		}
	})

	//topWindow.SetContent(makeNav(setTutorial, false))
	topWindow.SetContent(content)
	topWindow.Resize(fyne.NewSize(640, 460))

	myApp.Lifecycle().SetOnStopped(func() {
		write, err := storage.SaveFileToURI(infoPath)
		if err != nil {
			fyne.LogError("Unable to save file \"info.json\"", err)
		}

		data, err := json.Marshal(playerList)
		if err != nil {
			fyne.LogError("problem conversting to JSON", err)
		}
		_, err = write.Write(data)
		if err != nil {
			log.Panicln("Failed to write temporary file")
		}

		write.Close()
	})

	// Show window and run app
	topWindow.ShowAndRun()
}

func logLifecycle(a fyne.App) {
	a.Lifecycle().SetOnStarted(func() {
		log.Println("Lifecycle: Started")
	})
	a.Lifecycle().SetOnStopped(func() {
		log.Println("Lifecycle: Stopped")
	})
	a.Lifecycle().SetOnEnteredForeground(func() {
		log.Println("Lifecycle: Entered Foreground")
	})
	a.Lifecycle().SetOnExitedForeground(func() {
		log.Println("Lifecycle: Exited Foreground")
	})
}

func prep_init() {
	activePlayerIndex = binding.NewInt()

	b, br := storage.CanRead(infoPath)
	log.Printf("canRead %#v %#v", b, br)
	// attempt to load from file
	read, err := ioutil.ReadFile("info.json")
	log.Printf("err %#v", err)
	fileLoaded := false
	log.Println("prep")
	if err == nil {

		err = json.Unmarshal(read, &playerList)
		log.Printf("err2 %#v", err)
		if err == nil {
			fileLoaded = true
		}
		//read.Close()
	}
	log.Printf("playerList %#v", playerList)
	if !fileLoaded {
		playerList = PlayerList{Players: []Player{
			{Status: 0, Level: 1, Experience: 5, Name: "NoMan"},
		}}
	}

	if len(playerList.Players) > 0 {
		activePlayerIndex.Set(0)
	}

}
