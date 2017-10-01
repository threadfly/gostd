package list

import (
	"fmt"
	"reflect"
)

type List struct {
	num  int
	next *ListNode
}

type ListNode struct {
	next *ListNode
	pre  *ListNode
	V    interface{}
}

func NewList() *List {
	return &List{
		num:  0,
		next: &ListNode{nil, nil, nil},
	}
}

func (this *List) Len() int {
	return this.num
}

func (this *List) Head() *ListNode {
	if this.next != nil {
		return this.next
	}

	return nil
}

func (this *List) Tail() *ListNode {
	if this.num == 0 {
		return nil
	}

	n := this.Head()

	return n.pre
}

func (this *List) Find(v interface{}) bool {
	if this.num == 0 {
		return false
	}

	h := this.Head()
	n := h
	for h != nil {
		if reflect.DeepEqual(n.V, v) {
			return true
		}
		n = n.next
		if h == n {
			return false
		}
	}

	return false
}

func (this *List) insertAfter(base *ListNode, new *ListNode) {
	new.next = base.next
	new.pre = base

	base.next.pre = new
	base.next = new
	this.num++
}

/*
 *  Insert new node to place that after every node
 */
func (this *List) Add(base *ListNode, new *ListNode) error {
	if !(this.Find(base.V)) {
		return fmt.Errorf("Have no Found base Node")
	}

	if this.Find(new.V) {
		this.delete(new)
	}

	this.insertAfter(base, new)
	return nil
}

/*
 *	Add new to List's Tail
 */
func (this *List) AddTail(new *ListNode) {
	last := this.Tail()
	if last == nil {
		new.pre, new.next = new, new
		this.next = new
		this.num++
	} else {
		this.insertAfter(last, new)
	}
}

func (this *List) AddHead(new *ListNode) {
	this.AddTail(new)
	this.next = new
}

func (this *List) delete(n *ListNode) {
	if this.num == 0 || n == nil {
		return
	}

	if this.num == 1 {
		n.next, n.pre = nil, nil
		this.next = nil
		this.num--
		return
	}

	if this.Head() == n {
		this.next = n.next
	}

	n.pre.next = n.next
	n.next.pre = n.pre
	this.num--
}

func (this *List) DeleteHead() {
	this.delete(this.Head())
}

func (this *List) DeleteTail() {
	this.delete(this.Tail())
}

func (this *List) Delete(n *ListNode) bool {
	f := this.Find(n.V)
	if !f {
		return false
	}

	this.delete(n)
	return true
}

func (this *List) Traverse(f func(v interface{})) {
	if f == nil {
		return
	}
	n := this.Head()

	for n != nil {
		f(n.V)
		n = n.next
		if n == this.Head() {
			break
		}
	}
}

func (this *List) Clear() {
	n := this.Head()
	tail := this.Head()

	if n == nil {
		return
	}

	for {
		last := n.next
		n.pre = nil
		n.next = nil
		n = last
		if n == tail {
			last.next = nil
			last.pre = nil
			break
		}
	}

	this.next = nil
}

/*
*	 i from 0 to len(list)-1
 */
func (this *List) Index(i int) interface{} {
	n := this.Head()
	for ; n != nil; i-- {
		if i == 0 {
			return n.V
		}
		n = n.next

		if n == this.Head() {
			break
		}
	}

	return nil
}

func (this *List) IndexRange(b, e int) []interface{} {
	if e < b {
		return nil
	}

	n := this.Head()
	l := e - b
	if l > this.num {
		l = this.num
	}

	result := make([]interface{}, 0, l)
	for n != nil {
		if b == 0 {
			result = append(result, n.V)
		} else {
			b--
		}

		if e == 0 {
			break
		} else {
			e--
		}
		n = n.next
		if n == this.Head() {
			break
		}
	}

	return result
}
