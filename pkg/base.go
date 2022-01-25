package pkg

import (
	"barlights/types"
	"errors"
	"time"

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
	err = setLights(lights, color, opt.Channels[0].Brightness)
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
	return setLights(lights, c, opt.Channels[0].Brightness)

}

func Cycle(opt *ws2811.Option, settings types.BarlightSettings,
	done chan bool, errCh chan error) {

	lights, err := ws2811.MakeWS2811(opt)
	lights.Init()
	defer lights.Fini()
	if err != nil {
		log.WithFields(
			log.Fields{
				"function":          "Cycle",
				"settingsHexColors": settings.HexColors,
				"brightness":        settings.Brightness,
			},
		).Error(err)
		errCh <- err
		return
	}
	lights.SetBrightness(0, settings.Brightness)

	if len(settings.Colors) < 2 {

		err := errors.New("not enough colors")
		log.WithFields(
			log.Fields{
				"function":          "Cycle",
				"settingsHexColors": settings.HexColors,
				"brightness":        settings.Brightness,
			},
		).Error(err)

		errCh <- err
		return
	}

	err = setLights(lights, settings.Colors[0], settings.Brightness)
	if len(settings.Colors) < 2 {
		log.WithFields(
			log.Fields{
				"function":          "Cycle",
				"settingsHexColors": settings.HexColors,
				"brightness":        settings.Brightness,
			},
		).Error(err)
		errCh <- err
		return
	}

	ledIndex := 0
	colorIndex := 1
	for {
		select {
		case <-done:
			close(done)
			return
		case <-time.After(time.Millisecond * time.Duration(settings.Speed)):
		}

		lights.Leds(0)[ledIndex] = settings.Colors[colorIndex].UInt32
		err := lights.Render()
		if err != nil {

			log.WithFields(
				log.Fields{
					"function":          "Cycle",
					"settingsHexColors": settings.HexColors,
					"brightness":        settings.Brightness,
				},
			).Error(err)

			errCh <- err
			return
		}

		// make sure we stay within bounds
		ledIndex++
		if ledIndex >= len(lights.Leds(0)) {
			ledIndex = 0
			colorIndex++
			if colorIndex >= len(settings.Colors) {
				colorIndex = 0
			}
		}

	}

}
