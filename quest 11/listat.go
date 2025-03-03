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

func ListAt(l *NodeL, pos int) *NodeL {
	if l == nil || pos < 0 {
		return nil
	}

	current := l
	for i := 0; i < pos; i++ {
		if current.Next == nil {
			return nil // Reached the end of the list before reaching the desired position
		}
		current = current.Next
	}

	return current
}
