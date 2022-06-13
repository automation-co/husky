package lib

import (
	"fmt"
	"os"
)

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func Add(hook string, cmd string) {

	// check if hook name is valid
	validHooks := []string{
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
	if !contains(validHooks, hook) {
		fmt.Println("Invalid hook name.")
		return
	}

	// check if .git exists
	_, err := os.Stat(".git")
	if os.IsNotExist(err) {
		fmt.Println("git not initialized")
		return
	}

	// check if .husky exists
	_, err = os.Stat(".husky")

	if os.IsNotExist(err) {
		fmt.Println(".husky not initialized.")
		return
	}

	// check if .husky/hooks exists
	_, err = os.Stat(".husky/hooks")

	if os.IsNotExist(err) {
		fmt.Println("no pre-existing hooks found")

		// create .husky/hooks
		err = os.Mkdir(".husky/hooks", 0755)
		if err != nil {
			panic(err)
		}

		fmt.Println("created .husky/hooks")
	}

	// create hook
	file, err := os.Create(".husky/hooks/" + hook)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	cmd = "#!/bin/sh\n" + cmd
	_, err = file.WriteString(cmd)
	if err != nil {
		panic(err)
	}

}
