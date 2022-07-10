package lib

import (
	"errors"
	"fmt"
	"os"
)

func Add(hook string, cmd string) error {
	// validate hooks
	if !isValidHook(hook) {
		return errors.New("invalid hook name")
	}

	// check if .git exists
	if isExists, err := gitExists(); err == nil && !isExists {
		return errors.New("git not initialized")
	} else if err == nil {
		return err
	}

	// check if .husky exists
	if isExists, err := huskyExists(); err == nil && !isExists {
		return errors.New(".husky not initialized")
	} else if err != nil {
		return err
	}

	// check if .husky/hooks exists
	_, err := os.Stat(".husky/hooks")

	if os.IsNotExist(err) {
		fmt.Println("no pre-existing hooks found")

		// create .husky/hooks
		err = os.Mkdir(".husky/hooks", 0755)
		if err != nil {
			return err
		}

		fmt.Println("created .husky/hooks")
	}

	// create hook
	file, err := os.Create(".husky/hooks/" + hook)
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
