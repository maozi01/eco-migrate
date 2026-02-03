//go:build go1.16

package iofs_test

import (
	"embed"
	"log"

	migrate "github.com/maozi01/eco-migrate"
	_ "github.com/maozi01/eco-migrate/database/postgres"
	"github.com/maozi01/eco-migrate/source/iofs"
)

//go:embed testdata/migrations/*.sql
var fs embed.FS

func Example() {
	d, err := iofs.New(fs, "testdata/migrations")
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithSourceInstance("iofs", d, "postgres://postgres@localhost/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	err = m.Up()
	if err != nil {
		// ...
	}
	// ...
}
