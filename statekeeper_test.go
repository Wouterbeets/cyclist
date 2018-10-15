package cyclist

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWrite(t *testing.T) {
	k := keeper{}
	err := k.write("foo", Cycle{Entries: []Entry{{"foo"}}})
	require.NoError(t, err)
}

func TestRead(t *testing.T) {
	k := keeper{}
	c1 := Cycle{Entries: []Entry{{"foo"}}}
	err := k.write("foo", c1)
	require.NoError(t, err)
	c2, err := k.read("foo")
	require.NoError(t, err)
	require.Equal(t, c2, c1)
}
