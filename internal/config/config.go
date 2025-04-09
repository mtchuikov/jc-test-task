package config

var conf config = config{ServiceName: "shortener"}

type config struct {
	ServiceName string
	ServerAddr  string `env:"SERVER_ADDRESS"`
}

func ServerAddr() string {
	return ""
}

func Verbose() bool {
	return true
}
