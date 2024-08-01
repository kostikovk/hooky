package helpers

import (
	"fmt"
	"io"
	"os"
)

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

// copyFile copies a file from pathFrom to pathTo.
// It returns an error if the copy operation fails.
func copyFile(pathFrom, pathTo string) error {
	// Open the source file
	srcFile, err := os.Open(pathFrom)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer srcFile.Close()

	// Create the destination file
	destFile, err := os.Create(pathTo)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}
	defer destFile.Close()

	// Copy the contents of the source file to the destination file
	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return fmt.Errorf("failed to copy file: %w", err)
	}

	// Flush the file writer to ensure all data is written to disk
	err = destFile.Sync()
	if err != nil {
		return fmt.Errorf("failed to flush file writer: %w", err)
	}

	return nil
}
