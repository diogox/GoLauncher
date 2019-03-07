package sqlite

import (
	"database/sql"
	"errors"
	"github.com/diogox/GoLauncher/api/models"
)

func createExtensionsTable(db *sql.DB) error {
	query := "CREATE TABLE IF NOT EXISTS extensions (keyword TEXT PRIMARY KEY, name TEXT, description TEXT, icon_name TEXT, developer_name TEXT)"
	statement, _ := db.Prepare(query)
	_, err := statement.Exec()

	return err
}

func (ldb *LauncherDB) AddExtension(extension models.Extension) error {
	statement, err := ldb.db.Prepare("INSERT INTO extensions (keyword, name, description, icon_name, developer_name) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		panic(err)
	}
	_, err = statement.Exec(extension.Keyword, extension.Name, extension.Description, extension.IconName, extension.DeveloperName)

	return err
}

func (ldb *LauncherDB) DeleteExtension(extension models.Extension) error {
	statement, err := ldb.db.Prepare("DELETE FROM extensions WHERE keyword=?")
	if err != nil {
		panic(err)
	}
	_, err = statement.Exec(extension.Keyword)
	return err
}

func (ldb *LauncherDB) GetAllExtensions() ([]models.Extension, error) {
	var extensions []models.Extension

	rows, err := ldb.db.Query("SELECT * FROM extensions")
	if err != nil {
		return extensions, err
	}

	var _keyword string
	var _name string
	var _description string
	var _iconName string
	var _developerName string

	for rows.Next() {
		rows.Scan(&_keyword, &_name, &_description, &_iconName, &_developerName)

		extension := models.Extension{
			Keyword:       _keyword,
			Name:          _name,
			Description:   _description,
			IconName:      _iconName,
			DeveloperName: _developerName,
		}
		extensions = append(extensions, extension)
	}

	return extensions, nil
}

func (ldb *LauncherDB) FindExtensionByKeyword(keyword string) (models.Extension, error) {
	rows, err := ldb.db.Query("SELECT * FROM extensions WHERE keyword=(?)", keyword)
	if err != nil {
		return models.Extension{}, errors.New("failed to query extension")
	}

	var _keyword string
	var _name string
	var _description string
	var _iconName string
	var _developerName string
	for rows.Next() {
		rows.Scan(&_keyword, &_name, &_description, &_iconName, &_developerName)
	}

	return models.Extension{
		Keyword:       _keyword,
		Name:          _name,
		Description:   _description,
		IconName:      _iconName,
		DeveloperName: _developerName,
	}, nil
}
