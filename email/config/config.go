package emailconf

import "github.com/jemanzo/tutorials-go/config"

type Emails map[string]*EmailConfig

type ConfigFile struct {
	config.ConfigFile `json:"-" yaml:"-"`
	Emails            Emails `json:"emails" yaml:"emails"`
}

func (c *ConfigFile) AddEmail(emailAddr string) (*EmailConfig, error) {
	email := EmailConfig{}
	c.Emails[emailAddr] = &email
	return &email, nil
}

func (c *ConfigFile) ReadFile() error {
	return config.ReadFile(c.GetFilename(), c)
}

func (c *ConfigFile) SaveFile() error {
	return config.SaveFile(c.GetFilename(), c)
}

func (c *ConfigFile) AsYAML() (string, error) {
	data, err := config.AsYAML(c)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (c *ConfigFile) AsJSON(ident bool) (string, error) {
	data, err := config.AsJSON(c, ident)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func NewConfigFile(filename string) (*ConfigFile, error) {
	conf := ConfigFile{
		Emails: make(Emails),
	}
	if err := conf.SetFilename(filename); err != nil {
		return nil, err
	}
	return &conf, nil
}
