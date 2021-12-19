package pkg

import (
	"barlights/types"
	"barlights/utils"
	"fmt"
	"math/rand"
	"time"

	ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"
)

// Pong ball goes back and forth
func Pong(opt ws2811.Option,
	maxTime int64,
	ball *types.Ball, stripColors []uint32) error {

	rand.Seed(time.Now().UnixNano())
	lights, err := ws2811.MakeWS2811(&opt)

	lights.Init()
	defer lights.Fini()
	if err != nil {
		return err
	}

	setLights(lights, utils.RandColor(stripColors))

	// Should probably stop sometime
	start := time.Now().Unix()
	for time.Now().Unix()-start < maxTime {

		// bounce back and forth
		color := utils.RandColor(stripColors)
		err = moveBallDown(lights, color, ball, ball.Size, true)
		if err != nil {
			return err
		}

		color = utils.RandColor(stripColors)
		err = moveBallDown(lights, color, ball, ball.Size, false)
		if err != nil {
			return err
		}

	}

	return nil
}

// Siren -- flash through a range of colors, when the selection is made up of
// only two values it behaves like siren
func Siren(opt ws2811.Option,
	speed int, maxTime int64,
	selection []string) error {

	var cycle string
	if len(selection) > 0 {
		cycle = selection[0]
	} else {
		cycle = "fire"
	}
	fmt.Println(cycle)

	lights, err := ws2811.MakeWS2811(&opt)

	lights.Init()
	defer lights.Fini()
	if err != nil {
		return err
	}

	start := time.Now().Unix()
	for time.Now().Unix()-start < maxTime {

		for _, color := range CycleColors[cycle] {

			err = setLights(lights, color)
			if err != nil {
				return err
			}

			time.Sleep(time.Duration(speed) * time.Millisecond)

		}

	}

	return nil
}

// Cycle -- will cycle through a range of colors by invidually changing each
// led
func Cycle(opt ws2811.Option, speed int, selection []string) error {

	var cycle string
	if len(selection) > 0 {
		cycle = selection[0]
	} else {
		cycle = "default"
	}

	lights, err := ws2811.MakeWS2811(&opt)

	lights.Init()
	defer lights.Fini()
	if err != nil {
		return err
	}

	// TODO: Check if the cycle exists before running
	setLights(lights, DefaultSolidColors["blue"])
	for _, color := range CycleColors[cycle] {
		// for i := 0; i < opt.Channels[0].LedCount; i++ {
		cycleSet(lights, color, speed)
		// }
	}

	return nil
}

// Solid Color -- set all leds to the same value
func Solid(opt ws2811.Option, color string) error {

	var c uint32
	var err error

	if hex, ok := DefaultSolidColors[color]; ok {
		c = hex
	} else {
		c, err = utils.HexStrToInt(color)
		if err != nil {
			return err
		}

	}

	lights, err := ws2811.MakeWS2811(&opt)

	lights.Init()
	defer lights.Fini()
	if err != nil {
		return err
	}

	for i := 0; i < opt.Channels[0].LedCount; i++ {
		lights.Leds(0)[i] = c
	}

	err = lights.Render()
	if err != nil {
		return err
	}

	return nil
}

// Off -- turn off all the leds
func Off(opt ws2811.Option) error {
	lights, err := ws2811.MakeWS2811(&opt)

	lights.Init()
	defer lights.Fini()
	if err != nil {
		return err
	}

	return setLights(lights, DefaultSolidColors["off"])

}
