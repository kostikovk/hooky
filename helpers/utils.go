package helpers

import "os"

// dirExists checks if a directory exists and is a directory.
func dirExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}

	return info.IsDir()
}

// getAbsolutePath returns the absolute path of a file.
func getAbsolutePath(path string) string {
	wd, err := os.Getwd()
	if err != nil {
		return path
	}

	return wd + "/" + path
}

// contains checks if a slice contains a string.
func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}

	return false
}
