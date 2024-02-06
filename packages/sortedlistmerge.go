package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	current := n1
	hold := current
	for current != nil {
		hold = current
		current = current.Next
	}

	hold.Next = n2
	return ListSort(n1)
}
