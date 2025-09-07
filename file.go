package cascade

import (
	"fmt"
	"os"

	Drivers "github.com/Unfield/cascade/drivers"
)

func (l *Loader) loadFile(cfg interface{}) error {
	if l.file == "" {
		return nil
	}

	data, err := os.ReadFile(l.file)
	if err != nil {
		return err
	}

	drivers := l.drivers
	if len(drivers) == 0 {
		drivers = []FileDriver{Drivers.YAMLDriver{}, Drivers.TOMLDriver{}}
	}

	for _, d := range drivers {
		if d.CanHandle(l.file) {
			return d.Unmarshal(data, cfg)
		}
	}

	return fmt.Errorf("no driver found for file: %s", l.file)
}
