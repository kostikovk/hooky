package helpers

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDirExistsReturnsFalseOnInvalidPath(t *testing.T) {
	if dirExists(string([]byte{0})) {
		t.Fatal("expected false for invalid path")
	}
}

func TestDirExistsDistinguishesDirectoriesAndFiles(t *testing.T) {
	root := t.TempDir()
	file := filepath.Join(root, "file.txt")
	if err := os.WriteFile(file, []byte("hooky"), 0o600); err != nil {
		t.Fatalf("write temp file: %v", err)
	}

	if !dirExists(root) {
		t.Fatal("expected directory path to exist")
	}
	if dirExists(file) {
		t.Fatal("expected file path to return false")
	}
}
