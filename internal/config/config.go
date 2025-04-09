package config

import "github.com/caarlos0/env/v11"

var conf config = config{ServiceName: "js-test-task"}

type config struct {
	ServiceName string
	ServerAddr  string `env:"SERVER_ADDRESS"`
	DBConnURL   string `env:"DB_CONN_URL"`
}

func Init() {
	conf.loadFromFlags()
	env.Parse(&conf)
}

func ServiceName() string {
	return conf.ServiceName
}

func ServerAddr() string {
	return conf.ServerAddr
}

func DBConnURL() string {
	return conf.DBConnURL
}
