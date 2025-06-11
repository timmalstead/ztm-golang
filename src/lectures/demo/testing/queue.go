package queue

type Queue struct {
	items    []int
	capacity int
}

func (q *Queue) AppendItem(item int) bool {
	if len(q.items) == q.capacity {
		return false
	}
	q.items = append(q.items, item)
	return true
}

func (q *Queue) NextItem() (int, bool) {
	if len(q.items) > 0 {
		var next = q.items[0]
		q.items = q.items[1:]
		return next, true
	} else {
		return 0, false
	}
}

func NewQueue(capacity int) Queue {
	// make([]int, 0, 10) allocates an underlying array of size 10 and returns a slice of length 0 and capacity 10 that is backed by this underlying array.
	// not used to doing it this way but I understand the logic, I suppose
	return Queue{items: make([]int, 0, capacity), capacity: capacity}
}
