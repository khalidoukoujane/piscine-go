package piscine

func ListClear(l *List) {
	curr := l.Head
	for curr != nil {
		nextNode := curr.Next
		curr.Next = nil
		curr = nextNode
	}
	l.Head = nil
}
