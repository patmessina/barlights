package cmd

import (
	"barlights/pkg"
	"fmt"
	"os"

	ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"
	"github.com/spf13/cobra"
)

const (
	version = "v1.0"
)

var (
	port         int
	brightness   int
	ledCounts    int
	gpioPin      int
	lightOptions ws2811.Option

	rootCmd = &cobra.Command{

		Use:   "barlights",
		Short: "Barlights are lights on the bar!",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Do Stuff Here

			// err := server.Run(lights)

			return nil
		},
	}

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of Barlights",
		Long:  `All software has versions. This is Barlight's`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Barlights %v\n", version)
		},
	}

	serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Start barlight server",
		Long:  `Starts server that hosts a API with the light strip.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("starting Barlights server on port %v\n", port)
			// TODO: start server
			return nil
		},
	}

	setCmd = &cobra.Command{
		Use:   "set",
		Short: "Start the lights.",
		Long:  `Turn the lights on.`,
	}

	solidCmd = &cobra.Command{
		Use:   "solid [color]",
		Short: "Change lights to a solid color.",
		Args:  cobra.ExactArgs(1),
		Long:  `Turn the lights on.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("setting barlights to %v\n", args[0])
			return pkg.Solid(lightOptions, args[0])
		},
	}

	offCmd = &cobra.Command{
		Use:   "off",
		Short: "Turn the lights off.",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("Turning lights off.\n")
			return pkg.Off(lightOptions)
		},
	}
)

func init() {

	rootCmd.AddCommand(versionCmd)
	rootCmd.PersistentFlags().IntVarP(&brightness, "brightness", "b",
		60, "LED brightness")
	rootCmd.PersistentFlags().IntVarP(&ledCounts, "ledcount", "l",
		145, "number of LEDs")
	rootCmd.PersistentFlags().IntVarP(&gpioPin, "gpio-pin", "p",
		18, "pin on the raspberry pi where the signal will be available")

	// server
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().IntVarP(&port, "port", "p", 8080, "desired port number")

	// off
	rootCmd.AddCommand(offCmd)

	// set
	rootCmd.AddCommand(setCmd)

	// set solid
	setCmd.AddCommand(solidCmd)

	// set led options
	lightOptions = ws2811.DefaultOptions
	lightOptions.Channels[0].Brightness = brightness
	lightOptions.Channels[0].LedCount = ledCounts
	lightOptions.Channels[0].GpioPin = gpioPin

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
