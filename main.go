package main

import (
	"os"

	"github.com/godrei/go-echo/echo"
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "go-echo <message>",
	Short: "Prints a given message.",
	Run: func(cmd *cobra.Command, args []string) {
		message := args[0]
		echo.Log(message)
	},
}

func init() {
	cmd.Args = cobra.MinimumNArgs(1)
}

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
