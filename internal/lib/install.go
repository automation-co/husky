package lib

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// Install iterates over the files in .husky/hooks, creates hardlinks into git's
// hooks directory, and chmods them to 0755. Fails if the current directory
// is not inside a git working tree or if .husky is not initialized.
func Install() error {
	fmt.Println("Installing hooks")

	worktreeRoot, err := renderWorktreeRoot()
	if err != nil {
		return err
	}

	huskyDir := renderHuskyDir(worktreeRoot)
	if _, err := os.Stat(huskyDir); os.IsNotExist(err) {
		return errors.New(".husky not initialized")
	} else if err != nil {
		return fmt.Errorf("error when checking if %s already exists: %v", huskyDir, err)
	}

	huskyHooks := renderHuskyHooksDir(worktreeRoot)
	if _, err := os.Stat(huskyHooks); os.IsNotExist(err) {
		return fmt.Errorf("no hooks found in %s, use \"husky add\" to create some hooks first", huskyHooks)
	} else if err != nil {
		return fmt.Errorf("error when checking if %s already exists: %v", huskyHooks, err)
	}

	gitHooks, cmd, err := renderGitHooksDir(worktreeRoot)
	if err != nil {
		return fmt.Errorf("error when executing \"%v\" to find .git/hooks: %v", cmd, err)
	}

	fmt.Printf("deleting %s now\n", gitHooks)
	if err := os.RemoveAll(gitHooks); err != nil {
		return err
	}
	fmt.Printf("re-creating %s\n", gitHooks)
	if err := os.MkdirAll(gitHooks, 0755); err != nil {
		return err
	}

	entries, err := os.ReadDir(huskyHooks)
	if err != nil {
		return fmt.Errorf("error when listing the directory %s: %v", huskyHooks, err)
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		src := filepath.Join(huskyHooks, entry.Name())
		dst := filepath.Join(gitHooks, entry.Name())
		if err := os.Link(src, dst); err != nil {
			return fmt.Errorf("error when hard-linking from %s to %s: %v", src, dst, err)
		}
		if err := os.Chmod(dst, 0755); err != nil {
			return fmt.Errorf("error making %s executable: %v", dst, err)
		}
		fmt.Printf("hard-link for %s -> %s has been created\n", src, dst)
	}

	fmt.Println("All hooks installed successfully")
	return nil
}
