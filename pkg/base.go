package pkg

import (
	"barlights/utils"

	ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"
)

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

	dev, err := ws2811.MakeWS2811(&opt)

	dev.Init()
	if err != nil {
		return err
	}

	defer dev.Fini()
	if err != nil {
		return err
	}

	for i := 0; i < opt.Channels[0].LedCount; i++ {
		dev.Leds(0)[i] = c
	}

	err = dev.Render()
	if err != nil {
		return err
	}

	return nil
}

func Off(opt ws2811.Option) error {
	dev, err := ws2811.MakeWS2811(&opt)

	dev.Init()
	if err != nil {
		return err
	}

	defer dev.Fini()
	if err != nil {
		return err
	}

	for i := 0; i < opt.Channels[0].LedCount; i++ {
		dev.Leds(0)[i] = DefaultSolidColors["off"]
	}

	err = dev.Render()
	if err != nil {
		return err
	}

	return nil

}
