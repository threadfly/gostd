package list_test

import (
	"testing"

	//"fmt"

	"github.com/threadfly/gostd/list"
)

type Node struct {
	C string
	N list.ListNode
}

func NewNode(c string) *Node {
	n := &Node{
		C: c,
	}
	n.N.V = n
	return n
}

func Test_List(t *testing.T) {
	n1 := NewNode("1111")

	n2 := NewNode("2222")
	n3 := NewNode("3333")
	n4 := NewNode("4444")

	L := list.NewList()

	L.AddTail(&n1.N)
	L.AddTail(&n2.N)
	L.AddTail(&n3.N)
	L.AddTail(&n4.N)

	f := func(v interface{}) {
		if v == nil {
			return
		}
		V := v.(*Node)
		t.Errorf("node.C : %s", V.C)
	}

	L.Traverse(f)

	L.DeleteHead()
	L.DeleteHead()
	L.DeleteHead()
	L.DeleteHead()
	L.DeleteHead()
	L.Traverse(f)

	t.Errorf("&n4.N : %v", n4.N)
	L.AddHead(&n3.N)
	L.AddTail(&n4.N)
	L.DeleteTail()
	L.DeleteTail()

	L.Traverse(f)
	t.Errorf("////////////////////////////")

	t.Errorf("L.Len() : %d", L.Len())
	f(L.Index(L.Len() - 1))

	L.AddTail(&n1.N)
	L.AddTail(&n2.N)
	L.AddTail(&n3.N)
	L.AddTail(&n4.N)
	t.Errorf("L.Len() : %d", L.Len())
	err := L.Add(&n1.N, &n4.N)
	if err != nil {
		t.Errorf("L.Add() : err:%v", err)
	}

	indexs := L.IndexRange(0, L.Len()-1)

	for _, i := range indexs {
		f(i)
	}

	t.Errorf("////////////////////////////")
	n5 := NewNode("5555")
	L.AddHead(&n5.N)
	L.Traverse(f)

	t.Errorf("////////////////////////////")
	L.Delete(&n5.N)
	L.Traverse(f)

	t.Errorf("////////////////////////////")
	L.Delete(&n3.N)
	L.Traverse(f)

	t.Errorf("////////////////////////////")
	L.Delete(&n4.N)
	L.Traverse(f)

	t.Errorf("////////////////////////////")
	L.Delete(&n4.N)
	L.Traverse(f)

	t.Errorf("////////////////////////////")
	indexs = L.IndexRange(1, 8)
	for _, i := range indexs {
		f(i)
	}

	t.Errorf("L.Len() : %d", L.Len())
}
