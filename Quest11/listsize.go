package piscine

func ListSize(l *List) int {
	count := 0
	current := l.Head
	for current != nil {
		count += 1
		current = current.Next
	}
	return count
}
