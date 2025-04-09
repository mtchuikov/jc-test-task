package config

import "github.com/spf13/pflag"

var (
	serverAddrFlagDesc = "Specify the IP address and port for the server to listen on"
	serverAddrFlag     = pflag.StringP("server-addr", "a", "127.0.0.1:8080", serverAddrFlagDesc)

	dbConnURLDesc = "Define the database url to connect to the database"
	dbConnURLFlag = pflag.StringP("dsn", "d", "postgres://username:password@postgres:5432/postgres?sslmode=disable", dbConnURLDesc)
)

func (c *config) loadFromFlags() {
	pflag.CommandLine.SortFlags = false
	pflag.Parse()

	c.ServerAddr = *serverAddrFlag
	c.DBConnURL = *dbConnURLFlag

}
