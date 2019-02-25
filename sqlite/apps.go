package sqlite

import (
	"database/sql"
)

func createAppsTable(db *sql.DB) error {
	query := "CREATE TABLE IF NOT EXISTS apps (exec TEXT PRIMARY KEY, name TEXT, description TEXT)"
	statement, _ := db.Prepare(query)
	_, err := statement.Exec()

	return err
}

func (ldb *LauncherDB) AddApp(exec string, name string, descr string) error {
	statement, err := ldb.db.Prepare("INSERT INTO apps (exec, name, description) VALUES (?, ?, ?)")
	if err != nil {
		panic(err)
	}
	_, err = statement.Exec(exec, name, descr)

	return err
}

func (ldb *LauncherDB) FindAppByID(exec string) (string, error) {
	rows, err := ldb.db.Query("SELECT * FROM apps WHERE exec=(?)", exec)
	if err != nil {
		return "", err
	}

	var _exec string
	var _name string
	var _description string

	for rows.Next() {
		rows.Scan(&_exec, &_name, &_description)
	}

	return _name, nil
}

func (ldb *LauncherDB) FindAppByName(name string) ([]string, error) {
	query := "SELECT * FROM apps WHERE name LIKE '%" + name + "%'"
	rows, err := ldb.db.Query(query)
	if err != nil {
		return nil, err
	}

	var _exec string
	var _name string
	var _description string

	apps := make([]string, 0)
	for rows.Next() {
		err := rows.Scan(&_exec, &_name, &_description)
		if err != nil {
			panic(err)
		}
		apps = append(apps, _name)
	}

	return apps, nil
}