package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	node := l
	count := 0
	for count < pos && node != nil {
		node = node.Next
		count++
	}
	return node
}
