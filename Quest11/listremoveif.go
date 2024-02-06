package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	if l.Head == nil {
		return
	}
	current := l.Head
	var prev *NodeL
	var next *NodeL
	// the case where the head has to be replaced
	for current.Data == data_ref {
		if current.Next != nil {
			current = current.Next
		} else {
			l.Head = nil
			return
		}

	}
	l.Head = current

	for current != nil {
		next = current.Next
		if current.Data == data_ref {
			prev.Next = next
		} else {
			prev = current
		}
		current = current.Next
	}
	l.Tail = prev
}
