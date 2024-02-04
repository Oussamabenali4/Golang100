package piscine

func ListReverse(l *List) {
	var prev *NodeL
	current := l.Head
	next = l.Head.Next
	hold := l.Head
	for current != nil {
		current.Next = prev
		current = current.Next
	}
	l.Tail = hold
	current.Next = nil
}
