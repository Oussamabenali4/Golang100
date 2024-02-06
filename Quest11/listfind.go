package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	if l.Head == nil {
		return nil
	}
	current := l.Head
	for current != nil {
		if comp(ref, current.Data) {
			return &current.Data
		}
		current = current.Next
	}
	return nil
}
