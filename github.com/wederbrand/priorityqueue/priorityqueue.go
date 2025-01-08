package priorityqueue

import (
	"cmp"
	"slices"
)

type State struct {
	Data     interface{}
	Priority int
}

type Queue struct {
	states []*State
}

func NewQueue() *Queue {
	q := new(Queue)
	q.states = make([]*State, 0)

	return q
}

func (q *Queue) Add(s *State) {
	search, _ := slices.BinarySearchFunc(q.states, s, func(s1 *State, s2 *State) int {
		if cmp.Compare(s1.Priority, s2.Priority) == 0 {
			return -1 // maintain insert order if priorities are identical
		}
		return cmp.Compare(s1.Priority, s2.Priority)
	})
	q.states = slices.Insert(q.states, search, s)
}

func (q *Queue) HasNext() bool {
	return len(q.states) > 0
}

func (q *Queue) Next() *State {
	s := q.states[0]
	q.states = q.states[1:len(q.states)]
	return s
}
