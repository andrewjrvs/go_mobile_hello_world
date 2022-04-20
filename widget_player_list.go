package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type HeroList struct {
	widget.BaseWidget
	playerList  binding.Untyped
	activeIndex binding.Int
}

func NewHeroList(playerList binding.Untyped, activeIndex binding.Int) *HeroList {
	heroList := &HeroList{
		playerList:  playerList,
		activeIndex: activeIndex,
	}

	heroList.ExtendBaseWidget(heroList)
	return heroList
}

func (c *HeroList) Refresh() {

}

// CreateRenderer is a private method to Fyne which links this widget to its renderer
func (c *HeroList) CreateRenderer() fyne.WidgetRenderer {

	content := container.NewMax()
	c.playerList.AddListener(binding.NewDataListener(func() {

		ipl, _ := c.playerList.Get()
		pl, _ := ipl.(PlayerList)

		pplo := []fyne.CanvasObject{}

		for _, plyr := range pl.Players {
			pplo = append(pplo, container.New(layout.NewCenterLayout(), widget.NewLabel(strconv.Itoa(plyr.Level))))
		}

		rtnContainer := container.New(
			layout.NewGridLayoutWithColumns(5),
			pplo...,
		)

		content.Objects = []fyne.CanvasObject{rtnContainer}
		content.Refresh()
	}))

	objects := []fyne.CanvasObject{content}
	r := &HeroListRenderer{objects, c}
	r.applyTheme()
	return r
}

type HeroListRenderer struct {
	objects []fyne.CanvasObject

	heroList *HeroList
}

func (c *HeroListRenderer) Destroy() {
}

func (c *HeroListRenderer) Objects() []fyne.CanvasObject {

	return c.objects
}

// Layout the components of the card container.
func (c *HeroListRenderer) Layout(size fyne.Size) {

}

// MinSize calculates the minimum size of a card.
// This is based on the contained text, image and content.
func (c *HeroListRenderer) MinSize() fyne.Size {

	min := fyne.NewSize(theme.Padding(), theme.Padding())

	titlePad := theme.Padding() * 2
	min = min.Add(fyne.NewSize(0, titlePad*2))

	min = fyne.NewSize(fyne.Max(min.Width, titlePad*2+theme.Padding()),
		min.Height)

	lineHeight := float32(0)
	for _, s := range c.objects {
		lineHeight += s.Size().Height
	}
	min = min.Add((fyne.NewSize(0, lineHeight)))

	return min
}

func (c *HeroListRenderer) Refresh() {

}

// applyTheme updates this button to match the current theme
func (c *HeroListRenderer) applyTheme() {

}
