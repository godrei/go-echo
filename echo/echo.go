package echo

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

// Log logs message (msg)
func Log(msg string) {
	fmt.Fprintln(out, msg)
}
