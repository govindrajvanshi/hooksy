package lib

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// Install command will iterate the files from .hooksy/hooks create the hardlinks into .git/hooks and chmod to 0755 permission.
// The function intend to fail in case .git or .hooksy directory doesn't exist.
func Install() error {
	fmt.Println("Installing hooks")

	// check if .git exists
	if isExists, err := gitExists(); err == nil && !isExists {
		return errors.New("git not initialized")
	} else if err != nil {
		return err
	}

	// check if .hooksy exists
	if isExists, err := hooksyExists(); err == nil && !isExists {
		return errors.New(".hooksy not initialized")
	} else if err != nil {
		return err
	}

	gitHooksDir, hooksyHooksDir := getGitHooksDir(true), gethooksyHooksDir(true)
	// check if .hooksy/hooks exists
	_, err := os.Stat(hooksyHooksDir)
	if os.IsNotExist(err) {
		return errors.New("no hooks found")
	}

	// delete all files in .git/hooks
	if err := os.RemoveAll(gitHooksDir); err != nil {
		return err
	}

	// create .git/hooks
	if err := os.Mkdir(gitHooksDir, 0755); err != nil {
		return err
	}

	// copy all files in .hooksy/hooks to .git/hooks
	var hooks []string
	err = filepath.Walk(hooksyHooksDir,
		func(path string, info os.FileInfo, err error) error {
			hooks = append(hooks, path)
			return nil
		})
	if err != nil {
		return err
	}
	for _, hook := range hooks {

		// skip .hooksy/hooks
		if hook == hooksyHooksDir {
			continue
		}

		fmt.Println(hook)

		// copy file to .git/hooks
		err = os.Link(hook, filepath.Join(gitHooksDir, filepath.Base(hook)))
		if err != nil {
			return err
		}

		// make file executable
		err = os.Chmod(filepath.Join(gitHooksDir, filepath.Base(hook)), 0755)
		if err != nil {
			return err
		}

	}
	fmt.Println("Hooks installed")

	return nil
}
