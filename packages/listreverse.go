package piscine

func ListReverse(l *List) {
	var prev *NodeL
	next := l.Head.Next
	current := l.Head
	l.Tail = l.Head
	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	l.Head = prev
}
