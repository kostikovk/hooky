package helpers

import "os"

var AbsoluteGoHooksPath = getAbsolutePath(".gohooks")
var AbsoluteGoHooksGitHooksPath = getAbsolutePath(".gohooks/git/hooks")

// HasConfigFolder checks if the current directory is a GoHooks repository.
func IsGoHooksRepository() bool {
	return dirExists(AbsoluteGoHooksPath)
}

// HasGoHooksGitDirectory checks if the current directory has a .gohooks/git folder.
func HasGoHooksGitDirectory() bool {
	return dirExists(AbsoluteGoHooksGitHooksPath)
}

// CreateGoHooksGitDirectory creates a .gohooks/git/hooks folder.
func CreateGoHooksGitDirectory() error {
	return os.MkdirAll(AbsoluteGoHooksGitHooksPath, 0755)
}

// Delete .gohooks directory
func DeleteGoHooksDirectory() error {
	return os.RemoveAll(AbsoluteGoHooksPath)
}
