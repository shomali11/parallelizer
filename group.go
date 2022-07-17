package parallelizer

import (
	"errors"
	"sync"
)

const (
	nilFunctionError = "nil function"
)

// NewGroup create a new group of workers
func NewGroup(options ...GroupOption) *Group {
	groupOptions := newGroupOptions(options...)

	group := &Group{
		jobsChannel:  make(chan func() error, groupOptions.JobQueueSize),
		errorChannel: make(chan error),
		waitGroup:    &sync.WaitGroup{},
	}

	for i := 1; i <= groupOptions.PoolSize; i++ {
		go group.worker()
	}
	return group
}

// Group a group of workers executing functions concurrently
type Group struct {
	jobsChannel  chan func() error
	errorChannel chan error
	waitGroup    *sync.WaitGroup
}

// Add adds function to queue of jobs to execute
func (g *Group) Add(function func() error) error {
	if function == nil {
		return errors.New(nilFunctionError)
	}

	g.waitGroup.Add(1)
	g.jobsChannel <- function
	return nil
}

// Wait waits until workers finished the jobs in the queue
func (g *Group) Wait(options ...WaitOption) error {
	waitOptions := newWaitOptions(options...)

	waitChannel := make(chan bool)
	go func() {
		g.waitGroup.Wait()
		close(waitChannel)
	}()

	select {
	case <-waitOptions.Context.Done():
		return waitOptions.Context.Err()
	case err := <-g.errorChannel:
		return err
	case <-waitChannel:
		return nil
	}
}

// Close closes resources
func (g *Group) Close() {
	close(g.jobsChannel)
	close(g.errorChannel)
}

func (g *Group) worker() {
	for job := range g.jobsChannel {
		err := job()
		if err != nil {
			g.errorChannel <- err
			return
		}
		g.waitGroup.Done()
	}
}
