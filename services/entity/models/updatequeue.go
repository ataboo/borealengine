package models

import (
	"sync"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
)

const (
	UServer = "srv"
	UDamage = "dmg"
	UPlayer = "pl"
)

type RawEvent struct {
	Type string `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type Update interface {
	GetType() string
}

type ServerUpdate struct {
	Transform  Transform `json:"transform"`
	Attack     bool      `json:"attack"`
	Transition bool      `json:"transition"`
	Chat       string    `json:"chat"`
}

func(u ServerUpdate) GetType() string {
	return UServer
}

type PlayerUpdate struct {
	Vitals Vitals `json:"vitals"`
}

func(u PlayerUpdate) GetType() string {
	return UPlayer
}

type DamageUpdate struct {
	Target bson.ObjectId `json:"target"`
	Amount float32 `json:"amount"`
}

func(u DamageUpdate) GetType() string {
	return UDamage
}

type UpdateQueue struct {
	stack []Update
	lock  sync.Mutex
}

func (q *UpdateQueue) Push(u Update) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.stack = append(q.stack, u)
}

func (q *UpdateQueue) Pop() (Update, error) {
	q.lock.Lock()
	defer q.lock.Unlock()

	l := len(q.stack)
	if l == 0 {
		return nil, fmt.Errorf("event queue empty")
	}

	event := q.stack[l-1]
	q.stack = q.stack[:l-1]

	return event, nil
}

func (q *UpdateQueue) Clear() {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.stack = []Update{}
}

func (q *UpdateQueue) HasNext() bool {
	q.lock.Lock()
	defer q.lock.Unlock()

	return len(q.stack) > 0
}

func (q *UpdateQueue) PopAll() []Update {
	q.lock.Lock()
	defer q.lock.Unlock()

	updates := q.stack
	q.Clear()

	return updates
}
