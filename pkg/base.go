package pkg

import (
	"barlights/types"

	ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"
	log "github.com/sirupsen/logrus"
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

	log.WithFields(
		log.Fields{
			"function":   "Solid",
			"color":      color.Hex,
			"brightness": opt.Channels[0].Brightness,
		},
	).Debug("setting lights")
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

	log.WithFields(
		log.Fields{
			"function":   "Solid",
			"color":      c.Hex,
			"brightness": opt.Channels[0].Brightness,
		},
	).Debug("setting lights off")
	return setLights(lights, c)

}
