package list

import (
	"fmt"
	"iter"
	"strings"
)

// A DoublyLinkedList is a type of linked list where each node in the linked list contains
// data / actual value and pointer (reference) to next and previous data nodes.
//
// Linked list also has len property to keep track of its length / size.
// DoublyLinkedList implements list.LinkedList interface.
type DoublyLinkedList[T any] struct {
	head, tail *doublyLinkedNode[T]
	len        int
}

func (d *DoublyLinkedList[T]) AddLast(e T) LinkedList[T] {
	newNode := &doublyLinkedNode[T]{value: e}
	if d.IsEmpty() {
		d.tail = newNode
		d.head = d.tail
		d.len, newNode = d.len+1, nil
		return d
	}
	d.tail.next = newNode
	newNode.prev = d.tail
	d.tail = newNode
	d.len, newNode = d.len+1, nil
	return d
}

func (d *DoublyLinkedList[T]) AddFirst(e T) LinkedList[T] {
	newNode := &doublyLinkedNode[T]{value: e}
	// if this is the first item
	if d.IsEmpty() {
		d.head = newNode
		d.tail = d.head
		d.len, newNode = d.len+1, nil
		return d
	}
	newNode.next = d.head
	d.head.prev = newNode
	d.head = newNode
	d.len, newNode = d.len+1, nil
	return d
}

func (d *DoublyLinkedList[T]) Insert(e T, index int) (bool, error) {
	switch {
	case index < 0 || index > d.Len():
		return false, errIndexOutOfBounds(index, d.Len())
	case index == 0:
		d.AddFirst(e)
		return true, nil
	case index == d.Len():
		d.AddLast(e)
		return true, nil
	case index == 1:
		newNode := &doublyLinkedNode[T]{value: e}
		newNode.next, newNode.prev = d.head.next, d.head
		d.head.next.prev = newNode
		d.head.next = newNode
		d.len, newNode = d.len+1, nil
		return true, nil
	case index == 2:
		newNode := &doublyLinkedNode[T]{value: e}
		newNode.next, newNode.prev = d.head.next.next, d.head.next
		d.head.next.next.prev = newNode
		d.head.next.next = newNode
		d.len, newNode = d.len+1, nil
		return true, nil
	case index == d.Len()-1:
		newNode := &doublyLinkedNode[T]{value: e}
		newNode.next, newNode.prev = d.tail, d.tail.prev
		d.tail.prev.next = newNode
		d.tail.prev = newNode
		d.len, newNode = d.len+1, nil
		return true, nil
	case index == d.Len()-2:
		newNode := &doublyLinkedNode[T]{value: e}
		newNode.next, newNode.prev = d.tail.prev, d.tail.prev.prev
		d.tail.prev.prev.next = newNode
		d.tail.prev.prev = newNode
		d.len, newNode = d.len+1, nil
		return true, nil
	default:
		newNode := &doublyLinkedNode[T]{value: e}
		cur := d.head.next.next
		if index < d.Len()/2 {
			for i := 2; cur != nil && i < index; i++ {
				cur = cur.next
			}
		} else {
			cur = d.tail.prev.prev
			for i := d.Len() - 3; cur != nil && i > index; i-- {
				cur = cur.prev
			}
		}
		newNode.next, newNode.prev = cur, cur.prev
		cur.prev.next = newNode
		cur.prev, d.len = newNode, d.len+1
		// avoid memory leaks
		newNode, cur = nil, nil
		return true, nil
	}
}

func (d *DoublyLinkedList[T]) GetFirst() (T, error) {
	if d.IsEmpty() {
		var t T
		return t, errNoSuchElement()
	}
	return d.head.value, nil
}

func (d *DoublyLinkedList[T]) GetLast() (T, error) {
	if d.IsEmpty() {
		var t T
		return t, errNoSuchElement()
	}
	return d.tail.value, nil
}

func (d *DoublyLinkedList[T]) Get(index int) (T, error) {
	switch {
	case d.IsEmpty() || index < 0 || index >= d.Len():
		var zero T
		return zero, errIndexOutOfBounds(index, d.Len())
	case index == 0:
		return d.GetFirst()
	case index == d.Len()-1:
		return d.GetLast()
	case index == 1:
		return d.head.next.value, nil
	case index == d.Len()-2:
		return d.tail.prev.value, nil
	default:
		cur := d.tail.prev.prev
		if index > (d.Len() / 2) {
			// index right side of mid, so start from tail
			for i := d.Len() - 3; cur != nil && i > index; i-- {
				cur = cur.prev
			}
		} else {
			// index left of mid, so start from head
			cur = d.head.next.next
			for i := 2; cur != nil && i < index; i++ {
				cur = cur.next
			}
		}
		value := cur.value
		cur = nil // avoid memory leaks
		return value, nil
	}
}

func (d *DoublyLinkedList[T]) GetHeadNode() (ImmutableNode[T], error) {
	if d == nil || d.Len() == 0 {
		return nil, errNoSuchElement()
	}
	return d.head, nil
}

func (d *DoublyLinkedList[T]) GetTailNode() (ImmutableNode[T], error) {
	if d == nil || d.Len() == 0 {
		return nil, errNoSuchElement()
	}
	return d.tail, nil
}

func (d *DoublyLinkedList[T]) Len() int {
	return d.len
}

func (d *DoublyLinkedList[T]) IsEmpty() bool {
	return d.Len() == 0
}

func (d *DoublyLinkedList[T]) RemoveFirst() (T, error) {
	if d.Len() <= 1 {
		return d.removeZeroOrOne()
	}
	rm := d.head
	d.head = rm.next
	d.head.prev = nil
	rm.next, rm.prev = nil, nil
	value := rm.value
	d.len, rm = d.len-1, nil
	return value, nil
}

func (d *DoublyLinkedList[T]) RemoveLast() (T, error) {
	if d.Len() <= 1 {
		return d.removeZeroOrOne()
	}
	rm := d.tail
	d.tail = rm.prev
	d.tail.next = nil
	rm.prev = nil
	value := rm.value
	d.len, rm = d.len-1, nil
	return value, nil
}

func (d *DoublyLinkedList[T]) RemoveAt(index int) (T, error) {
	switch {
	case d.IsEmpty() || index < 0 || index >= d.Len():
		var zero T
		return zero, errIndexOutOfBounds(index, d.Len())
	case index == 0:
		return d.RemoveFirst()
	case index == d.Len()-1:
		return d.RemoveLast()
	case index == 1:
		removed, rv := d.head.next, d.head.next.value
		d.head.next, removed.next.prev = removed.next, d.head
		d.len, removed.next, removed.prev, removed = d.len-1, nil, nil, nil
		return rv, nil
	case index == d.Len()-2:
		removed, rv := d.tail.prev, d.tail.prev.value
		d.tail.prev, removed.prev.next = removed.prev, d.tail
		d.len, removed.next, removed.prev, removed = d.len-1, nil, nil, nil
		return rv, nil
	default:
		cur := d.tail.prev.prev
		if index > (d.Len() / 2) {
			// index right side of mid, so start from tail
			for i := d.Len() - 3; cur != nil && i > index; i-- {
				cur = cur.prev
			}
		} else {
			// index left of mid, so start from head
			cur = d.head.next.next
			for i := 2; cur != nil && i < index; i++ {
				cur = cur.next
			}
		}
		rv := cur.value
		cur.prev.next, cur.next.prev = cur.next, cur.prev
		cur.next, cur.prev, cur, d.len = nil, nil, nil, d.len-1
		return rv, nil
	}
}

func (d *DoublyLinkedList[T]) All() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i, cur := 0, d.head; cur != nil; i, cur = i+1, cur.next {
			if !yield(i, cur.value) {
				return
			}
		}
	}
}

func (d *DoublyLinkedList[T]) Values() iter.Seq[T] {
	return func(yield func(T) bool) {
		for cur := d.head; cur != nil; cur = cur.next {
			if !yield(cur.value) {
				return
			}
		}
	}
}

func (d *DoublyLinkedList[T]) ReverseAll() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i, cur := d.Len()-1, d.tail; cur != nil && i >= 0; i, cur = i-1, cur.prev {
			if !yield(i, cur.value) {
				return
			}
		}
	}
}

func (d *DoublyLinkedList[T]) ToSlice() []T {
	switch {
	case d == nil:
		return nil
	case d.IsEmpty():
		return make([]T, 0)
	default:
		slice := make([]T, d.Len())
		for i, v := range d.All() {
			slice[i] = v
		}
		return slice

	}
}

func (d *DoublyLinkedList[T]) String() string {
	if d == nil {
		return "nil"
	}
	var sb strings.Builder
	for i, v := range d.All() {
		sb.WriteString(fmt.Sprintf("%v", v))
		if i+1 < d.Len() {
			sb.WriteString(" <=> ")
		}
	}
	return sb.String()
}

func (d *DoublyLinkedList[T]) removeZeroOrOne() (T, error) {
	if d.IsEmpty() {
		var zero T
		return zero, errNoSuchElement()
	}
	value := d.head.value
	d.len, d.head, d.tail = d.len-1, nil, nil
	return value, nil
}
