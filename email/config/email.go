package emailconf

type IMAP = Protocol
type POP3 = Protocol
type SMTP = Protocol

type Protocol struct {
	Connection     `json:"conn" yaml:"conn"`
	Authentication `json:"auth" yaml:"auth"`
}

type EmailConfig struct {
	SMTP *Protocol `json:"smtp,omitempty" yaml:"smtp,omitempty"`
	IMAP *Protocol `json:"imap,omitempty" yaml:"imap,omitempty"`
	POP3 *Protocol `json:"pop3,omitempty" yaml:"pop3,omitempty"`
}

func (e *EmailConfig) SetSMTP(host string) *EmailConfig {
	e.SMTP = newDefaultConfig("smtp", host)
	return e
}

func (e *EmailConfig) SetIMAP(host string) *EmailConfig {
	e.IMAP = newDefaultConfig("imap", host)
	return e
}

func (e *EmailConfig) SetPOP3(host string) *EmailConfig {
	e.POP3 = newDefaultConfig("pop3", host)
	return e
}

func newDefaultConfig(protocol, host string) *Protocol {
	return &Protocol{
		newConnection(protocol, host),
		newAuthentication(),
	}
}
