package parallelizer

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRunner_Run_NilFunction(t *testing.T) {
	runner := Runner{}
	err := runner.Run(nil, nil)

	assert.NotNil(t, err)
}

func TestRunner_Add_NilFunction(t *testing.T) {
	runner := Runner{}

	err := runner.Add(nil)
	assert.NotNil(t, err)

	err = runner.Run()
	assert.Nil(t, err)
}

func TestRunner_Add_NoTimeout(t *testing.T) {
	value1 := 1
	value2 := 2

	runner := Runner{}

	err := runner.Add(func() {
		value1 = 11
	})

	assert.Nil(t, err)

	err = runner.Add(func() {
		value2 = 22
	})

	assert.Nil(t, err)

	err = runner.Run()

	assert.Nil(t, err)
	assert.Equal(t, value1, 11)
	assert.Equal(t, value2, 22)
}

func TestRunner_Run_NoTimeout(t *testing.T) {
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

func TestRunner_Add_LongTimeout(t *testing.T) {
	value1 := 1
	value2 := 2

	runner := Runner{Timeout: time.Minute}

	runner.Add(func() {
		value1 = 11
	})

	runner.Add(func() {
		value2 = 22
	})

	err := runner.Run()

	assert.Nil(t, err)
	assert.Equal(t, value1, 11)
	assert.Equal(t, value2, 22)
}

func TestRunner_Run_LongTimeout(t *testing.T) {
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

func TestRunner_Add_ShortTimeout(t *testing.T) {
	value1 := 1
	value2 := 2

	runner := Runner{Timeout: time.Second}

	runner.Add(func() {
		time.Sleep(time.Minute)
		value1 = 11
	})

	runner.Add(func() {
		time.Sleep(time.Minute)
		value2 = 22
	})

	err := runner.Run()

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "timeout")
	assert.Equal(t, value1, 1)
	assert.Equal(t, value2, 2)
}

func TestRunner_Run_ShortTimeout(t *testing.T) {
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
