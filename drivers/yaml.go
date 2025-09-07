package drivers

import (
	"strings"

	"gopkg.in/yaml.v3"
)

type YAMLDriver struct{}

func (YAMLDriver) CanHandle(path string) bool {
	return strings.HasSuffix(path, ".yaml") || strings.HasSuffix(path, ".yml")
}
func (YAMLDriver) Unmarshal(data []byte, cfg any) error {
	return yaml.Unmarshal(data, cfg)
}
