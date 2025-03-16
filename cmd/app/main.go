package main

import (
	"log"

	"github.com/nzoschke/gots/pkg/api"
	"github.com/nzoschke/gots/pkg/db"
	"github.com/pkg/errors"
)

func main() {
	if err := mainErr(); err != nil {
		log.Fatalf("ERR: %+v\n", err)
	}
}

func mainErr() error {
	if err := db.New(); err != nil {
		return errors.WithStack(err)
	}

	if err := api.New(":1234", ":3000"); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
