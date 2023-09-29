package main

import "github.com/dnanam/hashitout/ds/queue"

// main calls queue.NewQueue() and creates a Queue object to then EnQueue and DeQueue objects form the Queue
func main() {
	q := queue.NewQueue()
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.DeQueue()   // removes 1
	q.DeQueue()   // removes 2
	q.EnQueue(12) // add 12
	q.DeQueue()   // removes 3
	q.DeQueue()   // removes 12
	q.DeQueue()   // is empty over here

}
