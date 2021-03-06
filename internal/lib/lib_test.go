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

// TestInitNoGit validates that the lib.Init() will return error when git is not initialized
func TestInitNoGit(t *testing.T) {
	currentDir, _ := os.Getwd()
	repoPath, err := getRepoPath()
	defer os.RemoveAll(repoPath)
	defer os.Chdir(currentDir)

	if err != nil {
		t.Error(err)
	} else if err := os.Chdir(repoPath); err != nil {
		t.Error(err)
	}

	err = huskyLib.Init()
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
func TestInitHuskyExists(t *testing.T) {
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

	if err := huskyLib.Init(); err == nil {
		t.Error(nilErrorMsg)
	} else if err.Error() != ".husky already exist" {
		t.Error(err)
	}
}

// TestInit validates if the lib.Init() function runs accurately or not.
// It will skip the testing of the lib.Install() function in the end.
func TestInit(t *testing.T) {
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

	if err := huskyLib.Init(); err != nil {
		t.Error(err)
	}

	if stat, err := os.Stat(path.Join(repoPath, ".husky", "hooks")); err != nil && !os.IsExist(err) {
		t.Error(err)
	} else if !stat.IsDir() {
		t.Error(expectedDirErrorMsg)
	}

	if stat, err := os.Stat(path.Join(repoPath, ".husky", "hooks", "pre-commit")); err != nil && !os.IsExist(err) {
		t.Error(err)
	} else if stat.IsDir() {
		t.Error(expectedFileErrorMsg)
	}

	if content, err := ioutil.ReadFile(path.Join(repoPath, ".husky", "hooks", "pre-commit")); err != nil {
		t.Error(err)
	} else if "#!/bin/sh" != string(content) {
		t.Error(invalidContentsErrorMsg)
	}
}

// TestInstallNoGit validates that the lib.Install() will return error when git is not initialized
func TestInstallNoGit(t *testing.T) {
	currentDir, _ := os.Getwd()
	repoPath, err := getRepoPath()
	defer os.RemoveAll(repoPath)
	defer os.Chdir(currentDir)

	if err != nil {
		t.Error(err)
	} else if err := os.Chdir(repoPath); err != nil {
		t.Error(err)
	}

	err = huskyLib.Install()
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

// TestInstallNoHusky validates if the lib.Install() returns error if husky is not initialized
func TestInstallNoHusky(t *testing.T) {
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

	if err := huskyLib.Install(); err == nil {
		t.Error(nilErrorMsg)
	} else if err.Error() != ".husky not initialized" {
		t.Error(err)
	}
}

// TestInstallNoHuskyHooks validates if lib.Install() returns error when .husky/hooks directory doesn't exists
func TestInstallNoHuskyHooks(t *testing.T) {
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

	if err := huskyLib.Install(); err == nil {
		t.Error(nilErrorMsg)
	} else if err.Error() != "no hooks found" {
		t.Error(err)
	}
}

// TestInstall validates everything lib.Install() has done is correct
func TestInstall(t *testing.T) {
	currentDir, _ := os.Getwd()
	repoPath, err := getRepoPath()
	//defer os.RemoveAll(repoPath)
	defer os.Chdir(currentDir)
	if err != nil {
		t.Error(err)
	}
	setupRepo(repoPath)

	if err := os.MkdirAll(path.Join(repoPath, ".husky", "hooks"), 0755); err != nil {
		t.Error(err)
	}

	gibbrish := randomString(32)
	var file *os.File
	if file, err = os.Create(path.Join(repoPath, ".husky", "hooks", "pre-commit")); err != nil {
		t.Error(err)
	} else if _, err := file.WriteString(gibbrish); err != nil {
		t.Error(err)
	}
	file.Close()

	if err := os.Chdir(repoPath); err != nil {
		t.Error(err)
	}

	if err := huskyLib.Install(); err != nil {
		t.Error(err)
	}

	if file, err := os.Stat(path.Join(".git", "hooks", "pre-commit")); err != nil && !os.IsExist(err) {
		t.Error(err)
	} else if file.IsDir() {
		t.Error(expectedDirErrorMsg)
	} else if content, err := ioutil.ReadFile(path.Join(".git", "hooks", "pre-commit")); err != nil {
		t.Error(err)
	} else if string(content) != gibbrish {
		t.Error(invalidContentsErrorMsg)
	}
}
