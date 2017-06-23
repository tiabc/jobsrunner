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

	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup
}

// Run the configured jobs.
func (s *State) Run() {
	s.wg = sync.WaitGroup{}
	s.ctx, s.cancel = context.WithCancel(context.Background())
	for i, job := range s.Conf.Jobs {
		s.wg.Add(1)
		go func() {
			s.startJob(job)
			// TODO: Logger dependency.
			log.Printf("Job #%d finished", i)
			s.wg.Done()
		}()
	}
}

// Stop gracefully terminates the execution and waits until all jobs finished.
func (s *State) Stop() {
	// TODO: Stop gracefully.
	log.Printf("Finishing %d jobs...\n", len(s.Conf.Jobs))
	s.cancel()
	s.wg.Wait()
}

func (s *State) startJob(job ConfigJob) {
	for {
		select {
		case <-s.ctx.Done():
			return
		default:
		}
		// TODO: Run CMD and output the response in case of a non-zero exit code.
		log.Println(job.Cmd)
		select {
		case <-s.ctx.Done():
			return
		default:
		}
		// TODO: Replace sleep with some tick? Otherwise, one will have to wait till the end
		// of this long interval before the runner finishes.
		time.Sleep(time.Duration(job.Interval))
	}
}
