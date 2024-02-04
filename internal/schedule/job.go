package schedule

import (
	"fmt"
	"sync"
)

type Job struct {
	Name  string
	Cron  string
	Func  func() error
	State JobState
}

type JobFunc func(chan<- string, *sync.WaitGroup)

type JobState string

const (
	Waiting  JobState = "waiting"
	Running  JobState = "running"
	Finished JobState = "finished"
)

func (j *Job) Execute() error {
	if j.State == Finished {
		return fmt.Errorf("Job %s is already finished", j.Name)
	}
	j.State = Running
	err := j.Func()
	j.State = Finished
	return err
}

func NewJob(name, cron string, f func() error) *Job {
	return &Job{
		Name:  name,
		Cron:  cron,
		Func:  f,
		State: Waiting,
	}
}
