package dataloader

import (
	"log"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

func CloneRepo(url, username, token string) (*git.Repository, error) {
	// check if dir exists -> if not create it
	dir := "./repo"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.Mkdir(dir, 606)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
	}
	// clone the repo
	r, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
		Auth: &http.BasicAuth{
			Username: username,
			Password: token,
		}})

	return r, err
}
