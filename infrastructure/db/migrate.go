package db

import (
	"database/sql"
	"io/fs"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/pressly/goose/v3"

	"github.com/thanishsid/goserver/config"
)

// Runs database migrations.
func Migrate(fileSys fs.FS, path string) error {
	db, err := sql.Open("pgx", config.C.PostgresSource)
	defer db.Close()
	if err != nil {
		return err
	}

	goose.SetBaseFS(fileSys)

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	if err := goose.Up(db, path); err != nil {
		return err
	}

	return nil
}
