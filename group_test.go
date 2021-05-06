package parallelizer

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type store struct {
	value int
	mutex sync.Mutex
}

func (s *store) Get() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.value
}

func (s *store) Set(val int) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.value = val
}

func TestGroup_Add_NilFunction(t *testing.T) {
	group := NewGroup()
	defer group.Close()

	err := group.Add(nil)
	assert.NotNil(t, err)

	err = group.Wait()
	assert.Nil(t, err)
}

func TestGroup_NoTimeout(t *testing.T) {
	store1 := store{value: 1}
	store2 := store{value: 2}

	group := NewGroup()
	defer group.Close()

	err := group.Add(func() {
		store1.Set(11)
	})

	assert.Nil(t, err)

	err = group.Add(func() {
		store2.Set(22)
	})

	assert.Nil(t, err)

	err = group.Wait()

	assert.Nil(t, err)
	assert.Equal(t, store1.Get(), 11)
	assert.Equal(t, store2.Get(), 22)
}

func TestGroup_LongTimeout(t *testing.T) {
	store1 := store{value: 1}
	store2 := store{value: 2}

	group := NewGroup()
	defer group.Close()

	group.Add(func() {
		store1.Set(11)
	})

	group.Add(func() {
		store2.Set(22)
	})

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	err := group.Wait(WithContext(ctx))

	assert.Nil(t, err)
	assert.Equal(t, store1.Get(), 11)
	assert.Equal(t, store2.Get(), 22)
}

func TestGroup_ShortTimeout(t *testing.T) {
	store1 := store{value: 1}
	store2 := store{value: 2}

	group := NewGroup()
	defer group.Close()

	group.Add(func() {
		time.Sleep(2 * time.Second)
		store1.Set(11)
	})

	group.Add(func() {
		time.Sleep(2 * time.Second)
		store2.Set(22)
	})

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err := group.Wait(WithContext(ctx))

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "context deadline exceeded")
	assert.Equal(t, store1.Get(), 1)
	assert.Equal(t, store2.Get(), 2)
}

func TestGroup_CancelledContext(t *testing.T) {
	store1 := store{value: 1}
	store2 := store{value: 2}

	group := NewGroup()
	defer group.Close()

	group.Add(func() {
		time.Sleep(2 * time.Second)
		store1.Set(11)
	})

	group.Add(func() {
		time.Sleep(2 * time.Second)
		store2.Set(22)
	})

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	err := group.Wait(WithContext(ctx))

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "context canceled")
	assert.Equal(t, store1.Get(), 1)
	assert.Equal(t, store2.Get(), 2)
}

func TestGroup_LargeGroupSize(t *testing.T) {
	store1 := store{value: 1}
	store2 := store{value: 2}

	group := NewGroup(WithPoolSize(100))
	defer group.Close()

	group.Add(func() {
		store1.Set(11)
	})

	group.Add(func() {
		store2.Set(22)
	})

	err := group.Wait()

	assert.Nil(t, err)
	assert.Equal(t, store1.Get(), 11)
	assert.Equal(t, store2.Get(), 22)
}

func TestGroup_SmallGroupSize(t *testing.T) {
	store1 := store{value: 1}
	store2 := store{value: 2}

	group := NewGroup(WithPoolSize(1))
	defer group.Close()

	group.Add(func() {
		time.Sleep(time.Second)
		store1.Set(11)
	})

	group.Add(func() {
		time.Sleep(time.Second)
		store2.Set(22)
	})

	err := group.Wait()

	assert.Nil(t, err)
	assert.Equal(t, store1.Get(), 11)
	assert.Equal(t, store2.Get(), 22)
}

func TestGroup_SmallJobQueueSize(t *testing.T) {
	store1 := store{value: 1}
	store2 := store{value: 2}

	group := NewGroup(WithJobQueueSize(1))
	defer group.Close()

	group.Add(func() {
		time.Sleep(time.Second)
		store1.Set(11)
	})

	group.Add(func() {
		time.Sleep(time.Second)
		store2.Set(22)
	})

	err := group.Wait()

	assert.Nil(t, err)
	assert.Equal(t, store1.Get(), 11)
	assert.Equal(t, store2.Get(), 22)
}
