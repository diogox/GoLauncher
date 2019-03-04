package sqlite

import (
	"database/sql"
	"github.com/diogox/GoLauncher/api"
	"github.com/diogox/GoLauncher/api/models"
)

func createShortcutsTable(db *sql.DB) error {
	query := "CREATE TABLE IF NOT EXISTS shortcuts (keyword TEXT PRIMARY KEY, name TEXT, cmd TEXT, icon_name TEXT, is_default_search BIT, is_active BIT)"
	statement, _ := db.Prepare(query)
	_, err := statement.Exec()

	return err
}

func LoadDefaultShortcuts(db *api.DB) {
	shortcuts := []models.ShortcutInfo{
		{
			Keyword:         "g",
			Name:            "Google",
			Cmd:             "https://www.google.com/search?q=%s",
			IconName:        "google",
			IsDefaultSearch: true,
			IsActive:        true,
		},
		{
			Keyword:         "so",
			Name:            "Stack Overflow",
			Cmd:             "http://stackoverflow.com/search?q=%s",
			IconName:        "stack",
			IsDefaultSearch: true,
			IsActive:        true,
		},
		{
			Keyword:         "wiki",
			Name:            "Wikipedia",
			Cmd:             "https://en.wikipedia.org/wiki/%s",
			IconName:        "wikipedia",
			IsDefaultSearch: true,
			IsActive:        true,
		},
	}

	for _, shortcut := range shortcuts {
		_ = (*db).AddShortcut(shortcut)
	}
}

func (ldb *LauncherDB) AddShortcut(shortcut models.ShortcutInfo) error {
	statement, err := ldb.db.Prepare("INSERT INTO shortcuts (keyword, name, cmd, icon_name, is_default_search, is_active) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		panic(err)
	}

	var isDefaultSearchBit int
	if shortcut.IsDefaultSearch {
		isDefaultSearchBit = 1
	} else {
		isDefaultSearchBit = 0
	}

	var IsActiveBit int
	if shortcut.IsActive {
		IsActiveBit = 1
	} else {
		IsActiveBit = 0
	}

	_, err = statement.Exec(shortcut.Keyword, shortcut.Name, shortcut.Cmd, shortcut.IconName, isDefaultSearchBit, IsActiveBit)
	return err
}

func (ldb *LauncherDB) DeleteShortcut(shortcut models.ShortcutInfo) error {
	statement, err := ldb.db.Prepare("DELETE FROM shortcuts WHERE keyword=?")
	if err != nil {
		panic(err)
	}
	_, err = statement.Exec(shortcut.Keyword)
	return err
}

func (ldb *LauncherDB) GetAllShortcuts() ([]models.ShortcutInfo, error) {
	var shortcuts []models.ShortcutInfo

	rows, err := ldb.db.Query("SELECT * FROM shortcuts")
	if err != nil {
		return shortcuts, err
	}

	var _keyword string
	var _name string
	var _cmd string
	var _iconName string
	var _isDefaultSearchBit int
	var _isActiveBit int

	for rows.Next() {
		rows.Scan(&_keyword, &_name, &_cmd, &_iconName, &_isDefaultSearchBit, &_isActiveBit)

		var _isDefaultSearch bool
		if _isDefaultSearchBit == 1 {
			_isDefaultSearch = true
		} else {
			_isDefaultSearch = false
		}

		var _isActive bool
		if _isActiveBit == 1 {
			_isActive = true
		} else {
			_isActive = false
		}

		shortcut := models.ShortcutInfo{
			Keyword:         _keyword,
			Name:            _name,
			Cmd:             _cmd,
			IconName:        _iconName,
			IsDefaultSearch: _isDefaultSearch,
			IsActive:        _isActive,
		}
		shortcuts = append(shortcuts, shortcut)
	}

	return shortcuts, nil
}

func (ldb *LauncherDB) FindShortcutByKeyword(keyword string) (models.ShortcutInfo, error) {
	rows, err := ldb.db.Query("SELECT * FROM shortcuts WHERE keyword=(?)", keyword)
	if err != nil {
		return models.ShortcutInfo{}, err
	}

	var _keyword string
	var _name string
	var _cmd string
	var _iconName string
	var _isDefaultSearch bool
	var _isActive bool

	for rows.Next() {
		rows.Scan(&_keyword, &_name, &_cmd, &_iconName, &_isDefaultSearch, &_isActive)
	}

	return models.ShortcutInfo{
		Keyword:         _keyword,
		Name:            _name,
		Cmd:             _cmd,
		IconName:        _iconName,
		IsDefaultSearch: _isDefaultSearch,
		IsActive:        _isActive,
	}, nil
}
