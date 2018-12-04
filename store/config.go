package store

import (
	"log"

	"github.com/BurntSushi/toml"
)

var (
	config = Config{}
	DAOS   = DAO{}
)

func init() {
	config.Read()
	DAOS.Server = config.Server
	DAOS.Database = config.Database
	DAOS.Connect()
}

type Config struct {
	Server   string
	Database string
}

func (c *Config) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}
