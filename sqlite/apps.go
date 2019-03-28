package sqlite

import (
	"database/sql"
	"github.com/diogox/GoLauncher/api/models"
)

func createAppsTable(db *sql.DB) error {
	query := "CREATE TABLE IF NOT EXISTS apps (exec TEXT PRIMARY KEY, name TEXT, description TEXT, IconName TEXT, accessCounter INT)"
	statement, _ := db.Prepare(query)
	_, err := statement.Exec()

	return err
}

func (ldb *LauncherDB) AddApp(app models.AppInfo) error {
	statement, err := ldb.db.Prepare("INSERT INTO apps (exec, name, description, IconName, accessCounter) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		panic(err)
	}
	_, err = statement.Exec(app.Exec, app.Name, app.Description, app.IconName, 0)

	return err
}

func (ldb *LauncherDB) IncrementAppAccessCounter(exec string) error {
	statement, err := ldb.db.Prepare("UPDATE apps SET accessCounter=accessCounter+1 WHERE exec=?")
	if err != nil {
		panic(err)
	}
	_, err = statement.Exec(exec)

	return err
}

func (ldb *LauncherDB) UpdateApp(app models.AppInfo) error {
	statement, err := ldb.db.Prepare("UPDATE apps SET exec=?, name=?, description=?, IconName=? WHERE exec=?")
	if err != nil {
		panic(err)
	}
	_, err = statement.Exec(app.Exec, app.Name, app.Description, app.IconName, app.Exec)

	return err
}

func (ldb *LauncherDB) GetAllApps(name string) ([]models.AppInfo, error) {
	query := "SELECT * FROM apps"
	rows, err := ldb.db.Query(query)
	if err != nil {
		return nil, err
	}

	var _exec string
	var _name string
	var _description string
	var _iconName string
	var _accessCounter int

	apps := make([]models.AppInfo, 0)
	for rows.Next() {
		err := rows.Scan(&_exec, &_name, &_description, &_iconName, &_accessCounter)
		if err != nil {
			panic(err)
		}
		appInfo := models.AppInfo{
			Name:          _name,
			Description:   _description,
			Exec:          _exec,
			IconName:      _iconName,
			AccessCounter: _accessCounter,
		}
		apps = append(apps, appInfo)
	}

	return apps, nil
}

func (ldb *LauncherDB) RemoveAllApps() error {
	statement, err := ldb.db.Prepare("DROP TABLE apps")
	if err != nil {
		panic(err)
	}
	_, err = statement.Exec()

	return err
}

func (ldb *LauncherDB) FindAppByID(exec string) (models.AppInfo, error) {
	rows, err := ldb.db.Query("SELECT * FROM apps WHERE exec=(?)", exec)
	if err != nil {
		return models.AppInfo{}, err
	}

	var _exec string
	var _name string
	var _description string
	var _iconName string
	var _accessCounter int

	for rows.Next() {
		rows.Scan(&_exec, &_name, &_description, &_iconName, &_accessCounter)
	}

	app := models.AppInfo{
		Exec:          _exec,
		Name:          _name,
		Description:   _description,
		IconName:      _iconName,
		AccessCounter: _accessCounter,
	}

	return app, nil
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
	var _accessCounter int

	apps := make([]models.AppInfo, 0)
	for rows.Next() {
		err := rows.Scan(&_exec, &_name, &_description, &_iconName, &_accessCounter)
		if err != nil {
			panic(err)
		}
		appInfo := models.AppInfo{
			Name:          _name,
			Description:   _description,
			Exec:          _exec,
			IconName:      _iconName,
			AccessCounter: _accessCounter,
		}
		apps = append(apps, appInfo)
	}

	return apps, nil
}
