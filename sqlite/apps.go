package sqlite

import (
	"database/sql"
	"github.com/diogox/GoLauncher/models"
)

func createAppsTable(db *sql.DB) error {
	query := "CREATE TABLE IF NOT EXISTS apps (exec TEXT PRIMARY KEY, name TEXT, description TEXT, IconName TEXT)"
	statement, _ := db.Prepare(query)
	_, err := statement.Exec()

	return err
}

func (ldb *LauncherDB) AddApp(exec string, name string, descr string, _iconName string) error {
	statement, err := ldb.db.Prepare("INSERT INTO apps (exec, name, description, IconName) VALUES (?, ?, ?, ?)")
	if err != nil {
		panic(err)
	}
	_, err = statement.Exec(exec, name, descr, _iconName)

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
	var _iconName string

	for rows.Next() {
		rows.Scan(&_exec, &_name, &_description, &_iconName)
	}

	return _name, nil
}

func (ldb *LauncherDB) FindAppByName(name string) ([]models.AppInfo, error) {
	query := "SELECT * FROM apps WHERE name LIKE '%" + name + "%'"
	rows, err := ldb.db.Query(query)
	if err != nil {
		return nil, err
	}

	var _exec string
	var _name string
	var _description string
	var _iconName string

	apps := make([]models.AppInfo, 0)
	for rows.Next() {
		err := rows.Scan(&_exec, &_name, &_description, &_iconName)
		if err != nil {
			panic(err)
		}
		appInfo := models.AppInfo{
			Name: _name,
			Description: _description,
			Exec: _exec,
			IconName: _iconName,
		}
		apps = append(apps, appInfo)
	}

	return apps, nil
}