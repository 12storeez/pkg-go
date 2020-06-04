package scheduler

import (
	"github.com/robfig/cron"
)

type Scheduler struct {
	scheduler *cron.Cron
}

type Worker interface {
	Run()
	Schedule() string
}

func New(workers ...Worker) *Scheduler {
	s := &Scheduler{scheduler: cron.New()}
	for _, worker := range workers {
		temp := worker
		_ = s.scheduler.AddFunc(temp.Schedule(), func() {
			temp.Run()
		})
	}
	return s
}

func (s *Scheduler) Start() {
	s.scheduler.Start()
	select {}
}
