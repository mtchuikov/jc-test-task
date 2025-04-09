package config

import "github.com/spf13/pflag"

var (
	serverAddrFlagDesc = "IP address and port for the server to listen on"
	serverAddrFlag     = pflag.String("server.addr", "127.0.0.1:8080", serverAddrFlagDesc)

	dbConnURLDesc = "URL to connect to the Postgres database"
	dbConnURLFlag = pflag.String("db.conn.url", "postgres://username:password@postgres:5432/postgres?sslmode=disable", dbConnURLDesc)
)

func (c *config) loadFromFlags() {
	pflag.CommandLine.SortFlags = false
	pflag.Parse()

	c.ServerAddr = *serverAddrFlag
	c.DBConnURL = *dbConnURLFlag

}
