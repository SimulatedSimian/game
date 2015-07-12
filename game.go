package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"github.com/simulatedsimian/rect"
)

func main() {

	fmt.Println("Caps: ", GetJoyCaps(0))

	for {
		fmt.Println("Read: ", GetJoyPosEx(0))
	}

	OpenJoystick(0)
}

func mainx() {

	OpenJoystick(0)

	/*
		js, jerr := OpenJoystick("/dev/input/js0")

		if jerr == nil {
			fmt.Println("Axis Count: ", js.AxisCount)
			fmt.Println("Button Count: ", js.ButtonCount)
			fmt.Println("Name: ", js.Name)

			for {
				ev, err := js.GetEvent()
				if err != nil {
					panic(err)
				}

				fmt.Printf("%v\n", &ev)
			}
		}
	*/

	gc := GameCore{}

	var r rect.Rectangle
	x, y := 0, 0
	dx, dy := 1, 1

	b := MakeMemBuffer(10, 10)
	Fill(b, '0', termbox.ColorRed, termbox.ColorBlue, ALL)

	gc.OnInit = func(gc *GameCore) error {
		r = rect.WH(gc.BackBuffer.Size())
		r.Expand(rect.Vec{-1, -1})
		return nil
	}

	gc.OnEvent = func(gc *GameCore, ev *termbox.Event) error {
		if ev.Type == termbox.EventKey {
			if ev.Ch == 'q' {
				gc.DoQuit = true
			}
		}
		return nil
	}

	gc.OnTick = func(gc *GameCore) error {
		FillArea(gc.BackBuffer, r, 'x', termbox.ColorCyan, termbox.ColorGreen, ALL)
		BlitBuffer(b, gc.BackBuffer, x, y)
		x += dx
		y += dy

		if x > 60 || x <= 0 {
			dx = -dx
		}
		if y > 30 || y <= 0 {
			dy = -dy
		}
		return nil
	}

	err := gc.Run()
	if err != nil {
		panic(err)
	}
}
