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

	if l.Head == nil { // If the list is empty, set both Head and Tail to the new node
		l.Head = newNode
		l.Tail = newNode
	} else {
		l.Tail.Next = newNode // Link the current tail to the new node
		l.Tail = newNode      // Update the tail pointer to the new node
	}
}
