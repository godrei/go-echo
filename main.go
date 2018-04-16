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
		format := os.Getenv("format")
		if format == "" {
			if uppercase {
				format = "upper"
			} else if lowercase {
				format = "lower"
			}
		}
		echo.Log(message, format)
	},
}

var uppercase bool
var lowercase bool

func init() {
	cmd.PersistentFlags().BoolVarP(&uppercase, "upper", "u", false, "Converts message to uppercase characters.")
	cmd.PersistentFlags().BoolVarP(&lowercase, "lower", "l", false, "Converts message to lowercase characters.")
}

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
