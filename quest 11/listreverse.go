package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	newNode := &NodeL{Data: data}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		l.Tail.Next = newNode
		l.Tail = newNode
	}
}

func ListReverse(l *List) {
	var prev, current, next *NodeL
	current = l.Head

	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}

	// Swap head and tail pointers
	l.Tail, l.Head = l.Head, prev
}
