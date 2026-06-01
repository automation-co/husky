package lib

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

// git hooks currently supported
var validHooks = []string{
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

// contains will return true if str exists in s
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

// isValidHook will call the contains function internally.
func isValidHook(hook string) bool {
	return contains(validHooks, hook)
}

// renderWorktreeRoot returns the absolute path of the top-level directory of the
// current git working tree. In plain vanilla repository clones with only one
// worktree, this is the repository root directory. In clones with multiple linked
// worktrees, this is the linked worktree's own root, not the main repository's working
// tree. renderWorktreeRoot assumes to be called from inside the worktree.
func renderWorktreeRoot() (string, error) {
	out, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		return "", fmt.Errorf("this does not seem to be a git repository, make sure to call this from within the repository - or from within a worktree if you are using worktrees; error: %v", err)
	}
	return strings.TrimSpace(string(out)), nil
}

// renderGitHooksDir returns the absolute path to the hooks directory git will
// actually execute from. Honors linked worktrees (→ shared common dir),
// submodules, and core.hooksPath, because `git rev-parse --git-path hooks`
// does.
func renderGitHooksDir(worktreeRoot string) (string, exec.Cmd, error) {
	cmd := exec.Command("git", "rev-parse", "--git-path", "hooks")
	cmd.Dir = worktreeRoot
	out, err := cmd.Output()
	if err != nil {
		return "", *cmd, err
	}
	p := strings.TrimSpace(string(out))
	if !filepath.IsAbs(p) {
		p = filepath.Join(worktreeRoot, p)
	}
	return p, *cmd, nil
}

// renderHuskyDir returns the absolute path to the .husky directory inside the
// worktree root directory worktreeRoot.
func renderHuskyDir(worktreeRoot string) string {
	return filepath.Join(worktreeRoot, ".husky")
}

// renderHuskyHooksDir returns the absolute path to .husky/hooks inside worktreeRoot.
func renderHuskyHooksDir(worktreeRoot string) string {
	return filepath.Join(renderHuskyDir(worktreeRoot), "hooks")
}
