package task_scheduler

import (
	"fmt"
	"time"
)

// Task represents a scheduled job.
type Task struct {
	ID        int
	Name      string
	Scheduled time.Time
	Execute   func()
}

// SchedulerError is a typed error returned by scheduler operations.
type SchedulerError struct {
	Code    int
	Message string
}

func (e *SchedulerError) Error() string {
	return fmt.Sprintf("code=%d: %s", e.Code, e.Message)
}

// TaskScheduler holds scheduled tasks and next ID.
type TaskScheduler struct {
	tasks  []*Task
	nextID int
}

// NewTaskScheduler initializes and returns a TaskScheduler.
func NewTaskScheduler() *TaskScheduler {
	return &TaskScheduler{
		tasks:  make([]*Task, 0),
		nextID: 1,
	}
}

// AddTask validates input and appends a new task with auto-increment ID.
// Returns the created Task and nil error on success.
// If name is empty returns SchedulerError with Code 1.
func (ts *TaskScheduler) AddTask(name string, scheduled time.Time, execFn func()) (*Task, *SchedulerError) {
	if name == "" {
		return nil, &SchedulerError{
			Code:    1,
			Message: "task name cannot be empty",
		}
	}
	task := &Task{
		ID:        ts.nextID,
		Name:      name,
		Scheduled: scheduled,
		Execute:   execFn,
	}
	ts.tasks = append(ts.tasks, task)
	ts.nextID++
	return task, nil
}

// GetTask finds a task by ID. If not found returns SchedulerError with Code 2.
func (ts *TaskScheduler) GetTask(id int) (*Task, *SchedulerError) {
	for _, t := range ts.tasks {
		if t.ID == id {
			return t, nil
		}
	}
	return nil, &SchedulerError{
		Code:    2,
		Message: "task not found",
	}
}

// TaskIterator iterates over tasks in insertion order.
type TaskIterator struct {
	scheduler    *TaskScheduler
	currentIndex int
}

// Iterator returns a new TaskIterator for the scheduler.
func (ts *TaskScheduler) Iterator() *TaskIterator {
	return &TaskIterator{
		scheduler:    ts,
		currentIndex: 0,
	}
}

// Next returns the next task and true if available, otherwise nil,false.
func (it *TaskIterator) Next() (*Task, bool) {
	if it.scheduler == nil {
		return nil, false
	}
	if it.currentIndex >= len(it.scheduler.tasks) {
		return nil, false
	}
	t := it.scheduler.tasks[it.currentIndex]
	it.currentIndex++
	return t, true
}

// RunScheduledTasks executes tasks whose Scheduled time is <= now.
// This implementation calls task.Execute synchronously in the current goroutine.
func (ts *TaskScheduler) RunScheduledTasks() {
	now := time.Now()
	for _, t := range ts.tasks {
		if !t.Scheduled.After(now) {
			if t.Execute != nil {
				t.Execute()
			}
		}
	}
}
