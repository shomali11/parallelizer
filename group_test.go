package parallelizer

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGroup_Add_NilFunction(t *testing.T) {
	group := DefaultGroup()

	err := group.Add(nil)
	assert.NotNil(t, err)

	err = group.Run()
	assert.Nil(t, err)
}

func TestGroup_NoTimeout(t *testing.T) {
	value1 := 1
	value2 := 2

	group := DefaultGroup()

	err := group.Add(func() {
		value1 = 11
	})

	assert.Nil(t, err)

	err = group.Add(func() {
		value2 = 22
	})

	assert.Nil(t, err)

	err = group.Run()

	assert.Nil(t, err)
	assert.Equal(t, value1, 11)
	assert.Equal(t, value2, 22)
}

func TestGroup_LongTimeout(t *testing.T) {
	value1 := 1
	value2 := 2

	group := NewGroup(&Options{Timeout: time.Minute})

	group.Add(func() {
		value1 = 11
	})

	group.Add(func() {
		value2 = 22
	})

	err := group.Run()

	assert.Nil(t, err)
	assert.Equal(t, value1, 11)
	assert.Equal(t, value2, 22)
}

func TestGroup_ShortTimeout(t *testing.T) {
	value1 := 1
	value2 := 2

	group := NewGroup(&Options{Timeout: time.Second})

	group.Add(func() {
		time.Sleep(2 * time.Second)
		value1 = 11
	})

	group.Add(func() {
		time.Sleep(2 * time.Second)
		value2 = 22
	})

	err := group.Run()

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "timeout")
	assert.Equal(t, value1, 1)
	assert.Equal(t, value2, 2)

	time.Sleep(3 * time.Second)
}

func TestGroup_LargeWorkerPoolSize(t *testing.T) {
	value1 := 1
	value2 := 2

	group := NewGroup(&Options{WorkerPoolSize: 100})

	group.Add(func() {
		value1 = 11
	})

	group.Add(func() {
		value2 = 22
	})

	err := group.Run()

	assert.Nil(t, err)
	assert.Equal(t, value1, 11)
	assert.Equal(t, value2, 22)
}

func TestGroup_SmallWorkerPoolSize(t *testing.T) {
	value1 := 1
	value2 := 2

	group := NewGroup(&Options{WorkerPoolSize: 1})

	group.Add(func() {
		time.Sleep(time.Second)
		value1 = 11
	})

	group.Add(func() {
		time.Sleep(time.Second)
		value2 = 22
	})

	err := group.Run()

	assert.Nil(t, err)
	assert.Equal(t, value1, 11)
	assert.Equal(t, value2, 22)
}
