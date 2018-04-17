package echo

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLog(t *testing.T) {
	var b bytes.Buffer
	out = &b

	Log("Hello World!")

	require.Equal(t, "Hello World!\n", b.String())
}
