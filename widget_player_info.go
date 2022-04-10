package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type PlayerInfo struct {
	widget.BaseWidget
	player binding.Untyped
}

func NewPlayerInfo(player binding.Untyped) *PlayerInfo {
	playerInfo := &PlayerInfo{
		player: player,
	}

	playerInfo.ExtendBaseWidget(playerInfo)
	return playerInfo
}

func (c *PlayerInfo) SetPlayer(player Player) {
	c.player.Set(player)
}

func (c *PlayerInfo) Refresh() {

}

// CreateRenderer is a private method to Fyne which links this widget to its renderer
func (c *PlayerInfo) CreateRenderer() fyne.WidgetRenderer {
	dtName := binding.NewString()
	dtExp := binding.NewString()
	dtLevel := binding.NewString()

	c.ExtendBaseWidget(c)

	c.player.AddListener(binding.NewDataListener(func() {
		iplr, _ := c.player.Get()
		plr, _ := iplr.(Player)
		dtName.Set(plr.Name)
		dtExp.Set(strconv.FormatUint(plr.Experience, 10))
		dtLevel.Set(strconv.Itoa(plr.Level))
	}))

	header := canvas.NewText("Player Details", theme.ForegroundColor())
	header.TextStyle.Bold = true

	lblName := widget.NewLabel("Player:")
	txtName := widget.NewLabelWithData(dtName)
	lblLevel := widget.NewLabel("Level:")
	txtLevel := widget.NewLabelWithData(dtLevel)
	lblExp := widget.NewLabel("Experience:")
	txtExp := widget.NewLabelWithData(dtExp)

	rtnContainer := container.New(
		layout.NewGridLayoutWithRows(3),
		container.New(
			layout.NewGridLayoutWithColumns(2),
			lblName, txtName,
		),
		container.New(
			layout.NewGridLayoutWithColumns(2),
			lblLevel, txtLevel,
		),
		container.New(
			layout.NewGridLayoutWithColumns(2),
			lblExp, txtExp,
		),
	)

	rtnContainer.Move(fyne.Position{X: 0, Y: (theme.Padding() * 2) + header.MinSize().Height})

	objects := []fyne.CanvasObject{header, rtnContainer}

	r := &PlayerInfoRenderer{header, objects, c}
	r.applyTheme()
	return r
}

type PlayerInfoRenderer struct {
	header  *canvas.Text
	objects []fyne.CanvasObject

	playerInfo *PlayerInfo
}

func (c *PlayerInfoRenderer) Destroy() {
}

func (c *PlayerInfoRenderer) Objects() []fyne.CanvasObject {

	return c.objects
}

// Layout the components of the card container.
func (c *PlayerInfoRenderer) Layout(size fyne.Size) {

}

// MinSize calculates the minimum size of a card.
// This is based on the contained text, image and content.
func (c *PlayerInfoRenderer) MinSize() fyne.Size {

	min := fyne.NewSize(theme.Padding(), theme.Padding())

	titlePad := theme.Padding() * 2
	min = min.Add(fyne.NewSize(0, titlePad*2))

	headerMin := c.header.MinSize()
	min = fyne.NewSize(fyne.Max(min.Width, headerMin.Width+titlePad*2+theme.Padding()),
		min.Height+headerMin.Height)

	lineHeight := float32(0)
	for _, s := range c.objects {
		lineHeight += s.Size().Height
	}
	min = min.Add((fyne.NewSize(0, lineHeight)))

	return min
}

func (c *PlayerInfoRenderer) Refresh() {

}

// applyTheme updates this button to match the current theme
func (c *PlayerInfoRenderer) applyTheme() {
	if c.header != nil {
		c.header.TextSize = theme.TextHeadingSize()
		c.header.Color = theme.ForegroundColor()
	}
}
