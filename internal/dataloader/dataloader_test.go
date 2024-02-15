package dataloader

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSameRepo(t *testing.T) {
	testcases := []struct {
		url     string
		outcome bool
	}{
		{"github.com", false},
		{"http://github.com/pulsone21/SentinelDeployTest.git", true},
	}

	r, err := openRepo("../../repo")
	if err != nil {
		log.Fatal(err)
	}

	// running the test cases
	for _, tC := range testcases {
		res := isSameRepo(r, tC.url)
		assert.Equal(t, tC.outcome, res)
	}
}
