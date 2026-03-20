package db

import (
	"database/sql"
	"os"
	"strings"
)

func readSchemas(schemaUrl string) ([]string, error) {
	data, err := os.ReadFile(schemaUrl)
	if err != nil {
		return nil, err
	}

	text := string(data)
	schemas := strings.Split(text, ";")

	return schemas, nil
}

func ConfigDatabase(conn *sql.DB, schemaUrl string) error {
	schemas, err := readSchemas(schemaUrl)
	if err != nil {
		return err
	}

	for _, schema := range schemas {
		_, err := conn.Exec(schema)
		if err != nil {
			return err
		}
	}

	_, err = conn.Exec("PRAGMA foreign_keys=ON;")
	if err != nil {
		return err
	}

	_, err = conn.Exec("PRAGMA journal_mode=WAL;")
	if err != nil {
		return err
	}

	_, err = conn.Exec("PRAGMA synchronous=NORMAL;")
	if err != nil {
		return err
	}

	return nil
}
