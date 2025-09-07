package drivers

import (
	"strings"

	"github.com/BurntSushi/toml"
)

type TOMLDriver struct{}

func (TOMLDriver) CanHandle(path string) bool {
	return strings.HasSuffix(path, ".toml")
}
func (TOMLDriver) Unmarshal(data []byte, cfg any) error {
	return toml.Unmarshal(data, cfg)
}
