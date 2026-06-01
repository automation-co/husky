package lib

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Add creates a hook file in .husky/hooks under the current worktree.
// The cmd argument is prefixed with a #!/bin/sh shebang and written to
// .husky/hooks/<hook>. Fails if the hook name is invalid, the current
// directory is not inside a git working tree, or .husky is not initialized.
func Add(hook string, cmd string) error {
	if !isValidHook(hook) {
		return fmt.Errorf("invalid hook name %s, valid hooks are: %v", hook, strings.Join(validHooks, ", "))
	}

	// check if .git exists and find the root directory of the current git worktree
	worktreeRoot, err := renderWorktreeRoot()
	if err != nil {
		return err
	}

	// check if the .husky directory exists
	huskyDir := renderHuskyDir(worktreeRoot)
	if _, err := os.Stat(huskyDir); os.IsNotExist(err) {
		return fmt.Errorf("the directory %s does not exist, run \"husky init\" and try again", huskyDir)
	} else if err != nil {
		return err
	}

	huskyHooksDir := renderHuskyHooksDir(worktreeRoot)
	if _, err := os.Stat(huskyHooksDir); os.IsNotExist(err) {
		fmt.Printf("no pre-existing hooks found, creating %s now\n", huskyHooksDir)
		if err := os.MkdirAll(huskyHooksDir, 0755); err != nil {
			return err
		}
		fmt.Printf("created %s\n", huskyHooksDir)
	} else if err != nil {
		return err
	}

	file, err := os.Create(filepath.Join(huskyHooksDir, hook))
	if err != nil {
		return err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer file.Close()

	if _, err := file.WriteString("#!/bin/sh\n" + cmd); err != nil {
		return err
	}

	return nil
}
