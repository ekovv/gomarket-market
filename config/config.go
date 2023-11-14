package config

import (
	"flag"
	"os"
)

type Config struct {
	DB string
}

type F struct {
	db *string
}

var f F

func init() {
	f.db = flag.String("d", "", "-d=db")
}

func New() (c Config) {
	flag.Parse()
	if envDB := os.Getenv("DATABASE_DSN"); envDB != "" {
		f.db = &envDB
	}
	c.DB = *f.db
	return c
}
