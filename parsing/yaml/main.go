package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type Config struct {
	filename string
	Server   struct {
		Username string `yaml:"username" json:"username"`
		Password string `yaml:"password" json:"password"`
		Hostname string `yaml:"hostname" json:"hostname"`
		Port     int    `yaml:"port" json:"port"`
	} `yaml:"server" json:"server"`
}

func (c *Config) SaveAsYAML() {
	file, err := os.Create("config_saved.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}
	enc := yaml.NewEncoder(file)
	enc.Encode(c)
	enc.Close()
}

func (c *Config) SaveAsJSON(ident bool) {
	file, err := os.Create("config_saved.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	enc := json.NewEncoder(file)
	if ident {
		enc.SetIndent("", "  ")
	}
	enc.Encode(c)
}

func (c *Config) readConfFile(filename string) *Config {
	c.filename = filename
	fileData, err := ReadConfFile(c.filename)
	if err != nil {
		log.Printf("Reading config.yaml Error %e", err)
	}
	if yaml.Unmarshal(fileData, c); err != nil {
		log.Fatalf("Unmarshal: %e", err)
	}
	return c
}

func ReadConfFile(filename string) (fileData []byte, err error) {
	log.Println(strings.Repeat("-", 32))
	defer log.Println(strings.Repeat("-", 32))

	fileData, err = ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	fmt.Println(string(fileData))
	return
}

func main() {
	filename := "config.yaml"

	var conf Config
	log.Printf("NOT_INITIALIZED (%p) %v", &conf, conf)

	conf.readConfFile(filename)

	log.Printf("INIT_AND_FILLED (%p) %v", &conf, conf)
	conf.SaveAsYAML()
	conf.SaveAsJSON(true)
}
