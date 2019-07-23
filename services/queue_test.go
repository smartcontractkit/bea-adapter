package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPriorityQueue_Len(t *testing.T) {
	arr := make(PriorityQueue, 100)
	assert.Equal(t, 100, arr.Len())

	slice := PriorityQueue{}
	slice = append(slice, &Item{}, &Item{})
	assert.Equal(t, 2, slice.Len())
}

func TestPriorityQueue_Less(t *testing.T) {
	// Remember we inverse the Less() function
	// in order to get the newest items first

	newest := &Item{Year: 2019, Month: 7}
	mid := &Item{Year: 2019, Month: 6}
	oldest := &Item{Year: 2018, Month: 12}

	queue := PriorityQueue{}
	queue = append(queue, newest, mid, oldest)

	assert.False(t, queue.Less(2, 1))
	assert.False(t, queue.Less(1, 0))
	assert.False(t, queue.Less(2, 0))
	assert.True(t, queue.Less(1, 2))
	assert.True(t, queue.Less(0, 1))
}

func TestPriorityQueue_Push(t *testing.T) {
	i1 := &Item{Value: 1}
	i2 := &Item{Value: 2}

	queue := PriorityQueue{}
	queue.Push(i1)
	queue.Push(i2)

	assert.Equal(t, i1, queue[0])
	assert.Equal(t, i2, queue[1])
}

func TestPriorityQueue_Pop(t *testing.T) {
	i1 := &Item{Value: 1, Index: 0}
	i2 := &Item{Value: 2, Index: 1}

	queue := PriorityQueue{}
	queue = append(queue, i1, i2)

	assert.Equal(t, 2, len(queue))

	returnedItem := queue.Pop()
	assert.Equal(t, 1, len(queue))
	assert.Equal(t, i2, returnedItem)
	assert.Equal(t, i1, queue[len(queue)-1])
}

func TestPriorityQueue_Swap(t *testing.T) {
	i1 := &Item{Value: 1, Index: 0}
	i2 := &Item{Value: 2, Index: 1}

	queue := PriorityQueue{}
	queue = append(queue, i1, i2)
	queue.Swap(0, 1)

	assert.Equal(t, 1, i1.Index)
	assert.Equal(t, 0, i2.Index)
}
