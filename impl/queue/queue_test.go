package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	intQueue = &Queue[int]{}
	strQueue = &Queue[string]{}
)

func InitQueue() {
	intQueue = &Queue[int]{}
	strQueue = &Queue[string]{}
}

func TestPush(t *testing.T) {
	InitQueue()

	intQueue.Push(1, 2, 3)
	assert.Equal(t, 3, intQueue.Len())
	t.Logf("int queue: %v", intQueue.Display())

	strQueue.Push("a", "b", "c")
	t.Logf("str queue: %v", strQueue.Display())
}

func TestPop(t *testing.T) {
	InitQueue()

	v, err := intQueue.Pop()
	assert.ErrorContains(t, err, "queue is empty")
	assert.Equal(t, 0, v)
	t.Logf("int queue: %v", intQueue.Display())

	intQueue.Push(1, 2, 3)
	v, err = intQueue.Pop()
	assert.ErrorIs(t, err, nil)
	assert.Equal(t, 1, v)
	t.Logf("int queue: %v", intQueue.Display())
}
