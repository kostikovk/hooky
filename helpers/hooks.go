package helpers

// HasConfigFolder checks if the current directory is a GoHooks repository.
func IsGoHooksRepository() bool {
	return dirExists(".gohooks")
}

// HasGoHooksGitDirectory checks if the current directory has a .gohooks/git folder.
func HasGoHooksGitDirectory() bool {
	return dirExists(".gohooks/git/hooks")
}
