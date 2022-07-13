package lib

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func Install() error {
	fmt.Println("Installing hooks")

	// check if .git exists
	if isExists, err := gitExists(); err == nil && !isExists {
		return errors.New("git not initialized")
	} else if err != nil {
		return err
	}

	// check if .husky exists
	if isExists, err := huskyExists(); err == nil && !isExists {
		return errors.New(".husky not initialized")
	} else if err != nil {
		return err
	}

	gitHooksDir, huskyHooksDir := getGitHooksDir(true), getGitHooksDir(true)
	// check if .husky/hooks exists
	_, err := os.Stat(huskyHooksDir)
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

	// copy all files in .husky/hooks to .git/hooks
	var hooks []string
	err = filepath.Walk(huskyHooksDir,
		func(path string, info os.FileInfo, err error) error {
			hooks = append(hooks, path)
			return nil
		})
	if err != nil {
		return err
	}
	for _, hook := range hooks {

		// skip .husky/hooks
		if hook == huskyHooksDir {
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
