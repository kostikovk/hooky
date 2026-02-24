package helpers

import (
	"fmt"
	"os"
)

func dirExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}

	return info.IsDir()
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
