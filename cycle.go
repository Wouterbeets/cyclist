// cylcist creates and manages lists with and allows you to cylce through them
package cyclist

import (
	"github.com/pkg/errors"
)

type stateKeeper interface {
	write(id string, cycle Cycle) error
	read(id string) (Cycle, error)
}

type Cyclist struct {
	stateKeeper
}

func New(cyc map[string]Cycle) (Cyclist, error) {
	cyclist := Cyclist{&keeper{}}
	for id, c := range cyc {
		err := cyclist.write(id, c)
		if err != nil {
			return cyclist, err
		}
	}
	return cyclist, nil
}

// Cylce reads from the global state and returns the cycle entry
func (c Cyclist) Cycle(id string) (string, error) {
	cyc, err := c.read(id)
	if err != nil {
		return "", errors.Wrapf(err, "unable to read cycle for id %s", id)
	}
	cmd := cyc.Cycle()
	err = c.write(id, cyc)
	return cmd, err
}

// Cycle holds a list of Entries
type Cycle struct {
	Entries []Entry
	Current int
}

func (c *Cycle) Cycle() string {
	if len(c.Entries) == 0 {
		return ""
	}
	cmd := c.Entries[c.Current]
	c.Current++
	if len(c.Entries) == c.Current {
		c.Current = 0
	}
	return cmd.Cmd
}

type Entry struct {
	Cmd string
}
