package pkg

import (
	"barlights/types"

	ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"
)

// Solid Color -- set all leds to the same value
func Solid(opt *ws2811.Option, color *types.Color) error {

	var err error

	lights, err := ws2811.MakeWS2811(opt)
	lights.Init()
	defer lights.Fini()
	if err != nil {
		return err
	}

	err = setLights(lights, color)
	if err != nil {
		return err
	}

	return nil
}

// Off -- turn off all the leds
func Off(opt *ws2811.Option) error {
	lights, err := ws2811.MakeWS2811(opt)

	lights.Init()
	defer lights.Fini()
	if err != nil {
		return err
	}

	c, err := NewColorFromUInt32(OFF)
	if err != nil {
		return err
	}

	return setLights(lights, c)

}
