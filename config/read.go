package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func ReadFile(filename string, data interface{}) error {
	fileExt, err := getValidExtension(filename)
	if err != nil {
		return err
	}

	fileData, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	switch fileExt {
	case ".yaml", ".yml":
		err = yaml.Unmarshal(fileData, data)
	case ".json":
		err = json.Unmarshal(fileData, data)
	default:
		return fmt.Errorf("EmailConfigReadFileError: Unknown ERROR reading %q", filename)
	}
	return err
}
