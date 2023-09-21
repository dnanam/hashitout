package implementation

type Node struct {
	val  int
	next *Node
}

type LinkedList struct {
	head *Node
}

/*

Basic operations on a LinkedList will be
* Insert
* Delete

*/

// New returns a LinkedList
func New() *LinkedList {
	return &LinkedList{head: &Node{}}
}

func (l *LinkedList) Insert(value any) {
	newNode := &Node{}
	newNode.val = value.(int)

	current := l.head
	for current.next != nil {
		current = current.next
	}

	current.next = newNode
}

func (l *LinkedList) Delete(value any) {

	if l.head == nil {
		return
	}

	current := l.head
	for current.next != nil {
		if current.next.val != value.(int) {
			current = current.next
		}
	}

	if current.next != nil {
		current.next = current.next.next
	}
}
