package controllers

import (
	"sync"
	"fmt"
)

type Event interface {
	Validate(interface{}) error
	Implement(interface{}) error
}

type EventQueue struct {
	stack []Event
	lock sync.Mutex
}

func (q *EventQueue) Push(e Event) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.stack = append(q.stack, e)
}

func (q	*EventQueue) Pop() (Event, error) {
	q.lock.Lock()
	defer q.lock.Unlock()

	l := len(q.stack)
	if l == 0 {
		return nil, fmt.Errorf("event queue empty")
	}

	event := q.stack[l - 1]
	q.stack = q.stack[:l - 1]

	return event, nil
}

func (q *EventQueue) HasNext() bool {
	q.lock.Lock()
	defer q.lock.Unlock()

	return len(q.stack) > 0
}
