package dataloader

import (
	"fmt"
	"os"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

func ProcessData() (*DataHolder, error) {
	// Iterate over all playbook files
	start := time.Now()
	files, err := os.ReadDir("./data")
	if err != nil {
		return nil, err
	}
	// IDEA -> Go Concurency ??
	dataHolder := DataHolder{
		Playbooks: map[string]Playbook{},
	}

	for _, file := range files {
		fmt.Println(file.Name())

		fmt.Println("Trying to load file:", file.Name())
		rawContent, err := os.ReadFile(fmt.Sprintf("./data/%s", file.Name()))
		if err != nil {
			return nil, err
		}

		fmt.Println("File Content Loaded")
		var content map[string]interface{}
		err = yaml.Unmarshal(rawContent, &content)
		if err != nil {
			return nil, err
		}
		pb, err := extractPlaybook(content, file.Name())
		if err != nil {
			return nil, err
		}
		fmt.Println(pb)

		// store them in a map with the filename as key
		dataHolder.Playbooks[file.Name()] = *pb
	}

	// create the needed data structure for the visualisation

	// would be a good idea to messure how long it takes to create a big data strcuture
	// if its to long we should pre calc every data strucutre if not we can do it on the fly.
	fin := time.Now()
	fmt.Println("Finished in", fin.Sub(start))
	return &dataHolder, nil
}

func extractPlaybook(content map[string]interface{}, fileName string) (*Playbook, error) {
	fmt.Println(content)
	var playbook Playbook
	playbook.FileName = fileName
	playbook.Id = content["id"].(string)
	fmt.Println("Playbook created, starting to going over the Tasks")
	playbook.Name = content["task"].(map[string]interface{})["name"].(string)

	// Extract the tasks
	for k, v := range content["tasks"].(map[string]interface{}) {
		task := extractTask(k, v.(map[string]interface{}))
		playbook.Tasks = append(playbook.Tasks, task)
	}
	// Extract Inputs
	playbook.Inputs = extractInputs(content["inputs"].([]interface{}))

	// Extract Outputs
	playbook.Outputs = extractOutputs(content["outputs"].([]interface{}))

	return &playbook, nil
}

func extractTask(k string, val map[string]interface{}) Task {
	str := strings.ToLower(val["type"].(string))
	task := val["task"].(map[string]string)
	taskId := val["taskid"].(string)
	nTs := val["nexttasks"].(map[string][]string)

	switch str {
	case "collection":
		return CreateCollectionTask(k, taskId, task["name"], task["description"], nTs)

	case "condition":
		return CreateConditionTask(k, taskId, nTs)

	case "title":
		return CreateTitleTask(k, taskId, nTs)

	case "start":
		return CreateStartTask(k, taskId, nTs)

	case "playbook":
		args := []string{}
		for k := range val["scriptarguments"].(map[string]string) {
			args = append(args, k)
		}
		pbId := task["playbookId"]
		desc := task["description"]
		name := task["name"]
		return CreatePlaybookTask(k, taskId, pbId, desc, name, args, nTs)

	case "builtin":
		args := []string{}
		for k := range val["scriptarguments"].(map[string]string) {
			args = append(args, k)
		}
		return CreateAutomationTask(k, taskId, task["description"], task["name"], args, nTs)

	case "automation":
		args := []string{}
		for k := range val["scriptarguments"].(map[string]string) {
			args = append(args, k)
		}
		return CreateAutomationTask(k, taskId, task["description"], task["name"], args, nTs)
	}

	fmt.Println(fmt.Errorf("PlaybookTaskType: %s not mapped in extract task", str))
	return nil
}

func extractInputs(inputs []interface{}) []PlaybookInput {
	pInputs := []PlaybookInput{}
	for _, v := range inputs {
		val := v.(map[string]interface{})
		pInput := PlaybookInput{
			Name:        val["key"].(string),
			Required:    val["required"].(bool),
			Description: val["description"].(string),
		}
		pInputs = append(pInputs, pInput)
	}
	return pInputs
}

func extractOutputs(outputs []interface{}) []PlaybookOutput {
	pOutputs := []PlaybookOutput{}
	for _, v := range outputs {
		val := v.(map[string]interface{})
		pOutput := PlaybookOutput{
			Name:        val["key"].(string),
			Description: val["description"].(string),
			Type:        val["type"].(string),
		}
		pOutputs = append(pOutputs, pOutput)
	}
	return pOutputs
}
