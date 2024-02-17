package piscine

func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	}
	curr := l.Head
	for curr != nil {
		if curr.Next == nil {
			return curr.Data
		}
		curr = curr.Next
	}
	return curr.Data
}
