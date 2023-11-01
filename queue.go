package main

import (
	"sync"
)

type Queue struct {
	activQ []*Item
	cond   *sync.Cond
}

type Item struct {
	body map[string]interface{}
}

func New() *Queue {
	return &Queue{activQ: []*Item{}, cond: sync.NewCond(&sync.Mutex{})}
}

func (q *Queue) Add(Item *Item) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()

	q.activQ = append(q.activQ, Item)
	q.cond.Signal()
}

func (q *Queue) Get() *Item {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	if len(q.activQ) == 0 {
		return nil
	}
	item := q.activQ[0]
	q.activQ = q.activQ[1:]
	return item
}
