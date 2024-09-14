package lib

import (
	"errors"
	"os"
	"path"
)

// Init command will set up the .hooksy directory as sibling of .git directory if not exists install pre-commit hook by default
// If .hooksy exists, it will remove all the files from .git/hooks directory and copy from .hooksy directory.
func Init() error {
	// check if .git exists
	if isExists, err := gitExists(); err == nil && !isExists {
		return errors.New("git not initialized")
	} else if err != nil {
		return err
	}

	// check if .hooksy exists
	if isExists, err := hooksyExists(); err == nil && isExists {
		return errors.New(".hooksy already exist")
	} else if err != nil {
		return err
	}

	// if not, create .hooksy/hooks
	err := os.MkdirAll(gethooksyHooksDir(true), 0755)
	if err != nil {
		return err
	}

	// create default pre-commit hook
	file, err := os.Create(path.Join(gethooksyHooksDir(true), "pre-commit"))
	if err != nil {
		return err
	}

	//goland:noinspection GoUnhandledErrorResult
	defer file.Close()

	_, err = file.WriteString(`#!/bin/sh`)
	if err != nil {
		return err
	}

	// add hooks to .git/hooks
	err = Install()
	if err != nil {
		return err
	}

	return nil
}
