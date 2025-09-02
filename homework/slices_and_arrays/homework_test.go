package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type CircularQueue struct {
	values []int
	head   int
	tail   int
	count  int
}

func NewCircularQueue(capacity int) CircularQueue {
	return CircularQueue{
		values: make([]int, capacity),
		head:   0,
		tail:   0,
		count:  0,
	}
}

func (q *CircularQueue) Push(value int) bool {
	if q.Full() {
		return false
	}
	q.values[q.tail] = value
	q.tail = (q.tail + 1) % len(q.values)
	q.count++
	return true
}

func (q *CircularQueue) Pop() bool {
	if q.Empty() {
		return false
	}
	q.head = (q.head + 1) % len(q.values)
	q.count--
	return true
}

func (q *CircularQueue) Front() int {
	if q.Empty() {
		return -1
	}
	return q.values[q.head]
}

func (q *CircularQueue) Back() int {
	if q.Empty() {
		return -1
	}
	capacity := len(q.values)
	idx := (q.tail - 1 + capacity) % capacity
	return q.values[idx]
}

func (q *CircularQueue) Empty() bool {
	return q.count == 0
}

func (q *CircularQueue) Full() bool {
	return q.count == len(q.values)
}

func TestCircularQueue(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue(queueSize)

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())

	assert.Equal(t, -1, queue.Front())
	assert.Equal(t, -1, queue.Back())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Push(1))
	assert.True(t, queue.Push(2))
	assert.True(t, queue.Push(3))
	assert.False(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{1, 2, 3}, queue.values))

	assert.False(t, queue.Empty())
	assert.True(t, queue.Full())

	assert.Equal(t, 1, queue.Front())
	assert.Equal(t, 3, queue.Back())

	assert.True(t, queue.Pop())
	assert.False(t, queue.Empty())
	assert.False(t, queue.Full())
	assert.True(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{4, 2, 3}, queue.values))

	assert.Equal(t, 2, queue.Front())
	assert.Equal(t, 4, queue.Back())

	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())
}

func TestCircularQueueOverflow(t *testing.T) {
	q := NewCircularQueue(3)

	// заполняем до capacity
	if !q.Push(1) {
		t.Errorf("expected Push(1) = true, got false")
	}
	if !q.Push(2) {
		t.Errorf("expected Push(2) = true, got false")
	}
	if !q.Push(3) {
		t.Errorf("expected Push(3) = true, got false")
	}

	// проверяем, что очередь теперь полная
	if !q.Full() {
		t.Errorf("expected queue to be full")
	}

	// пытаемся вставить лишний элемент
	if q.Push(4) {
		t.Errorf("expected Push(4) = false (overflow), got true")
	}
}

func TestCircularQueueWrapAround(t *testing.T) {
	q := NewCircularQueue(3)

	// Заполняем очередь [1,2,3]
	q.Push(1)
	q.Push(2)
	q.Push(3)

	// Удаляем один элемент (теперь ожидаем [2,3])
	q.Pop()

	// Добавляем новый элемент (ожидаем очередь[2,3,4], причём 4 попадёт в начало массива [4,2,3] где голова начинается с 2)
	q.Push(4)

	// Теперь вынимаем все элементы в массив
	var got []int
	for !q.Empty() {
		got = append(got, q.Front())
		q.Pop()
	}

	want := []int{2, 3, 4}
	if len(got) != len(want) {
		t.Fatalf("expected %v elements, got %v", len(want), len(got))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("expected sequence %v, got %v", want, got)
			break
		}
	}
}
