package parallelizer

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGroup_Add_NilFunction(t *testing.T) {
	group := NewGroup()
	defer group.Close()

	err := group.Add(nil)
	assert.NotNil(t, err)

	err = group.Wait()
	assert.Nil(t, err)
}

func TestGroup_NoTimeout(t *testing.T) {
	value1 := 1
	value2 := 2

	group := NewGroup()
	defer group.Close()

	err := group.Add(func() {
		value1 = 11
	})

	assert.Nil(t, err)

	err = group.Add(func() {
		value2 = 22
	})

	assert.Nil(t, err)

	err = group.Wait()

	assert.Nil(t, err)
	assert.Equal(t, value1, 11)
	assert.Equal(t, value2, 22)
}

func TestGroup_LongTimeout(t *testing.T) {
	value1 := 1
	value2 := 2

	group := NewGroup()
	defer group.Close()

	group.Add(func() {
		value1 = 11
	})

	group.Add(func() {
		value2 = 22
	})

	err := group.Wait(WithTimeout(time.Minute))

	assert.Nil(t, err)
	assert.Equal(t, value1, 11)
	assert.Equal(t, value2, 22)
}

func TestGroup_ShortTimeout(t *testing.T) {
	value1 := 1
	value2 := 2

	group := NewGroup()
	defer group.Close()

	group.Add(func() {
		time.Sleep(2 * time.Second)
		value1 = 11
	})

	group.Add(func() {
		time.Sleep(2 * time.Second)
		value2 = 22
	})

	err := group.Wait(WithTimeout(time.Second))

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "timeout")
	assert.Equal(t, value1, 1)
	assert.Equal(t, value2, 2)
}

func TestGroup_LargeGroupSize(t *testing.T) {
	value1 := 1
	value2 := 2

	group := NewGroup(WithPoolSize(100))
	defer group.Close()

	group.Add(func() {
		value1 = 11
	})

	group.Add(func() {
		value2 = 22
	})

	err := group.Wait()

	assert.Nil(t, err)
	assert.Equal(t, value1, 11)
	assert.Equal(t, value2, 22)
}

func TestGroup_SmallGroupSize(t *testing.T) {
	value1 := 1
	value2 := 2

	group := NewGroup(WithPoolSize(1))
	defer group.Close()

	group.Add(func() {
		time.Sleep(time.Second)
		value1 = 11
	})

	group.Add(func() {
		time.Sleep(time.Second)
		value2 = 22
	})

	err := group.Wait()

	assert.Nil(t, err)
	assert.Equal(t, value1, 11)
	assert.Equal(t, value2, 22)
}

func TestGroup_SmallJobQueueSize(t *testing.T) {
	value1 := 1
	value2 := 2

	group := NewGroup(WithJobQueueSize(1))
	defer group.Close()

	group.Add(func() {
		time.Sleep(time.Second)
		value1 = 11
	})

	group.Add(func() {
		time.Sleep(time.Second)
		value2 = 22
	})

	err := group.Wait()

	assert.Nil(t, err)
	assert.Equal(t, value1, 11)
	assert.Equal(t, value2, 22)
}
