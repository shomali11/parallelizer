package parallelizer

import (
	"errors"
	"sync"
	"time"
)

const (
	timeoutError     = "timeout"
	nilFunctionError = "nil function"
)

// Runner allows you to parallelize function calls with an optional timeout
type Runner struct {
	Timeout   time.Duration
	functions []func()
}

// Add adds function to list of functions to parallelize
func (p *Runner) Add(function func()) error {
	if function == nil {
		return errors.New(nilFunctionError)
	}

	p.functions = append(p.functions, function)
	return nil
}

// Run parallelizes the function calls
func (p *Runner) Run(functions ...func()) error {
	err := p.appendFunctions(functions...)
	if err != nil {
		return errors.New(nilFunctionError)
	}

	var waitGroup sync.WaitGroup
	waitGroup.Add(len(p.functions))

	for _, function := range p.functions {
		go func(copy func()) {
			defer waitGroup.Done()
			copy()
		}(function)
	}

	return p.wait(&waitGroup)
}

func (p *Runner) appendFunctions(functions ...func()) error {
	for _, function := range functions {
		err := p.Add(function)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Runner) wait(waitGroup *sync.WaitGroup) error {
	// If no timeout was provided
	if p.Timeout <= 0 {
		waitGroup.Wait()
		return nil
	}

	channel := make(chan struct{})

	go func() {
		waitGroup.Wait()
		close(channel)
	}()

	select {
	case <-channel:
		return nil
	case <-time.After(p.Timeout):
		return errors.New(timeoutError)
	}
}
