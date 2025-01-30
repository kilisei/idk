package migrator

import (
	"database/sql"
	"embed"
	_ "embed"
	"github.com/pressly/goose/v3"
	"log"
)

//go:embed migrations/*.sql
var migrations embed.FS

func Migrate(db *sql.DB) {
	goose.SetBaseFS(migrations)
	if err := goose.SetDialect("sqlite3"); err != nil {
		log.Fatal(err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		log.Fatal(err)
	}
}
