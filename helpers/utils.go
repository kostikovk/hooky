package helpers

import (
	"fmt"
	"os"
)

func dirExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}

	return info.IsDir()
}

func getAbsolutePath(path string) string {
	wd, err := os.Getwd()
	if err != nil {
		return path
	}

	return wd + "/" + path
}

func Contains[T any](arr []T, compare func(T) bool) bool {
	for _, a := range arr {
		if compare(a) {
			return true
		}
	}
	return false
}

func ContainsFile(arr []os.DirEntry, fileName string) bool {
	return Contains(arr, func(f os.DirEntry) bool {
		return f.Name() == fileName
	})
}

func Prompt(prompt string) error {
	fmt.Println(prompt)
	fmt.Print("Y/n: ")

	var response string
	_, err := fmt.Scanln(&response)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	if response == "Y" || response == "y" {
		return nil
	}

	return fmt.Errorf("user response: %s", response)
}
