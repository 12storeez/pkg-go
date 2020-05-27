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
		_ = s.scheduler.AddFunc(worker.Schedule(), func() {
			worker.Run()
		})
	}
	return s
}

func (s *Scheduler) Start() {
	s.scheduler.Start()
	select {}
}
