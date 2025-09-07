package main

import (
	"container/heap"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Task struct {
	Identifier int
	Priority   int
}

type QueueTask struct {
	Task
	heapPriority int
	idx          int
}

type Scheduler struct {
	tasks    *TaskHeap
	registry map[int]*QueueTask
}

func NewScheduler() Scheduler {
	tasks := TaskHeap([]*QueueTask{})
	heap.Init(&tasks)

	return Scheduler{
		tasks:    &tasks,
		registry: make(map[int]*QueueTask),
	}
}

func (s *Scheduler) AddTask(task Task) {
	hTask := QueueTask{task, task.Priority, 0}
	s.registry[task.Identifier] = &hTask
	heap.Push(s.tasks, &hTask)
}

func (s *Scheduler) ChangeTaskPriority(taskID int, newPriority int) {
	if t, exists := s.registry[taskID]; exists {
		t.heapPriority = newPriority
		heap.Fix(s.tasks, t.idx)
	}
}

func (s *Scheduler) GetTask() (Task, error) {
	if s.tasks.Len() == 0 {
		return Task{}, errors.New("task queue is empty")
	}
	task := heap.Pop(s.tasks).(*QueueTask).Task
	delete(s.registry, task.Identifier)
	return task, nil
}

type TaskHeap []*QueueTask

func (h *TaskHeap) Len() int {
	return len(*h)
}

func (h *TaskHeap) Less(i, j int) bool {
	return (*h)[i].heapPriority > (*h)[j].heapPriority
}

func (h *TaskHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
	(*h)[i].idx = i
	(*h)[j].idx = j
}

func (h *TaskHeap) Push(x interface{}) {
	item := x.(*QueueTask)
	item.idx = len(*h)
	*h = append(*h, x.(*QueueTask))
}

func (h *TaskHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func TestTrace(t *testing.T) {
	task1 := Task{Identifier: 1, Priority: 10}
	task2 := Task{Identifier: 2, Priority: 20}
	task3 := Task{Identifier: 3, Priority: 30}
	task4 := Task{Identifier: 4, Priority: 40}
	task5 := Task{Identifier: 5, Priority: 50}

	scheduler := NewScheduler()
	scheduler.AddTask(task1)
	scheduler.AddTask(task2)
	scheduler.AddTask(task3)
	scheduler.AddTask(task4)
	scheduler.AddTask(task5)

	task, err := scheduler.GetTask()
	assert.NoError(t, err)
	assert.Equal(t, task5, task)

	task, err = scheduler.GetTask()
	assert.NoError(t, err)
	assert.Equal(t, task4, task)

	scheduler.ChangeTaskPriority(1, 100)

	task, err = scheduler.GetTask()
	assert.NoError(t, err)
	assert.Equal(t, task1, task)

	task, err = scheduler.GetTask()
	assert.NoError(t, err)
	assert.Equal(t, task3, task)
}
