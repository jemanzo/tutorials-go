package emailconf

var (
	AUTH_METHOD = [...]string{`NONE`, `PLAIN`, `LOGIN`, `CRAM-MD5`, `DIGEST-MD5`}
)

type Authentication struct {
	Method   string `json:"method" yaml:"method"`
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
}

func newAuthentication() Authentication {
	return Authentication{
		Method:   "CRAM-MD5",
		Username: "",
		Password: "",
	}
}
