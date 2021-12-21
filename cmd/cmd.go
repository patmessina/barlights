package cmd

import (
	"barlights/pkg"
	"barlights/server"
	"barlights/types"
	"barlights/utils"
	"errors"
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	version = "v1.0"
)

var (
	port       int
	brightness int
	ledCounts  int
	gpioPin    int
	debug      bool

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
			SetLogging(debug)

			log.WithFields(
				log.Fields{
					"args": args,
					"port": port,
				}).Info("starting barlights server")

			lightOptions := utils.SetLightOptions(brightness, ledCounts, gpioPin)
			server.Start(lightOptions, port)
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
			SetLogging(debug)

			log.WithFields(
				log.Fields{
					"args": args,
				}).Info("setting barlights to solid color")
			var c *types.Color
			var err error

			if v, ok := pkg.DefaultColors[strings.ToLower(args[0])]; ok {
				c, err = pkg.NewColorFromUInt32(v)
				if err != nil {
					return err
				}
			} else if strings.HasPrefix(args[0], "#") ||
				strings.HasPrefix(args[0], "0x") ||
				strings.HasPrefix(args[0], "0X") {
				c, err = pkg.NewColorFromHex(args[0])
				if err != nil {
					return err
				}

			} else {
				return errors.New("please choose a default color or use a hex value")
			}

			lightOptions := utils.SetLightOptions(brightness, ledCounts, gpioPin)
			return pkg.Solid(lightOptions, c)
		},
	}

	offCmd = &cobra.Command{
		Use:   "off",
		Short: "Turn the lights off.",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Infoln("Turning lights off.")
			lightOptions := utils.SetLightOptions(0, ledCounts, gpioPin)
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

	// set
	rootCmd.AddCommand(setCmd)

	// set off
	rootCmd.AddCommand(offCmd)
	setCmd.AddCommand(offCmd)

	// set solid
	setCmd.AddCommand(solidCmd)

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
