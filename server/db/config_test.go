package db

import "testing"

func TestReadSchemas(t *testing.T) {
	schemas, err := readSchemas("./schema.sql")
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	if len(schemas) == 0 {
		t.Errorf("expected > 0, got %v", len(schemas))
	}
}

func TestConfigDatabase(t *testing.T) {
	conn := GetConnection(":memory:")

	err := ConfigDatabase(conn, "./schema.sql")
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}
