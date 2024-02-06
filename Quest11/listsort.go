package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	current := l
	if l == nil {
		return l
	}
	for current != nil {
		ListSort(current.Next)
		if current.Next != nil {
			if current.Data > current.Next.Data {
				current.Data, current.Next.Data = current.Next.Data, current.Data
			}
		}
		current = current.Next
	}
	return l
}
