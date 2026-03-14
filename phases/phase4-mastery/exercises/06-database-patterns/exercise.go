package exercise_06

import (
	"database/sql"
	"fmt"
)

// User is a database model.
type User struct {
	ID    int
	Name  string
	Email string
}

// UserRepo manages users in a SQL database.
type UserRepo struct {
	db *sql.DB
}

// NewUserRepo creates a UserRepo.
func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

// CreateTable creates the users table if it doesn't exist.
// TODO: db.Exec CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT, email TEXT UNIQUE)
func (r *UserRepo) CreateTable() error {
	return fmt.Errorf("not implemented")
}

// Insert adds a user and returns the assigned ID.
// TODO: db.Exec INSERT, LastInsertId
func (r *UserRepo) Insert(name, email string) (int64, error) {
	return 0, fmt.Errorf("not implemented")
}

// FindByEmail finds a user by email.
// Returns sql.ErrNoRows if not found.
// TODO: db.QueryRow SELECT, Scan
func (r *UserRepo) FindByEmail(email string) (User, error) {
	return User{}, fmt.Errorf("not implemented")
}

// ListAll returns all users.
// TODO: db.Query SELECT, rows.Next + Scan
func (r *UserRepo) ListAll() ([]User, error) {
	return nil, fmt.Errorf("not implemented")
}

var _ = sql.ErrNoRows
