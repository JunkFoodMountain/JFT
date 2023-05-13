package main

import (
	"JFT/internal/repository"
	"github.com/vrischmann/envconfig"
)

func main() {
	conf := repository.Config{}
	err := envconfig.Init(&conf)
	ifKill(err)

	db, err := repository.OpenDB(conf)
	ifKill(err)

	err = repository.MigrateDB(db)
	ifKill(err)
}

func ifKill(err error) {
	if err != nil {
		panic(err)
	}
}
