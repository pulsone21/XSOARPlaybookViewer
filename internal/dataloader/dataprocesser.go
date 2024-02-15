package dataloader

import (
	"fmt"
	"os"
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
	dataHolder := DataHolder{}

	for _, file := range files {
		// fmt.Println(file.Name())

		// fmt.Println("Trying to load file:", file.Name())
		rawContent, err := os.ReadFile(file.Name())
		if err != nil {
			return nil, err
		}

		// fmt.Println("File Content Loaded")
		var content map[string]interface{}
		err = yaml.Unmarshal(rawContent, &content)
		if err != nil {
			return nil, err
		}
		pb, err := extractPlaybook(content, file.Name())
		if err != nil {
			return nil, err
		}
		// fmt.Println(pb)

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
	//    fmt.Println(content)
	var playbook Playbook
	playbook.Name = content["name"].(string)
	playbook.FileName = fileName
	playbook.Id = content["id"].(string)
	// fmt.Println("Playbook created, starting to going over the Tasks")

	// Extract the tasks
	for k, v := range content["tasks"].(map[string]interface{}) {
		task := extractTask(k, v.(map[string]interface{}))
		// fmt.Println("NextTasks assignt")
		playbook.NextTasks = append(playbook.NextTasks, *task)
	}
	// Extract Inputs
	playbook.Inputs = extractInputs(content["inputs"].([]interface{}))

	// Extract Outputs
	playbook.Outputs = extractOutputs(content["outputs"].([]interface{}))

	return &playbook, nil
}

func extractNextTask(nextTasks map[string]interface{}) map[string][]string {
	nT := make(map[string][]string)
	for k, v := range nextTasks {
		strings := []string{}
		for _, item := range v.([]interface{}) {
			strings = append(strings, item.(string))
		}
		nT[k] = strings
	}
	return nT
}

func extractTask(k string, val map[string]interface{}) *PlaybookTask {
	// fmt.Println("Converted val")
	task := PlaybookTask{
		Id:   k,
		Type: toTaskType(val["type"].(string)),
	}

	// fmt.Println("Created Task with base infos")
	task.Task = val["task"].(map[string]interface{})

	// fmt.Println("Created Task with task infos")
	if val["nexttasks"] != nil {
		task.NextTasks = extractNextTask(val["nexttasks"].(map[string]interface{}))
	}

	return &task
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
