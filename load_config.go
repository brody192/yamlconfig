package yamlconfig

import (
	"os"

	"gopkg.in/validator.v2"
	"gopkg.in/yaml.v3"
)

func loadConfig(path string, output any) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	defer f.Close()

	if err := yaml.NewDecoder(f).Decode(output); err != nil {
		return err
	}

	if err := validator.Validate(output); err != nil {
		return err
	}

	return nil
}
