package pkg

import (
	"barlights/utils"
	"fmt"
	"time"

	ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"
)

// Siren
func Siren(opt ws2811.Option, speed int, maxTime int64, selection []string) error {

	var cycle string
	if len(selection) > 0 {
		cycle = selection[0]
	} else {
		cycle = "fire"
	}
	fmt.Println(cycle)

	dev, err := ws2811.MakeWS2811(&opt)

	dev.Init()
	if err != nil {
		return err
	}

	defer dev.Fini()
	if err != nil {
		return err
	}

	start := time.Now().Unix()
	for start-time.Now().Unix() < maxTime {

		for _, color := range CycleColors[cycle] {

			for i := 0; i < opt.Channels[0].LedCount; i++ {
				dev.Leds(0)[i] = color
			}

			err = dev.Render()
			if err != nil {
				return err
			}
			time.Sleep(time.Duration(speed) * time.Millisecond)

		}

	}

	return nil
}

func Cycle(opt ws2811.Option, speed int, selection []string) error {

	var cycle string
	if len(selection) > 0 {
		cycle = selection[0]
	} else {
		cycle = "default"
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

	// TODO: Check if the cycle exists before running
	for _, color := range CycleColors[cycle] {
		for i := 0; i < opt.Channels[0].LedCount; i++ {
			dev.Leds(0)[i] = color
			err = dev.Render()
			if err != nil {
				return err
			}
			time.Sleep(time.Duration(speed) * time.Millisecond)

		}
	}

	return nil
}

// Solid Color
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

// Off
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
