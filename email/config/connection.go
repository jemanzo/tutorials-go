package emailconf

import "fmt"

var (
	ENCRIPTION = [...]string{`Auto`, `StartTLS`, `SSL`}
)

type Connection struct {
	Host       string `json:"host" yaml:"host"`
	Port       int    `json:"port" yaml:"port"`
	Encryption string `json:"encryption" yaml:"encryption"`
}

func (c *Connection) GetURL() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

func newConnection(protocol, host string) Connection {
	var (
		port       int
		encryption string
	)
	switch protocol {
	case "smtp":
		port = 587
		encryption = "StartTLS"
	case "imap":
		port = 993
		encryption = "SSL"
	case "pop3":
		port = 995
		encryption = "SSL"
	}
	return Connection{
		Host:       host,
		Port:       port,
		Encryption: encryption,
	}
}
