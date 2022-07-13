package lib_test

import (
	huskyLib "github.com/automation-co/husky/internal/lib"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"path"
	"strings"
	"testing"
	"time"
)

// global error message declration block
const (
	nilErrorMsg             = "method has return nil error"
	invalidContentsErrorMsg = "invalid file contents"
	expectedDirErrorMsg     = "expected the path to be a directory"
	expectedFileErrorMsg    = "expected the path to be a file"
)

// getRepoPath is a testing utility function to create a random directory and return its path
func getRepoPath() (string, error) {
	name, err := ioutil.TempDir(os.TempDir(), "husky")
	if err != nil {
		return "", err
	}

	return name, nil
}

// setupRepo creates a git repository in the path
func setupRepo(path string) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return
	}

	p := exec.Command("git", "init", path)
	p.Start()
	p.Wait()
}

// randomString returns a random string of a specific length
func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())

	alphabet := "abcdefghijklmnopqrstuvwxyz"
	var sb strings.Builder

	l := len(alphabet)

	for i := 0; i < length; i++ {
		c := alphabet[rand.Intn(l)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// TestAddInvalidHook validates that the lib.Add() will return error when invalid hook name is provided
func TestAddInvalidHook(t *testing.T) {
	err := huskyLib.Add(randomString(13), randomString(20))

	if err == nil {
		t.Error(nilErrorMsg)
	} else if err.Error() != "invalid hook name" {
		t.Error(err)
	}
}

// TestAddNoGit validates that the lib.Add() will return error when git is not initialized
func TestAddNoGit(t *testing.T) {
	currentDir, _ := os.Getwd()
	repoPath, err := getRepoPath()
	defer os.RemoveAll(repoPath)
	defer os.Chdir(currentDir)

	if err != nil {
		t.Error(err)
	} else if err := os.Chdir(repoPath); err != nil {
		t.Error(err)
	}

	err = huskyLib.Add("pre-commit", "whoami")
	if err == nil {
		t.Error(nilErrorMsg)
	} else if err.Error() != "git not initialized" {
		t.Error(err)
	}

	if err := os.Chdir(currentDir); err != nil {
		t.Error(err)
	}
	if err := os.RemoveAll(repoPath); err != nil {
		t.Error(err)
	}
}

// TestAddNoHusky validates if the lib.Add() returns error if husky is not initialized
func TestAddNoHusky(t *testing.T) {
	currentDir, _ := os.Getwd()
	repoPath, err := getRepoPath()
	defer os.RemoveAll(repoPath)
	defer os.Chdir(currentDir)

	if err != nil {
		t.Error(err)
	}
	setupRepo(repoPath)

	if err := os.Chdir(repoPath); err != nil {
		t.Error(err)
	}

	if err := huskyLib.Add("pre-commit", "whoami"); err == nil {
		t.Error(nilErrorMsg)
	} else if err.Error() != ".husky not initialized" {
		t.Error(err)
	}
}

// TestAddNoHusky validates if the lib.Add() returns error if husky is not initialized
func TestAdd(t *testing.T) {
	currentDir, _ := os.Getwd()
	repoPath, err := getRepoPath()
	defer os.RemoveAll(repoPath)
	defer os.Chdir(currentDir)

	if err != nil {
		t.Error(err)
	}
	setupRepo(repoPath)

	if err := os.Mkdir(path.Join(repoPath, ".husky"), 0755); err != nil {
		t.Error(err)
	}

	if err := os.Chdir(repoPath); err != nil {
		t.Error(err)
	}

	if err := huskyLib.Add("pre-commit", "whoami"); err != nil {
		t.Error(err)
	}

	preCommit := path.Join(repoPath, ".husky", "hooks", "pre-commit")
	if _, err := os.Stat(preCommit); os.IsNotExist(err) {
		t.Error(err)
	}

	if content, err := ioutil.ReadFile(preCommit); err != nil {
		t.Error(err)
	} else if "#!/bin/sh\nwhoami" != string(content) {
		t.Error(invalidContentsErrorMsg)
	}
}
