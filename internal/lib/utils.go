package lib

import (
	"os"
	"path"
)

// git hooks currently supported
var validHooks = []string{
	"applypatch-msg",
	"commit-msg",
	"fsmonitor-watchman",
	"post-checkout",
	"post-update",
	"pre-applypatch",
	"pre-commit",
	"pre-push",
	"pre-rebase",
	"prepare-commit-msg",
	"update",
	"pre-receive",
	"pre-merge-commit",
	"push-to-checkout",
}

// contains will return true if str exists in s
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

// isValidHook will call the contains function internally.
func isValidHook(hook string) bool {
	return contains(validHooks, hook)
}

// gitExists will return true if the comman is executed under .git directory
// TODO: support recursive find .git directory till home
func gitExists() (bool, error) {
	// check if .git exists
	_, err := os.Stat(".git")
	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

// hooksyExists will return true if exists, otherwise false
// TODO: support recursive find .hooksy directory till home
func hooksyExists() (bool, error) {
	_, err := os.Stat(".hooksy")

	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

// gethooksyHooksDir will return the relative or absolute .hooksy hooks directory
func gethooksyHooksDir(relative bool) string {
	if relative {
		return path.Join(".hooksy", "hooks")
	}

	cwd, _ := os.Getwd()
	return path.Join(cwd, ".hooksy", "hooks")
}

// getGitHooksDir will return the relative or absolute .git hooks directory
func getGitHooksDir(relative bool) string {
	if relative {
		return path.Join(".git", "hooks")
	}

	cwd, _ := os.Getwd()
	return path.Join(cwd, ".git", "hooks")
}
