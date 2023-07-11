package main

import (
	"time"

	"github.com/AlbinBerisha/go-snake/pkg/graphics"
)

func main() {
	const WindowWidth int = 50
	const WindowHeight int = 25
	window := graphics.Window{Width: WindowWidth, Height: WindowHeight, Buffer: make([]byte, WindowWidth*WindowHeight+WindowHeight)}
	object := graphics.Object{X: 0, Y: 0, Width: 5, Height: 2, Buffer: []byte("D V DVIDEO")}
	xSpeed := 1
	ySpeed := 1
	for {
		window.Clear()
		window.Draw(object)
		window.Display()
		object.X += xSpeed
		object.Y += ySpeed
		if object.X+object.Width >= window.Width || object.X < 0 {
			xSpeed = -xSpeed
			object.X += xSpeed + 1
		}
		if object.Y+object.Height >= window.Height || object.Y < 0 {
			ySpeed = -ySpeed
			object.Y += ySpeed + 1
		}
		time.Sleep(100 * time.Millisecond)
	}
}
