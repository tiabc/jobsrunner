package jobsrunner

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"sync"
	"time"
)

// Runtime of jobrunner.
type Runtime struct {
	// TODO: Logger dependency.
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
			log.Printf("Job #%d finished", i)
			wg.Done()
		}()
	}
	wg.Wait()
}

func (r *Runtime) startJob(job ConfigJob) {
	r.runCmd(job.Cmd)
	for {
		select {
		case <-r.ctx.Done():
			return
		case <-time.Tick(time.Duration(job.Interval)):
			r.runCmd(job.Cmd)
		}
	}
}

func (r *Runtime) runCmd(cmd string) {
	c := exec.Command("sh", "-c", cmd)

	stderr, err := c.StderrPipe()
	if err != nil {
		log.Printf(`"%v" didn'r run: failed to pipe stderr: %s`, cmd, err)
		return
	}

	stdout, err := c.StdoutPipe()
	if err != nil {
		log.Printf(`"%v" didn'r run: failed to pipe stdout: %s`, cmd, err)
		return
	}

	if err := c.Start(); err != nil {
		log.Printf(`"%v" didn'r run: failed to start: %s`, cmd, err)
		return
	}
	b := time.Now()
	var (
		slurp, _  = ioutil.ReadAll(stderr)
		output, _ = ioutil.ReadAll(stdout)
	)
	if err := c.Wait(); err != nil {
		if errExit, ok := err.(*exec.ExitError); ok {
			o := fmt.Sprintf(`"%v" finished with %s`, cmd, errExit)
			if len(slurp) != 0 {
				o = fmt.Sprintf("%s, stderr: %s", o, slurp)
			}
			if len(output) != 0 {
				o = fmt.Sprintf("%s, stdout: %s", o, output)
			}
			log.Println(o)
		} else {
			log.Printf(`"%v" finished with error %s`, cmd, err)
		}
	}
	log.Printf(`"%v" took %s`, cmd, time.Now().Sub(b))
}
