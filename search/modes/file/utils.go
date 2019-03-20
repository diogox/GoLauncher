package file

import (
	"github.com/diogox/GoLauncher/api"
	"github.com/mitchellh/go-homedir"
	"os"
	"path"
)

// TODO: This icon doesn't exist!
const FILE_ICON = "file"
const DIR_ICON = "nemo"

func isFile(filepath string) bool {
	fi, err := os.Stat(filepath)
	if err != nil {
		return false
	}

	return !fi.Mode().IsDir()
}

func isDir(filepath string) bool {
	fi, err := os.Stat(filepath)
	if err != nil {
		return false
	}

	return fi.Mode().IsDir()
}

func getAbsPath(pathStr string) (string, error) {
	if !path.IsAbs(pathStr) {
		absPath, err := homedir.Expand(pathStr)
		if err != nil {
			return "", err
		}
		pathStr = absPath
	}
	return pathStr, nil
}

func getN(results []api.Result, nOfResults int) []api.Result {
	_results := make([]api.Result, 0)
	for i, r := range results {

		// Ensure maximum number of results
		if i >= nOfResults {
			break
		}

		// Add result
		_results = append(_results, r)
	}

	return _results
}
