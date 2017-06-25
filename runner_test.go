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
	err := runner.Run(fun1, fun2)

	assert.Nil(t, err)
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
	err := runner.Run(fun1, fun2)

	assert.Nil(t, err)
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
	err := runner.Run(fun1, fun2)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "timeout")
	assert.Equal(t, value1, 1)
	assert.Equal(t, value2, 2)
}
