package parallelizer

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRunnerNoTimeout(t *testing.T) {
	value1 := 1
	value2 := 2

	fun1 := func() {
		value1 = 11
	}

	fun2 := func() {
		value2 = 22
	}

	runner := Runner{}
	finished := runner.Run(fun1, fun2)

	assert.True(t, finished)
	assert.Equal(t, value1, 11)
	assert.Equal(t, value2, 22)
}

func TestRunnerLongTimeout(t *testing.T) {
	value1 := 1
	value2 := 2

	fun1 := func() {
		value1 = 11
	}

	fun2 := func() {
		value2 = 22
	}

	runner := Runner{Timeout: time.Minute}
	finished := runner.Run(fun1, fun2)

	assert.True(t, finished)
	assert.Equal(t, value1, 11)
	assert.Equal(t, value2, 22)
}

func TestRunnerShortTimeout(t *testing.T) {
	value1 := 1
	value2 := 2

	fun1 := func() {
		time.Sleep(time.Minute)
		value1 = 11
	}

	fun2 := func() {
		time.Sleep(time.Minute)
		value2 = 22
	}

	runner := Runner{Timeout: time.Second}
	finished := runner.Run(fun1, fun2)

	assert.False(t, finished)
	assert.Equal(t, value1, 1)
	assert.Equal(t, value2, 2)
}
