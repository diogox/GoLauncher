package sqlite

import (
	"database/sql"
	"errors"
)

func createSessionsTable(db *sql.DB) error {
	query := "CREATE TABLE IF NOT EXISTS sessions (query TEXT PRIMARY KEY, execName TEXT)"
	statement, _ := db.Prepare(query)
	_, err := statement.Exec()

	return err
}

func (ldb *LauncherDB) AddSession(input string, appExec string) error {
	statement, err := ldb.db.Prepare("REPLACE INTO sessions (query, execName) VALUES (?, ?)")
	if err != nil {
		return err
	}
	_, err = statement.Exec(input, appExec)

	return err
}

func (ldb *LauncherDB) GetSession(input string) (string, error) {
	rows, err := ldb.db.Query("SELECT * FROM sessions WHERE query=(?)", input)
	if err != nil {
		return "", err
	}

	var _input string
	var _execName string

	for rows.Next() {
		err := rows.Scan(&_input, &_execName)
		if err != nil {
			return "", err
		}
	}

	if _execName == "" {
		return "", errors.New("session not found")
	}

	return _execName, nil
}
