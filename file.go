package cascade

import (
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v3"
)

func loadFile(path string, cfg any) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	switch filepath.Ext(path) {
	case ".yaml", ".yml":
		return yaml.Unmarshal(data, cfg)
	case ".toml":
		return toml.Unmarshal(data, cfg)
	default:
		return nil
	}
}
