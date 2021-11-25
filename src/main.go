package main

import (
	"golang.org/x/mobile/app"
	// "golang.org/x/mobile/app/debug"
	"golang.org/x/mobile/gl"
)

func main() {
	app.Run(app.Callbacks{draw: draw})
}

func draw() {
	gl.ClearColor(1, 0, 0, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT)

	// debug.DrawFPS()
}
