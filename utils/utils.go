package utils

import ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"

// SetLedOptions
func SetLightOptions(brightness, ledCounts, gpioPin int) *ws2811.Option {
	// set led options
	lightOptions := ws2811.DefaultOptions
	lightOptions.Channels[0].Brightness = brightness
	lightOptions.Channels[0].LedCount = ledCounts
	lightOptions.Channels[0].GpioPin = gpioPin

	return &lightOptions
}
