//go:build darwin || linux || windows
// +build darwin linux windows

package main

import (
	"golang.org/x/mobile/app"
	// "golang.org/x/mobile/exp/app/debug"
)

func main() {
	app.Main(func(a app.App) {
		//debug.NewFPS()
	})
}

// func draw() {
// 	gl.ClearColor(1, 0, 0, 1)
// 	gl.Clear(gl.COLOR_BUFFER_BIT)

// }
