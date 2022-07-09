package lib

import (
	"fmt"
	"os"
)

func Init() {

	// check if .git exists
	_, err := os.Stat(".git")
	if os.IsNotExist(err) {
		fmt.Println("git not initialized")
		return
	}

	// check if .husky exists
	_, err = os.Stat(".husky")

	if os.IsNotExist(err) {
	} else {
		fmt.Println(".husky already exist.")
		return
	}

	// if not, create .husky
	err = os.Mkdir(".husky", 0755)
	if err != nil {
		panic(err)
	}

	err = os.Mkdir(".husky/hooks", 0755)
	if err != nil {
		panic(err)
	}

	// create default pre-commit hook
	file, err := os.Create(".husky/hooks/pre-commit")
	if err != nil {
		panic(err)
	}

	//goland:noinspection GoUnhandledErrorResult
	defer file.Close()

	_, err = file.WriteString(`#!/bin/sh`)

	if err != nil {
		panic(err)
	}

	// add hooks to .git/hooks
	Install()
}
