package piscine

func ListMerge(l1 *List, l2 *List) {
	if l1.Head == nil || l2.Head == nil {
		return
	}
	current := l2.Head
	for current != nil {
		ListPushBack(l1, current.Data)
		current = current.Next
	}
}
