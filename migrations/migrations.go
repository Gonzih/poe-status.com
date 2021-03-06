package migrations

import (
	"errors"
	"fmt"
	"log"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func SQLMigrationsFolderPath(migrationsFolderAbsPath string) string {
	path, err := filepath.Abs(filepath.Join(migrationsFolderAbsPath, "/sql"))
	if err != nil {
		log.Fatal(err)
	}

	return path
}

func FormatSQLError(err error) error {
	if err == nil {
		return nil
	}

	switch err := err.(type) {
	case database.Error:
		return errors.New(err.Error())
	default:
		if err.Error() == "no change" {
			log.Printf("Ignoring SQL error \"%s\" during migrations run", err)
			return nil
		}
		return err
	}
}

func Up(migrationsFolderAbsPath, databaseURL string) error {
	return Run(migrationsFolderAbsPath, databaseURL, "up")
}

func Down(migrationsFolderAbsPath, databaseURL string) error {
	return Run(migrationsFolderAbsPath, databaseURL, "down")
}

func Run(migrationsFolderAbsPath, databaseURL, direction string) error {
	path := SQLMigrationsFolderPath(migrationsFolderAbsPath)
	log.Printf("Running migrations from %s, with db %s", path, databaseURL)
	m, err := migrate.New(fmt.Sprintf("file://%s", path), databaseURL)

	if err != nil {
		return FormatSQLError(err)
	}

	switch direction {
	case "up":
		if err := m.Up(); err != nil {
			return FormatSQLError(err)
		}
	case "down":
		if err := m.Down(); err != nil {
			return FormatSQLError(err)
		}
	default:
		return fmt.Errorf("Wrong migration direction %s", direction)
	}

	return nil
}
