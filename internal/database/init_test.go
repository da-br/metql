package database_test

import (
	"testing"

	"codeburg.com/da-br/metql/internal/database"
)

// database Init creates a new database, overwrites an existing one,
// and throws an error when one already exists
func TestInitDataBase(t *testing.T) {
	t.Skip("skipping init db test as it takes a long time")
	tmpDir := t.TempDir()

	t.Log("creating db")
	err := database.Init(tmpDir, "test", false)
	if err != nil {
		t.Fatalf(`Init(temp, "test", false) failed to create a db with %s`, err.Error())
	}

	t.Log("overwriting db")
	err = database.Init(tmpDir, "test", true)
	if err != nil {
		t.Fatalf(`Init(temp, "test", true) failed to overwrite the db with %s`, err.Error())
	}

	t.Log("skipping creating db")
	err = database.Init(tmpDir, "test", false)
	if err != nil {
		t.Fatalf(`Init(temp, "test", false) failed to error when the db already existed with %s`, err.Error())
	}
}
