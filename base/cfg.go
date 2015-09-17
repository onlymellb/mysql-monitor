package base

type Cfg struct {
	LogFile      string
	SocketDir    string
	LogLevel     int
	FalconClient string
	Endpoint     string
	Appname      string

	User string
	Pass string
	Host string
}

func NewConfig(sdir, user, passwd string) *Cfg {
	cfg := &Cfg{
		LogFile:      "mymon.log",
		SocketDir:    sdir,
		LogLevel:     5,
		FalconClient: "http://127.0.0.1:1988/v1/push",
		Appname:      "",
		Host:         "127.0.0.1",
		Endpoint:     "",
		User:         user,
		Pass:         passwd,
	}
	return cfg
}
