package exercise_06

import (
	"database/sql"
	"errors"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func newTestDB(t *testing.T) *sql.DB {
	t.Helper()
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	t.Cleanup(func() { db.Close() })
	return db
}

func TestUserRepo(t *testing.T) {
	db := newTestDB(t)
	repo := NewUserRepo(db)

	if err := repo.CreateTable(); err != nil {
		t.Fatalf("CreateTable: %v", err)
	}

	id, err := repo.Insert("Alice", "alice@example.com")
	if err != nil {
		t.Fatalf("Insert: %v", err)
	}
	if id == 0 {
		t.Error("Insert should return non-zero ID")
	}

	repo.Insert("Bob", "bob@example.com")

	user, err := repo.FindByEmail("alice@example.com")
	if err != nil {
		t.Fatalf("FindByEmail: %v", err)
	}
	if user.Name != "Alice" {
		t.Errorf("user.Name = %q, want Alice", user.Name)
	}

	_, err = repo.FindByEmail("ghost@example.com")
	if !errors.Is(err, sql.ErrNoRows) {
		t.Errorf("FindByEmail(missing) should return sql.ErrNoRows, got %v", err)
	}

	users, err := repo.ListAll()
	if err != nil {
		t.Fatalf("ListAll: %v", err)
	}
	if len(users) != 2 {
		t.Errorf("ListAll len = %d, want 2", len(users))
	}
}
