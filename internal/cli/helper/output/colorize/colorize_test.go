package colorize

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRed(t *testing.T) {
	t.Parallel()

	t.Run("", func(t *testing.T) {
		s := "foo"
		got := Red("foo")
		want := fmt.Sprintf("%s%s%s", red, s, reset)

		require.Equal(t, want, got)
	})
}

func TestGreen(t *testing.T) {
	t.Parallel()

	t.Run("", func(t *testing.T) {
		s := "foo"
		got := Green(s)
		want := fmt.Sprintf("%s%s%s", green, s, reset)

		require.Equal(t, want, got)
	})
}

func TestYellow(t *testing.T) {
	t.Parallel()

	t.Run("", func(t *testing.T) {
		s := "foo"
		got := Yellow(s)
		want := fmt.Sprintf("%s%s%s", yellow, s, reset)

		require.Equal(t, want, got)
	})
}

func TestCyan(t *testing.T) {
	t.Parallel()

	t.Run("", func(t *testing.T) {
		s := "foo"
		got := Cyan(s)
		want := fmt.Sprintf("%s%s%s", cyan, s, reset)

		require.Equal(t, want, got)
	})
}
