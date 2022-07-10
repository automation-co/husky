package lib

import (
	"errors"
	"os"
)

func Init() error {

	// check if .git exists
	_, err := os.Stat(".git")
	if os.IsNotExist(err) {
		return errors.New("git not initialized")
	}

	// check if .husky exists
	_, err = os.Stat(".husky")

	if err == nil {
		return errors.New(".husky already exist")
	}

	// if not, create .husky
	err = os.Mkdir(".husky", 0755)
	if err != nil {
		return err
	}

	err = os.Mkdir(".husky/hooks", 0755)
	if err != nil {
		return err
	}

	// create default pre-commit hook
	file, err := os.Create(".husky/hooks/pre-commit")
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
