package linkedlist

type Node struct {
	Value int32
	Next  *Node
}

func (n *Node) Append(value int32) *Node {
	n.Next = &Node{Value: value, Next: nil}
	return n.Next
}

func Create(values []int32) *Node {
	if len(values) == 0 {
		return nil
	}

	var head = &Node{Value: values[0], Next: nil}
	cur := head

	for i := 1; i < len(values); i++ {
		cur = cur.Append(values[i])
	}

	return head
}
