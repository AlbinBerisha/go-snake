package graphics

import (
	"fmt"
	"os"
	"os/exec"
)

type Window struct {
	Width  int
	Height int
	Buffer []byte
}

func (window *Window) Clear() {
	command := exec.Command("clear")
	command.Stdout = os.Stdout
	command.Run()
	for y := 0; y < window.Height; y++ {
		for x := 0; x <= window.Width; x++ {
			if x == window.Width {
				window.Buffer[y*(window.Width+1)+x] = '\n'
			} else {
				window.Buffer[y*(window.Width+1)+x] = '-'
			}
		}
	}
}

func (window *Window) Display() {
	fmt.Printf("%s", window.Buffer)
}

func (window *Window) Draw(object Object) {
	for y := object.Y; y < object.Y+object.Height; y++ {
		for x := object.X; x < object.X+object.Width; x++ {
			if x >= 0 && x < window.Width && y >= 0 && y < window.Height {
				window.Buffer[y*(window.Width+1)+x] = object.Buffer[(y-object.Y)*object.Width+(x-object.X)]
			}
		}
	}
}

type Object struct {
	X      int
	Y      int
	Width  int
	Height int
	Buffer []byte
}
