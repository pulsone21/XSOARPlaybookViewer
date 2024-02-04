package schedule

import "fmt"

type Scheduler struct {
	Jobs    []Job
	CurrJob *Job
	Running bool
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		Jobs: []Job{},
	}
}

func (s *Scheduler) AddJob(job Job) {
	s.Jobs = append(s.Jobs, job)
	if !s.Running {
		s.CurrJob = &job
		s.Start()
	}
}

func (s *Scheduler) Start() {
	s.Running = true
	go func() {
		for s.Running {
			if s.CurrJob == nil {
				s.Running = false
				break
			}
			err := s.CurrJob.Execute()
			if err != nil {
				fmt.Println(fmt.Errorf("Error executing job %s: %s", s.CurrJob.Name, err.Error()))
			}
			s.setNextJob()
		}
	}()
}

func (s *Scheduler) setNextJob() {
	if len(s.Jobs) == 0 {
		s.Running = false
		return
	}
	s.CurrJob = &s.Jobs[0]
	s.Jobs = s.Jobs[1:]
}
