package config

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

func SaveFile(filename string, data interface{}) error {
	fileExt, err := getValidExtension(filename)
	if err != nil {
		return err
	}

	switch fileExt {
	case ".yaml", ".yml":
		err = saveAsYAML(filename, data)
	case ".json":
		err = saveAsJSON(filename, data, true)
	default:
		return fmt.Errorf("SaveFileError: Unknown ERROR saving %q", filename)
	}
	return err
}

func saveAsYAML(filename string, data interface{}) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	enc := yaml.NewEncoder(file)
	enc.Encode(data)
	enc.Close()
	return nil
}

func saveAsJSON(filename string, data interface{}, ident bool) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	enc := json.NewEncoder(file)
	if ident {
		enc.SetIndent("", "  ")
	}
	enc.Encode(data)
	return nil
}

func AsJSON(data interface{}, ident bool) ([]byte, error) {
	if ident {
		return json.MarshalIndent(data, "", "  ")
	}
	return json.Marshal(data)
}

func AsYAML(data interface{}) ([]byte, error) {
	return yaml.Marshal(data)
}
