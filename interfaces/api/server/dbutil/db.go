package dbutil

import (
	"database/sql"
	"fmt"
)

func Init(datasource string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", datasource)
	if err != nil {
		return nil, fmt.Errorf("failed db init. %s", err)
	}
	if err := createTodoTable(db); err != nil {
		return nil, err
	}
	return db, nil
}

func createTodoTable(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE todo(id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT, text TEXT)")
	return err
}
