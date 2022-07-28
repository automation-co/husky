package lib

import (
	"errors"
	"os"
	"path"
)

// Init command will set up the .husky directory as sibling of .git directory if not exists install pre-commit hook by default
// If .husky exists, it will remove all the files from .git/hooks directory and copy from .husky directory.
func Init() error {
	// check if .git exists
	if isExists, err := gitExists(); err == nil && !isExists {
		return errors.New("git not initialized")
	} else if err != nil {
		return err
	}

	// check if .husky exists
	if isExists, err := huskyExists(); err == nil && isExists {
		return errors.New(".husky already exist")
	} else if err != nil {
		return err
	}

	// if not, create .husky/hooks
	err := os.MkdirAll(getHuskyHooksDir(true), 0755)
	if err != nil {
		return err
	}

	// create default pre-commit hook
	file, err := os.Create(path.Join(getHuskyHooksDir(true), "pre-commit"))
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
