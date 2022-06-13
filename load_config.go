package yamlconfig

import (
	"fmt"
	"os"

	"gopkg.in/validator.v2"
	"gopkg.in/yaml.v3"
)

func LoadConfig(path string, output any) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	defer f.Close()

	if err := yaml.NewDecoder(f).Decode(output); err != nil {
		return fmt.Errorf("error decoding config: %s", err)
	}

	if err := validator.Validate(output); err != nil {
		return fmt.Errorf("config validation failed: %s", err)
	}

	return nil
}
