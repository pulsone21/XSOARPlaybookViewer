package dataloader

import (
	"fmt"
	"log"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

func CloneRepo(url, username, token string) (*git.Repository, error) {
	// check if dir exists -> if not create it
	dir := "./repo"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.Mkdir(dir, 0666)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
	}
	fmt.Println(url)
	// clone the repo
	r, err := git.PlainClone(dir, false, &git.CloneOptions{
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		URL:               url,
		Progress:          os.Stdout,
		Auth: &http.BasicAuth{
			Username: "NA",
			Password: token,
		},
	})

	// TODO -> we should check if the repo is the same
	if err != nil && err.Error() == "repository already exists" {
		r, err = openRepo(dir)
		if err != nil {
			fmt.Println("Error in open Repo")
			fmt.Println(err.Error())
			return nil, err
		}

		if !isSameRepo(r, url) {
			os.RemoveAll("./repo")
			r, err = CloneRepo(url, username, token)
			if err != nil {
				return nil, fmt.Errorf("there was already an Repo exisiting, but not the same as configured, tried to remove everything and clone it again, but didn't worked out.\n %v", err.Error())
			}
		}
		err = PullRepo(r, os.Getenv("GIT_REPO_USERNAME"), os.Getenv("GIT_REPO_TOKEN"))
	}

	return r, err
}

func PullRepo(repo *git.Repository, username, token string) error {
	w, err := repo.Worktree()
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = w.Pull(&git.PullOptions{
		RemoteName: "origin",
		Auth: &http.BasicAuth{
			Username: username,
			Password: token,
		},
	})

	if err != nil && err != git.NoErrAlreadyUpToDate {
		fmt.Println("Error in Pull")
		log.Fatal(err)
	}

	return nil
}

func openRepo(dir string) (*git.Repository, error) {
	return git.PlainOpen(dir)
}

func isSameRepo(repo *git.Repository, url string) bool {
	remotes, _ := repo.Remotes()
	containsURL := false
	for _, remote := range remotes {
		if containsURL {
			break
		}
		for _, rurl := range remote.Config().URLs {
			if rurl == url {
				containsURL = true
				break
			}
		}
	}
	return containsURL
}
