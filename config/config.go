package config

import (
	"fmt"
	"path"
	"strings"
)

type IConfigFile interface {
	GetFilename() string
	SetFilename(filename string) error
	ReadFile() error
	SaveFile() error
	AsJSON(ident bool) (string, error)
	AsYAML() (string, error)
}

type ConfigFile struct {
	filename string `json:"-" yaml:"-"`
}

func (c *ConfigFile) GetFilename() string {
	return c.filename
}

func (c *ConfigFile) SetFilename(filename string) error {
	c.filename = sanitizeFilename(filename)
	_, err := getValidExtension(filename)
	if err != nil {
		return err
	}
	c.filename = filename
	return nil
}

func (c *ConfigFile) SetAsYAML() {
	filename := c.filename
	filename = sanitizeFilename(filename)
	filename = forceFileExtension(filename, ".yaml")
	c.filename = filename
}

func (c *ConfigFile) SetAsJSON() {
	filename := c.filename
	filename = sanitizeFilename(filename)
	filename = forceFileExtension(filename, ".json")
	c.filename = filename
}

var validExtensions = []string{".yaml", ".yml", ".json"}

func getValidExtension(filename string) (string, error) {
	fileExt := strings.ToLower(path.Ext(filename))
	for _, v := range validExtensions {
		if fileExt == v {
			return fileExt, nil
		}
	}
	err := fmt.Errorf("SaveEmailConfigError: filename extension %q should be one of %v", fileExt, validExtensions)
	return fileExt, err
}

func sanitizeFilename(filename string) string {
	if filename == "" {
		filename = "config.yaml"
	}
	filename = path.Clean(filename)
	return filename
}

func forceFileExtension(filename, newExtension string) string {
	fileExt := path.Ext(filename)
	filename = filename[:len(filename)-len(fileExt)]
	// Force the file extension to be lowercase
	filename += strings.ToLower(newExtension)
	return filename
}
