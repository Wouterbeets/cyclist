package cyclist

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCycle(t *testing.T) {
	c := Cycle{
		Entries: []string{
			"cmd1",
			"cmd2",
			"cmd3",
		},
	}

	require.Equal(t, "cmd1", c.Cycle())
	require.Equal(t, "cmd2", c.Cycle())
	require.Equal(t, "cmd3", c.Cycle())
	require.Equal(t, "cmd1", c.Cycle())
	require.Equal(t, "cmd2", c.Cycle())
	require.Equal(t, "cmd3", c.Cycle())
}
