package sqlite

import (
	"database/sql"
	"github.com/diogox/GoLauncher/api"
)

func createPreferencesTable(db *sql.DB) error {
	query := "CREATE TABLE IF NOT EXISTS preferences (preference TEXT PRIMARY KEY, prefValue TEXT)"
	statement, _ := db.Prepare(query)
	_, err := statement.Exec()

	return err
}

func (ldb *LauncherDB) LoadDefaultPreferences() []error {
	errors := make([]error, 0)

	err := ldb.AddPreference(api.PreferenceHotkey, "<Ctrl>space")
	if err != nil {
		errors = append(errors, err)
	}

	err = ldb.AddPreference(api.PreferenceKeepInputOnHide, api.PreferenceFALSE)
	if err != nil {
		errors = append(errors, err)
	}

	err = ldb.AddPreference(api.PreferenceLaunchAtStartUp, api.PreferenceFALSE)
	if err != nil {
		errors = append(errors, err)
	}

	err = ldb.AddPreference(api.PreferenceNResultsToShow, "9")
	if err != nil {
		errors = append(errors, err)
	}

	err = ldb.AddPreference(api.PreferenceNAppResults, "10")
	if err != nil {
		errors = append(errors, err)
	}

	return nil
}

func (ldb *LauncherDB) AddPreference(preference string, value string) error {
	statement, err := ldb.db.Prepare("INSERT INTO preferences (preference, prefValue) VALUES (?, ?)")
	if err != nil {
		panic(err)
	}
	_, err = statement.Exec(preference, value)

	return err
}

func (ldb *LauncherDB) SetPreference(preference string, value string) error {
	statement, err := ldb.db.Prepare("UPDATE preferences SET prefValue=? WHERE preference=?")
	if err != nil {
		panic(err)
	}

	_, err = statement.Exec(value, preference)

	return err
}

func (ldb *LauncherDB) GetPreference(preference string) string {
	rows, err := ldb.db.Query("SELECT * FROM preferences WHERE preference=(?)", preference)
	if err != nil {
		return ""
	}

	var _preference string
	var value string
	for rows.Next() {
		rows.Scan(&_preference, &value)
	}

	return value
}
