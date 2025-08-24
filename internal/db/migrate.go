package db

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func (s *Store) Migrate() error {
	driver, err := postgres.WithInstance(s.db.DB, &postgres.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"postgres", driver)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Println("No change in migration")
			return nil
		}
		return err
	}

	return nil

}
