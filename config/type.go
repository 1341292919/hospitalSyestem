package config

type config struct {
	MySQL mySQL
}

type mySQL struct {
	Addr     string
	Database string
	Username string
	Password string
	Charset  string
}

type service struct {
	Name     string
	AddrList []string
	LB       bool `mapstructure:"load-balance"`
}
