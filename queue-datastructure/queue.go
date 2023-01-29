package queue_datastructure

type Queue struct {
	data chan interface{}
}

func NewQueue(i int) *Queue {
	return &Queue{
		data: make(chan interface{}, i),
	}
}

func (q *Queue) Push(d interface{}) {
	q.data <- d
}

func (q *Queue) Pop() interface{} {
	return <-q.data
}

func (q *Queue) Length() int {
	return len(q.data)
}
