package jobsrunner

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

// Runtime of jobrunner.
type Runtime struct {
	Conf Config
	//Logger log.Logger

	ctx context.Context
}

// NewFromFile creates Runtime with Config parsed from the specified file.
func NewFromFile(filename string) (Runtime, error) {
	conf, err := NewConfigFromFile(filename)
	if err != nil {
		return Runtime{}, fmt.Errorf("can't create config: %s", err)
	}
	return Runtime{
		Conf: conf,
	}, nil
}

// Run the configured jobs.
func (r *Runtime) Run(ctx context.Context) {
	var wg sync.WaitGroup
	r.ctx = ctx
	for i, job := range r.Conf.Jobs {
		wg.Add(1)
		go func() {
			r.startJob(job)
			// TODO: Logger dependency.
			log.Printf("Job #%d finished", i)
			wg.Done()
		}()
	}
	wg.Wait()
}

func (r *Runtime) startJob(job ConfigJob) {
	// TODO: Run CMD and output the response in case of a non-zero exit code.
	// TODO: Logger dependency.
	log.Println(job.Cmd)
	for {
		select {
		case <-r.ctx.Done():
			return
		case <-time.Tick(time.Duration(job.Interval)):
			// TODO: Run CMD and output the response in case of a non-zero exit code.
			// TODO: Logger dependency.
			log.Println(job.Cmd)
		}
	}
}
