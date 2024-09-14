package lib

import (
	"errors"
	"fmt"
	"os"
	"path"
)

// Add command will create the file from hook value into .hooksy/hooks directory. The cmd appended to shebang string and
// written in the file .hooksy/hooks/<hook>. The function intend to fail if the git hook name is invalid.
func Add(hook string, cmd string) error {
	// validate hooks
	if !isValidHook(hook) {
		return errors.New("invalid hook name")
	}

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

	// check if .hooksy/hooks exists
	_, err := os.Stat(gethooksyHooksDir(true))
	fmt.Println(err)
	if os.IsNotExist(err) {
		fmt.Println("no pre-existing hooks found")

		// create .hooksy/hooks
		err = os.MkdirAll(gethooksyHooksDir(true), 0755)
		if err != nil {
			return err
		}

		fmt.Println("created .hooksy/hooks")
	}

	// create hook
	file, err := os.Create(path.Join(gethooksyHooksDir(true), hook))
	if err != nil {
		return err
	}

	//goland:noinspection GoUnhandledErrorResult
	defer file.Close()

	cmd = "#!/bin/sh\n" + cmd
	_, err = file.WriteString(cmd)
	if err != nil {
		return err
	}

	return nil
}
