package jobrunner

import (
	"context"
	"log"
	"sync"
	"time"
)

// State of jobrunner.
type State struct {
	Conf Config
	//Logger log.Logger

	ctx context.Context
}

// Run the configured jobs.
func (s *State) Run(ctx context.Context) {
	var wg sync.WaitGroup
	s.ctx = ctx
	for i, job := range s.Conf.Jobs {
		wg.Add(1)
		go func() {
			s.startJob(job)
			// TODO: Logger dependency.
			log.Printf("Job #%d finished", i)
			wg.Done()
		}()
	}
	wg.Wait()
}

func (s *State) startJob(job ConfigJob) {
	// TODO: Run CMD and output the response in case of a non-zero exit code.
	// TODO: Logger dependency.
	log.Println(job.Cmd)
	for {
		select {
		case <-s.ctx.Done():
			return
		case <-time.Tick(time.Duration(job.Interval)):
			// TODO: Run CMD and output the response in case of a non-zero exit code.
			// TODO: Logger dependency.
			log.Println(job.Cmd)
		}
	}
}
