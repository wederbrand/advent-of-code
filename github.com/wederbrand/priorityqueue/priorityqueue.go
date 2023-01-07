package priorityqueue

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
	// add State to Queue and sort it by priority
	if len(q.states) == 0 {
		q.states = append(q.states, s)
	} else {
		for i, s2 := range q.states {
			if s.Priority < s2.Priority {
				// insert here
				q.states = append(q.states, nil)   // make room for copy
				copy(q.states[i+1:], q.states[i:]) // make room at index i, overwriting the nil
				q.states[i] = s                    // insert s
				return
			}
		}
		q.states = append(q.states, s)
	}
}

func (q *Queue) HasNext() bool {
	return len(q.states) > 0
}

func (q *Queue) Next() *State {
	s := q.states[0]
	q.states = q.states[1:len(q.states)]
	return s
}
