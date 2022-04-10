package main

import (
	"math"
)

var exponent float64 = 3
var baseXP int = 100

// find and calculate the 'next' level at which the player will level up
func Next_Level_At(player *Player) uint64 {
	// right now we'll just use sets of 100 for levels...
	//return Base_Level_Experience(player) + 100
	lvl := Find_Level_From_Experience(player.Experience)
	return Get_Exp_Base_for_Level(lvl + 1)
}

func Get_Exp_Base_for_Level(level int) uint64 {
	return uint64(math.Floor(float64(baseXP) * math.Pow(float64(level-1), exponent)))
}

// find the calucated experince that this user reached the active level at
func Base_Level_Experience(player *Player) uint64 {
	//math.Ceil(math.Pow(float64(player.Experience / uint64(baseXP)), 1 / exponent))
	// return int64(math.Floor(float64(player.Experience)/100) * 100)
	// get calculated level first
	lvl := Find_Level_From_Experience(player.Experience)
	return Get_Exp_Base_for_Level(lvl)
}

func Find_Level_From_Experience(exp uint64) int {
	// step1 := float64(exp) / float64(baseXP)
	// step2 := 1 / exponent
	// step3 := math.Pow(float64(step1), step2)
	// step4 := math.Ceil(step3)
	// step5 := int(step4) + 1
	// log.Printf("one %f, %f, %f, %f, %d", step1, step2, step3, step4, step5)
	return int(math.Ceil(math.Pow(float64(float64(exp)/float64(baseXP)), 1/exponent)))
}
