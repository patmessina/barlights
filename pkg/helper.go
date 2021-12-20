package pkg

import (
	"barlights/types"
	"time"

	ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"
)

// SetLights -- Set lights to a specified color
func setLights(lights *ws2811.WS2811, color *types.Color) error {
	for i := 0; i < len(lights.Leds(0)); i++ {
		lights.Leds(0)[i] = color.UInt32
	}
	err := lights.Render()
	if err != nil {
		return err
	}
	return nil
}

// cycleSet move down the strip to the right changing the color behind to the
// given color
func cycleSet(lights *ws2811.WS2811, color uint32, speed int) error {
	for i := 0; i < len(lights.Leds(0)); i++ {
		lights.Leds(0)[i] = color
		err := lights.Render()
		if err != nil {
			return err
		}
		time.Sleep(time.Duration(speed) * time.Millisecond)

	}
	return nil
}

// moveBallDown moves a ball down the light strip. Offset will determine how
// far off strip the ball will go
func moveBallDown(lights *ws2811.WS2811, stripColor uint32,
	ball *types.Ball, offset int, startToEnd bool) error {

	var err error
	if startToEnd {
		err = moveBallStartToEnd(lights, stripColor, ball, offset)
	} else {
		err = moveBallEndToStart(lights, stripColor, ball, offset)
	}

	return err
}

// moveBallStartToEnd move the ball from the start to the end with a given offset
func moveBallStartToEnd(lights *ws2811.WS2811, stripColor uint32,
	ball *types.Ball, offset int) error {

	for i := 0; i < len(lights.Leds(0))+offset; i++ {
		for j := 0; j < ball.Size; j++ {
			if i-j >= 0 && i-j < len(lights.Leds(0)) {
				lights.Leds(0)[i-j] = ball.Color
			}
		}
		if i-ball.Size >= 0 && i-ball.Size < len(lights.Leds(0)) {
			lights.Leds(0)[i-ball.Size] = stripColor
		}
		err := lights.Render()
		if err != nil {
			return err
		}
		time.Sleep(time.Duration(ball.Speed) * time.Millisecond)

	}

	return nil
}

// moveBallEndToStart move the ball from the end to the start with a given offset
func moveBallEndToStart(lights *ws2811.WS2811, stripColor uint32,
	ball *types.Ball, offset int) error {

	for i := len(lights.Leds(0)); i > -offset; i-- {

		for j := 0; j < ball.Size; j++ {
			if i+j >= 0 && i+j < len(lights.Leds(0)) {
				lights.Leds(0)[i+j] = ball.Color
			}
		}

		if i+ball.Size >= 0 && i+ball.Size < len(lights.Leds(0)) {
			lights.Leds(0)[i+ball.Size] = stripColor
		}

		err := lights.Render()
		if err != nil {
			return err
		}
		time.Sleep(time.Duration(ball.Speed) * time.Millisecond)

	}

	return nil
}
