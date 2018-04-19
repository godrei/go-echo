package main

import (
	"os"

	"github.com/bitrise-io/go-utils/log"
	"github.com/godrei/go-echo/echo"
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "go-echo <message>",
	Short: "Prints a given message.",
	Run: func(cmd *cobra.Command, args []string) {
		message := os.Getenv("message")
		if message == "" && len(args) > 0 {
			message = args[0]
		}
		if message == "" {
			log.Errorf("message to print not specified")
			os.Exit(1)
		}
		echo.Log(message)
	},
}

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
