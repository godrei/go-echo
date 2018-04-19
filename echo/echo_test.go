package echo

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLog(t *testing.T) {
	{
		var b bytes.Buffer
		out = &b

		Log("Hello World!", "no_format")

		require.Equal(t, "Hello World!\n", b.String())
	}

	{
		var b bytes.Buffer
		out = &b

		Log("Hello World!", "upper")

		require.Equal(t, "HELLO WORLD!\n", b.String())
	}

	{
		var b bytes.Buffer
		out = &b

		Log("Hello World!", "lower")

		require.Equal(t, "hello world!\n", b.String())
	}
}
