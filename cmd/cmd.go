package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	version = "v1.0"
)

var (
	port int

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

	startLightsCmd = &cobra.Command{
		Use:   "start",
		Short: "Start the lights.",
		Long:  `Turn the lights on.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("starting barlight server on port %v\n", port)
			// TODO: start server
			return nil
		},
	}

	offCmd = &cobra.Command{
		Use:   "off",
		Short: "Turn the lights off.",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("Turning lights off.\n")
			// TODO: Turn off the lights
			return nil
		},
	}
)

func init() {

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(serverCmd)

	serverCmd.Flags().IntVarP(&port, "port", "p", 8080, "desired port number")

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
