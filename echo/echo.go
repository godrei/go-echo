package echo

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var out io.Writer = os.Stdout

// Log logs message (msg)
func Log(msg, format string) {
	if format == "upper" {
		msg = strings.ToUpper(msg)
	} else if format == "lower" {
		msg = strings.ToLower(msg)
	}
	fmt.Fprintln(out, msg)
}
