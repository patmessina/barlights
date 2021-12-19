package cmd

import (
	"barlights/pkg"
	"fmt"
	"os"

	ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"
	log "github.com/sirupsen/logrus"
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
	debug        bool
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

	colorCmd = &cobra.Command{
		Use:   "color [color]",
		Short: "testing color",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			SetLogging(debug)
			pkg.NewColorFromRGB(93, 138, 168)
			pkg.NewColorFromUInt32(6130344)
			_, err := pkg.NewColorFromHex("0X5d8aa8")
			return err
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

	cobra.OnInitialize()
	rootCmd.AddCommand(versionCmd)
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d",
		false, "debug mode")
	rootCmd.PersistentFlags().IntVarP(&brightness, "brightness", "b",
		60, "LED brightness")
	rootCmd.PersistentFlags().IntVarP(&ledCounts, "ledcount", "l",
		145, "number of LEDs")
	rootCmd.PersistentFlags().IntVarP(&gpioPin, "gpio-pin", "g",
		18, "pin on the raspberry pi where the signal will be available")

	// server
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().IntVarP(&port, "port", "p", 8080, "desired port number")

	rootCmd.AddCommand(colorCmd)

	// off
	rootCmd.AddCommand(offCmd)

	// set
	rootCmd.AddCommand(setCmd)

	// set solid
	setCmd.AddCommand(solidCmd)

	if debug {
		fmt.Println("here")
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	// set led options
	lightOptions = ws2811.DefaultOptions
	lightOptions.Channels[0].Brightness = brightness
	lightOptions.Channels[0].LedCount = ledCounts
	lightOptions.Channels[0].GpioPin = gpioPin

}

// SetLogging set logging level
func SetLogging(debug bool) {
	if debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
