package linkedlist

import (
	"testing"
)

func TestCreateEmpty(t *testing.T) {
	var values []int32
	node := Create(values)
	if node != nil {
		t.Error("LinkedList is not empty")
	}
}

func TestCreate(t *testing.T) {
	values := []int32{1, 2, 3, 4, 5}
	node := Create(values)
	cnt := 0
	for node != nil {
		if node.Value != values[cnt] {
			t.Error("LinkedList is invalid")
		}
		cnt++
		node = node.Next
	}
}
