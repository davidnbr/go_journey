# Exercise 06: Database Patterns

## Learning Objectives
- Use `database/sql` for SQL database access
- Use prepared statements via `db.QueryRow` and `db.Query`
- Handle `sql.ErrNoRows` for missing records
- Test with an in-memory SQLite database

## Why This Matters (DevOps context)
Most backend services persist data in SQL databases. Go's `database/sql` is the standard abstraction. Using `sql.ErrNoRows` explicitly (not string matching) is the correct pattern for "not found" handling.

## Concept
`db.Exec` for writes. `db.QueryRow(...).Scan(...)` for single rows. `db.Query` + `rows.Next()` + `rows.Scan()` for multiple rows. Always `defer rows.Close()`. SQLite with `go-sqlite3` driver for testing.

## Note on Dependencies
This exercise requires the `github.com/mattn/go-sqlite3` driver. Run:
```bash
cd /path/to/go_journey && go get github.com/mattn/go-sqlite3
```
This requires CGO (a C compiler). If unavailable, skip this exercise and use the pure-Go `modernc.org/sqlite` instead.

## Your Task
Implement `CreateTable`, `Insert`, `FindByEmail`, and `ListAll`.

## Run Tests
```bash
go test ./phases/phase4-mastery/exercises/06-database-patterns/
```

## Hints
<details><summary>Hint 1</summary>CreateTable: `r.db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL, email TEXT UNIQUE NOT NULL)")`</details>
<details><summary>Hint 2</summary>Insert: `result, err := r.db.Exec("INSERT INTO users(name,email) VALUES(?,?)", name, email); id, _ := result.LastInsertId()`</details>
<details><summary>Hint 3</summary>FindByEmail: `row := r.db.QueryRow("SELECT id,name,email FROM users WHERE email=?", email); return user, row.Scan(&user.ID, &user.Name, &user.Email)`</details>

## References
- [pkg.go.dev — database/sql](https://pkg.go.dev/database/sql)
- [Go Wiki — SQLDrivers](https://go.dev/wiki/SQLDrivers)
