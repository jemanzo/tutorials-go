package main

import (
	"fmt"
	"path"

	emailconf "github.com/jemanzo/tutorials-go/email/config"
)

func CreateConfigSample(filename string) (*emailconf.ConfigFile, error) {
	if filename == "" {
		filename = "./email.sample.yaml"
	}

	config, err := emailconf.NewConfigFile(filename)
	if err != nil {
		return nil, err
	}

	email, _ := config.AddEmail("contact@example.com")
	if err != nil {
		fmt.Println(err)
		return config, err
	}

	email.SetSMTP("smtp.example.com")
	email.SetIMAP("mail.example.com")

	err = config.SaveFile()
	if err != nil {
		fmt.Println(err)
		return config, err
	}

	return config, nil
}

func ReadConfig(filename string) (*emailconf.ConfigFile, error) {
	config, _ := emailconf.NewConfigFile(filename)
	err := config.ReadFile()
	return config, err
}

func SaveYAMLAndJSON(conf *emailconf.ConfigFile) {
	// YAML
	conf.SetAsYAML()
	fmt.Printf(" trying to save file %q", conf.GetFilename())
	if err := conf.SaveFile(); err != nil {
		fmt.Printf(" successfully saved as %q", path.Ext(conf.GetFilename()))
	}

	// JSON
	conf.SetAsJSON()
	fmt.Printf(" trying to save file %q", conf.GetFilename())
	if err := conf.SaveFile(); err != nil {
		fmt.Printf(" successfully saved as %q", path.Ext(conf.GetFilename()))
	}
	conf.SaveFile()
}

func PrintEmailInfo(config *emailconf.ConfigFile, emailAddress string) {
	email := config.Emails[emailAddress]
	fmt.Println(email)
}
