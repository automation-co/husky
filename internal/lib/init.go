package lib

import (
	"fmt"
	"os"
	"path/filepath"
)

// Init sets up the .husky/hooks directory under the current worktree if it
// does not already exist, creates a default pre-commit hook, and then calls
// Install() to link the hooks into git's hooks directory.
func Init() error {
	worktreeRoot, err := renderWorktreeRoot()
	if err != nil {
		return err
	}

	// check if .husky exists
	huskyDir := renderHuskyDir(worktreeRoot)
	if _, err = os.Stat(huskyDir); err == nil {
		return fmt.Errorf("%s already exists", huskyDir)
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("error when checking if %s already exists: %v", huskyDir, err)
	}

	// create .husky/hooks
	huskyHooks := renderHuskyHooksDir(worktreeRoot)
	fmt.Printf("creating the %s now\n", huskyHooks)
	if err := os.MkdirAll(huskyHooks, 0755); err != nil {
		return fmt.Errorf("error when creating %s: %v", huskyHooks, err)
	}

	preCommitHook := filepath.Join(huskyHooks, "pre-commit")
	file, err := os.Create(preCommitHook)
	if err != nil {
		return fmt.Errorf("error when creating %s: %v", preCommitHook, err)
	}
	//goland:noinspection GoUnhandledErrorResult
	defer file.Close()

	if _, err := file.WriteString("#!/bin/sh\n"); err != nil {
		return fmt.Errorf("error when writing to %s: %v", preCommitHook, err)
	}

	return Install()
}
