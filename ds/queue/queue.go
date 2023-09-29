package queue

import "fmt"

// Queue follows FIFO order
type Queue struct {
	queue []int
}

// NewQueue returns a Queue object
func NewQueue() *Queue {
	return &Queue{
		queue: []int{},
	}
}

// EnQueue adds an element to the end of the queue
func (q *Queue) EnQueue(x int) {
	if q.queue == nil {
		fmt.Println("need to initialize the queue.")
	}
	q.queue = append(q.queue, x)
}

// DeQueue will remove the first element from Queue
func (q *Queue) DeQueue() {
	if len(q.queue) > 0 {
		fmt.Println("deQueuing ", q.queue[0])
		q.queue = q.queue[1:]
		fmt.Println("after deQueuing queue is", q.queue)
	} else {
		fmt.Println("looks like the queue is empty")
	}
}
