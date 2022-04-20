//go:build darwin || linux || windows
// +build darwin linux windows

package main

import (
	"flag"
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
	infoPath          fyne.URI = storage.NewFileURI(filepath.Join(".", "data", "info.json"))
	DefaultPlayerPath string   = "player.json"
)

func main() {
	t := flag.Bool("t", false, "Dark theme")
	flag.Parse()

	myApp := app.NewWithID("us.cvbn.button_quest")

	log.SetOutput(NewCustomWriter())

	// //if !fyne.CurrentDevice().IsMobile() {
	// file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.SetOutput(file)
	// //}

	// for i := 1; i < 10000; i += 20 {
	// 	lvl := Find_Level_From_Experience(uint64(i))
	// 	log.Printf("ex %d [%d] %d-%d", i, lvl, Get_Exp_Base_for_Level(lvl), Get_Exp_Base_for_Level(lvl+1))
	// }

	//log.Println("Application Starting")
	// log.Printf("dt %#v", fyne.CurrentDevice())
	// log.Println("Is Mobile", fyne.CurrentDevice().IsMobile())

	prep_init()

	if *t {
		myApp.Settings().SetTheme(theme.DarkTheme())
	} else {
		myApp.Settings().SetTheme(theme.LightTheme())
	}

	log.Println("pref testinPref", myApp.Preferences().BoolWithFallback("TestinPref", false))
	myApp.Preferences().SetBool("TestinPref", true)
	//activePlayer = binding.BindUntyped(&Player{Status: 0, Level: 1, Experience: 5, Name: "NoMan"})
	activePlayerIndex.AddListener(binding.NewDataListener(func() {
		actIndx, _ := activePlayerIndex.Get()
		activePlayer = binding.BindUntyped(&playerList.Players[actIndx])
	}))

	// uri, err := storage.Child(myApp.Storage().RootURI(), "test.txt")
	// infoPath = uri

	//logLifecycle(myApp)

	topWindow = myApp.NewWindow("Button Quest")

	//topWindow.Resize(fyne.NewSize(400, 400))

	topWindow.SetMaster()
	content := container.NewMax()

	content.Objects = []fyne.CanvasObject{welcomeScreen(topWindow, activePlayer)}
	content.Refresh()

	go func() {
		time.Sleep(5 * time.Second)
		content.Objects = []fyne.CanvasObject{mainScreen(topWindow, activePlayer, &playerList)}
		content.Refresh()
	}()

	topWindow.SetMainMenu(GetMainMenu(&myApp, &topWindow, func(name string) {
		switch name {
		case "debug":
			content.Objects = []fyne.CanvasObject{debugScreen(topWindow)}
		case "home":
			content.Objects = []fyne.CanvasObject{mainScreen(topWindow, activePlayer, &playerList)}
		case "save":
			BeforeAppDestroy()
		}
	}))

	// Close the App when Escape key is pressed
	topWindow.Canvas().SetOnTypedKey(func(keyEvent *fyne.KeyEvent) {

		if keyEvent.Name == fyne.KeyEscape {
			//myApp.Quit()
			topWindow.Close()
		}
	})

	//topWindow.SetContent(makeNav(setTutorial, false))
	topWindow.SetContent(content)
	topWindow.Resize(fyne.NewSize(640, 460))

	w, wEs := myApp.Storage().Save("sSave")
	if wEs != nil {
		log.Println("Error", wEs)
		u, _ := myApp.Storage().Create("sSave")

		defer u.Close()
		u.Write([]byte("INITAL CREATE"))

	} else {
		defer w.Close()
		w.Write([]byte("THIS IS A TEST!"))
	}

	topWindow.SetCloseIntercept(func() {
		log.Println("Close Intercept")
		//BeforeAppDestroy()
		// write, err := storage.SaveFileToURI(infoPath)
		// if err != nil {
		// 	fyne.LogError("Unable to save file \"info.json\"", err)
		// }

		// data, err := json.Marshal(playerList)
		// if err != nil {
		// 	fyne.LogError("problem conversting to JSON", err)
		// }
		// _, err = write.Write(data)
		// if err != nil {
		// 	log.Panicln("Failed to write temporary file")
		// }

		// write.Close()
		topWindow.Close()
	})

	myApp.Lifecycle().SetOnStopped(func() {
		log.Println("app stopping")
		BeforeAppDestroy()
		log.Println("app killed")
	})
	log.Println("PRINTING URI []", myApp.Storage().RootURI())

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
	log.Println("prep init")

	activePlayerIndex = binding.NewInt()

	fileLoaded, _ := FirstLoad()

	// v, err := storage.CanRead(infoPath)
	// if v == false || err != nil {
	// 	log.Println("Error E21: ", err)
	// 	return
	// }
	// reader, err := storage.Reader(infoPath)
	// if err != nil {
	// 	log.Println("Error: E22: ", err)
	// 	return
	// }
	// defer reader.Close()

	// data, err := ioutil.ReadAll(reader)
	// if err != nil {
	// 	log.Println("Error: E23: ", err)
	// 	return
	// }
	// log.Println(data)

	// b, br := storage.CanRead(infoPath)
	// log.Printf("canRead %#v %#v", b, br)

	// l, lr := storage.CanList(infoPath)
	// log.Printf("canList %#v %#v", l, lr)

	// s, sr := storage.Reader()

	// // attempt to load from file
	// read, err := ioutil.ReadFile(infoPath.Path())
	// log.Printf("err %#v", err)
	// log.Println("prep")
	// if err == nil {

	// 	err = json.Unmarshal(read, &playerList)
	// 	log.Printf("err2 %#v", err)
	// 	if err == nil {
	// 		fileLoaded = true
	// 	}
	// 	//read.Close()
	// }
	// log.Printf("playerList %#v", playerList)
	if !fileLoaded {
		playerList = PlayerList{Players: []Player{
			{Status: 0, Level: 1, Experience: 5, Name: "NoMan"},
		}}
	}

	if len(playerList.Players) > 0 {
		activePlayerIndex.Set(0)
	}

}
