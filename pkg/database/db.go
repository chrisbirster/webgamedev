package database

import (
	"database/sql"

	_ "github.com/libsql/libsql-client-go/libsql"
	_ "modernc.org/sqlite"
)

const (
	createMigrationsTable  = `CREATE TABLE IF NOT EXISTS migrations (version INTEGER PRIMARY KEY)`
	createUsersTable       = `CREATE TABLE IF NOT EXISTS users (name TEXT NOT NULL)`
	createDeezTable        = `CREATE TABLE IF NOT EXISTS deez (what TEXT NOT NULL)`
	queryMigrationsVersion = `SELECT MAX(version) FROM migrations`
)

func InitDB(url string) (*sql.DB, error) {
	db, err := sql.Open("libsql", url)
	if err != nil {
		return nil, err
	}

	// Create migrations table if it doesn't exist
	_, err = db.Exec(createMigrationsTable)
	if err != nil {
		return nil, err
	}

	// Get the current migration version
	var version *int
	err = db.QueryRow(queryMigrationsVersion).Scan(&version)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	// Apply migrations based on the current version
	currentVersion := 0
	if version != nil {
		currentVersion = *version
	}
	switch currentVersion {
	case 0:
		// Apply initial migrations and data seeding
		if err := applyInitialMigrations(db); err != nil {
			return nil, err
		}
		_, err = db.Exec(`INSERT INTO migrations (version) VALUES (1)`)
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}

func applyInitialMigrations(db *sql.DB) error {
	// Your initial DB setup and data seeding
	if _, err := db.Exec(createUsersTable); err != nil {
		return err
	}
	if _, err := db.Exec(createDeezTable); err != nil {
		return err
	}
	if _, err := db.Exec(`INSERT INTO users (name) VALUES ("GM")`); err != nil {
		return err
	}
	if _, err := db.Exec(`INSERT INTO users (name) VALUES ("Bongo")`); err != nil {
		return err
	}
	if _, err := db.Exec(`INSERT INTO deez (what) VALUES ("NUTS")`); err != nil {
		return err
	}

	return nil
}
