package pkg

import (
	ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"
)

// Solid Color -- set all leds to the same value
func Solid(opt ws2811.Option, color string) error {

	var c uint32
	var err error

	if hex, ok := DefaultSolidColors[color]; ok {
		c = hex
	} else {
		c, err = HexStrToUInt32(color)
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
