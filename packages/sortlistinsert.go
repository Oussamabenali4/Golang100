package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	newnode := NodeI{Data: data_ref}
	current := l
	if data_ref < current.Data {
		newnode.Next = l
		l = &newnode
		return l
	}
	for current != nil {
		if data_ref >= current.Data {
			if current.Next != nil {
				newnode.Next = current.Next // if it is not nill
			}
			current.Next = &newnode
			break
		}
		current = current.Next
	}
	return l
}
