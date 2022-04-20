package main

import (
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func outsideScreen(_ fyne.Window, player binding.ExternalUntyped) fyne.CanvasObject {
	progress := widget.NewProgressBar()
	mPlr, _ := player.Get()
	plr, _ := mPlr.(Player)
	progress.Min = float64(Base_Level_Experience(&plr))
	progress.Max = float64(Next_Level_At(&plr))

	//infinite := widget.NewProgressBarInfinite()

	playerInfo := NewPlayerInfo(player) // Display_Player_Info(&player)
	// go func() {
	// 	for i := 0.0; i <= 1.0; i += 0.1 {
	// 		time.Sleep(time.Millisecond * 250)
	// 		progress.SetValue(i)
	// 	}
	// }()

	// update the progress bars.
	player.AddListener(binding.NewDataListener(func() {
		iplayer, _ := player.Get()
		plyr, _ := iplayer.(Player)
		if progress.Max < float64(plyr.Experience) {
			progress.Min = float64(Base_Level_Experience(&plyr))
			progress.Max = float64(plyr.NextLevelExp)
		}
		progress.SetValue(float64(plyr.Experience))
		//log.Printf("user %#v %d %d", plyr, plyr.NextLevelExp, plyr.Experience)
		//log.Printf("progress values %#v, %#v, %#v", progress.Min, progress.Max, progress.Value)
	}))

	fightBtn := New_Color_Button("Fight", func() {
		xPlr, _ := player.Get()
		fPlr, _ := xPlr.(Player)
		fPlr.Experience += 20
		lvl := Find_Level_From_Experience(fPlr.Experience)
		log.Printf("fight level check [cur, calc] %d %d", fPlr.Level, lvl)
		if lvl != fPlr.Level {
			fPlr.Level = lvl
			fPlr.NextLevelExp = Get_Exp_Base_for_Level(lvl + 1)
			log.Printf("Updated level [%d] and next level at [%d]", fPlr.Level, fPlr.NextLevelExp)
		}
		// gotta be a better way to do this!?!
		player.Set(fPlr)
		playerInfo.Refresh()
	}, color.NRGBA{R: 200, G: 200, B: 200, A: 255})

	return container.NewVBox(
		playerInfo,
		progress,
		fightBtn,
	)
}
