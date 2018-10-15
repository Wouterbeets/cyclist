package cyclist

import (
	"encoding/json"
	"os"
)

type keeper struct {
}

func (k *keeper) write(id string, cyc Cycle) error {

	f, err := os.OpenFile("/home/bugless/.screens/."+id, os.O_WRONLY|os.O_CREATE, 0755)
	defer f.Close()
	if err != nil {
		return err
	}
	enc := json.NewEncoder(f)
	enc.SetIndent("", "\t")
	err = enc.Encode(cyc)
	return err
}

func (k *keeper) read(id string) (Cycle, error) {
	var c Cycle

	f, err := os.OpenFile("/home/bugless/.screens/."+id, os.O_RDONLY, 0755)
	defer f.Close()
	if err != nil {
		return c, err
	}

	err = json.NewDecoder(f).Decode(&c)
	return c, err
}
