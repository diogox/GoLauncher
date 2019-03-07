package sqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

const DBPath = "./go-launcher.db"

func NewLauncherDB() LauncherDB {

	// Open DB
	db, err := sql.Open("sqlite3", DBPath)
	if err != nil {
		panic(err)
	}
	//defer db.Close() // TODO: Have this outside of the function?

	// Initialize tables, if not there
	err = createAppsTable(db)
	if err != nil {
		panic(err)
	}
	err = createPreferencesTable(db)
	if err != nil {
		panic(err)
	}
	err = createShortcutsTable(db)
	if err != nil {
		panic(err)
	}
	err = createExtensionsTable(db)
	if err != nil {
		panic(err)
	}

	return LauncherDB{
		db: db,
	}
}

type LauncherDB struct {
	db *sql.DB
}
