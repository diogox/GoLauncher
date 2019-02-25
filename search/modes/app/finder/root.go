package finder

import (
	"bufio"
	"github.com/diogox/GoLauncher/sqlite"
	"io/ioutil"
	"strings"
)

const BasePath = "/usr/share/applications/"

// Finds all available '.desktop' file paths
func FindApps(db *sqlite.LauncherDB) []AppInfo {
	desktopFiles := getDesktopFiles()

	apps := make([]AppInfo, 0)
	for _, file := range desktopFiles {
		appInfo := getAppInfo(file)

		// Check if exists
		res, err := db.FindAppByID(appInfo.Exec)
		if res == "" {
			err = db.AddApp(appInfo.Exec, appInfo.Name, appInfo.Description)
			if err != nil {
				panic(err)
			}
		}
	}

	return apps
}

func getDesktopFiles() []string {
	fileInfo, err := ioutil.ReadDir(BasePath)
	if err != nil {
		panic(err)
	}

	files := make([]string, 0)
	for _, info := range fileInfo {
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".desktop") {
			files = append(files, BasePath + info.Name())
		}
	}

	return files
}

func getAppInfo(filePath string) AppInfo {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	contents := string(data)
	scanner := bufio.NewScanner(strings.NewReader(contents))

	appInfo := AppInfo{
		Exec: "",
		Name: "",
		Description: "",
		IconPath: "",
	}

	isExcerpt := false
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "[Desktop Entry]") {
			isExcerpt = true
			continue
		} else if isExcerpt && strings.HasPrefix(line, "[") {
			break
		}

		if isExcerpt {
			keyValue := strings.Split(line, "=")
			if len(keyValue) != 2 {
				continue
			}

			key := keyValue[0]
			value := keyValue[1]
			switch key {
			case "Exec":
				appInfo.Exec = value
			case "Name":
				appInfo.Name = value
			case "Comment":
				appInfo.Description = value
			}
		}
	}

	return appInfo
}