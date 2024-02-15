package dataloader

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestExtractPlaybook(t *testing.T) {
	testcases := []struct {
		FileName string
		error    bool
	}{
		{"../../data/playbook-TDR_-_Unclassified_Incident.yml", false},
	}
	for _, tC := range testcases {
		t.Run(tC.FileName, func(t *testing.T) {
			rawContent, err := os.ReadFile(tC.FileName)
			if err != nil {
				assert.Error(t, err)
			}

			// fmt.Println("File Content Loaded")
			var content map[string]interface{}
			err = yaml.Unmarshal(rawContent, &content)
			if err != nil {
				assert.Error(t, err)
			}
			_, err = extractPlaybook(content, tC.FileName)

			// TODO playbook Validation should be done, needing a simpler example
			// fmt.Println(pb)
			assert.Equal(t, tC.error, err != nil)
		})
	}
}
