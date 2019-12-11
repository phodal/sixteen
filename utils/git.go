package utils

import (
	"bytes"
	"fmt"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"io/ioutil"
	"os/exec"
	"regexp"
	"strings"
	"syscall"
	"time"
)

func CommitByMessage(message string) {
	r, err := git.PlainOpen(".")
	CheckIfError(err)
	w, err := r.Worktree()
	CheckIfError(err)

	username, err := Username()
	email, err := Email()
	CheckIfError(err)

	commit, err := w.Commit(message, &git.CommitOptions{
		Author: &object.Signature{
			Name:  username,
			Email: email,
			When:  time.Now(),
		},
	})

	CheckIfError(err)

	// Prints the current HEAD to verify that all worked well.
	Info("git show -s")
	obj, err := r.CommitObject(commit)
	CheckIfError(err)

	fmt.Println(obj)
}

type ErrNotFound struct {
	Key string
}

func (e *ErrNotFound) Error() string {
	return fmt.Sprintf("the key `%s` is not found", e.Key)
}

// Entire extracts configuration value from `$HOME/.gitconfig` file ,
// `$GIT_CONFIG`, /etc/gitconfig or include.path files.
func Entire(key string) (string, error) {
	return execGitConfig(key)
}

// Username extracts git user name from `Entire gitconfig`.
// This is same as Entire("user.name")
func Username() (string, error) {
	return Entire("user.name")
}

// Email extracts git user email from `$HOME/.gitconfig` file or `$GIT_CONFIG`.
// This is same as Global("user.email")
func Email() (string, error) {
	return Entire("user.email")
}

func execGitConfig(args ...string) (string, error) {
	gitArgs := append([]string{"config", "--get", "--null"}, args...)
	var stdout bytes.Buffer
	cmd := exec.Command("git", gitArgs...)
	cmd.Stdout = &stdout
	cmd.Stderr = ioutil.Discard

	err := cmd.Run()
	if exitError, ok := err.(*exec.ExitError); ok {
		if waitStatus, ok := exitError.Sys().(syscall.WaitStatus); ok {
			if waitStatus.ExitStatus() == 1 {
				return "", &ErrNotFound{Key: args[len(args)-1]}
			}
		}
		return "", err
	}

	return strings.TrimRight(stdout.String(), "\000"), nil
}

var RepoNameRegexp = regexp.MustCompile(`.+/([^/]+)(\.git)?$`)

func retrieveRepoName(url string) string {
	matched := RepoNameRegexp.FindStringSubmatch(url)
	return strings.TrimSuffix(matched[1], ".git")
}
