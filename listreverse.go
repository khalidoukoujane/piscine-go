package piscine

func ListReverse(l *List) {
	var prev, current, next *NodeL
	current = l.Head

	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}

	l.Head, l.Tail = prev, l.Head
}
