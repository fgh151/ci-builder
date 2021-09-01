package main

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"os"
)

//goland:noinspection GoUnusedFunction
func clone(toDir string, gitUrl string, ref string) *git.Repository {

	auth, err := ssh.NewPublicKeysFromFile("git", os.Getenv("GIT_SSH_PRIVATE_KEY"), "")
	checkErr(err)

	r, err := git.PlainClone(toDir, false, &git.CloneOptions{
		URL:           gitUrl,
		Progress:      os.Stdout,
		Auth:          auth,
		ReferenceName: plumbing.NewBranchReferenceName(ref),
	})

	checkErr(err)

	return r
}
