package utils

import (
	"fmt"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"time"
)

func CommitByMessage(message string) {
	r, err := git.PlainOpen(".")
	CheckIfError(err)
	w, err := r.Worktree()
	CheckIfError(err)

	commit, err := w.Commit(message, &git.CommitOptions{
		Author: &object.Signature{
			When: time.Now(),
		},
	})

	CheckIfError(err)

	// Prints the current HEAD to verify that all worked well.
	Info("git show -s")
	obj, err := r.CommitObject(commit)
	CheckIfError(err)

	fmt.Println(obj)
}

