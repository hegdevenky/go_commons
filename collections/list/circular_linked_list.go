package list

import (
	"fmt"
	"iter"
	"strings"
)

type CircularLinkedList[T any] struct {
	head, tail *doublyLinkedNode[T]
	len        int
}

func (c *CircularLinkedList[T]) AddLast(e T) LinkedList[T] {
	newNode := &doublyLinkedNode[T]{value: e}
	// if this is the first item
	if c.IsEmpty() {
		c.tail = newNode
		c.head = c.tail
		c.len, newNode = c.len+1, nil
		return c
	}
	newNode.next, newNode.prev = c.head, c.tail
	c.tail.next = newNode
	c.head.prev = newNode
	c.tail = newNode
	c.len, newNode = c.len+1, nil
	return c
}

func (c *CircularLinkedList[T]) AddFirst(e T) LinkedList[T] {
	newNode := &doublyLinkedNode[T]{value: e}
	// if this is the first item
	if c.IsEmpty() {
		c.head = newNode
		c.tail = c.head
		c.len, newNode = c.len+1, nil
		return c
	}
	newNode.next, newNode.prev = c.head, c.tail
	c.head.prev = newNode
	c.tail.next = newNode
	c.head = newNode
	c.len, newNode = c.len+1, nil
	return c
}

func (c *CircularLinkedList[T]) Insert(e T, index int) (bool, error) {
	switch {
	case index < 0 || index > c.Len():
		return false, errIndexOutOfBounds(index, c.Len())
	case index == 0:
		c.AddFirst(e)
		return true, nil
	case index == c.Len():
		c.AddLast(e)
		return true, nil
	case index == 1:
		newNode := &doublyLinkedNode[T]{value: e}
		newNode.next, newNode.prev = c.head.next, c.head
		c.head.next.prev = newNode
		c.head.next = newNode
		c.len, newNode = c.len+1, nil
		return true, nil
	case index == 2:
		newNode := &doublyLinkedNode[T]{value: e}
		newNode.next, newNode.prev = c.head.next.next, c.head.next
		c.head.next.next.prev = newNode
		c.head.next.next = newNode
		c.len, newNode = c.len+1, nil
		return true, nil
	case index == c.Len()-1:
		newNode := &doublyLinkedNode[T]{value: e}
		newNode.next, newNode.prev = c.tail, c.tail.prev
		c.tail.prev.next = newNode
		c.tail.prev = newNode
		c.len, newNode = c.len+1, nil
		return true, nil
	case index == c.Len()-2:
		newNode := &doublyLinkedNode[T]{value: e}
		newNode.next, newNode.prev = c.tail.prev, c.tail.prev.prev
		c.tail.prev.prev.next = newNode
		c.tail.prev.prev = newNode
		c.len, newNode = c.len+1, nil
		return true, nil
	default:
		newNode := &doublyLinkedNode[T]{value: e}
		cur := c.head.next.next
		if index < c.Len()/2 {
			// position is left of mid, start from head.
			for i := 2; cur != nil && i < index; i++ {
				cur = cur.next
			}
		} else {
			// index is >= mid, so position is mid, or it's right, start from the tail
			cur := c.tail.prev.prev
			for i := c.Len() - 3; cur != nil && i > index; i-- {
				cur = cur.prev
			}
		}
		newNode.next, newNode.prev = cur, cur.prev
		cur.prev.next = newNode
		cur.prev, c.len = newNode, c.len+1

		newNode, cur = nil, nil
		return true, nil
	}
}

func (c *CircularLinkedList[T]) GetFirst() (T, error) {
	if c.IsEmpty() {
		var t T
		return t, errNoSuchElement()
	}
	return c.head.value, nil
}

func (c *CircularLinkedList[T]) GetLast() (T, error) {
	if c.IsEmpty() {
		var t T
		return t, errNoSuchElement()
	}
	return c.tail.value, nil
}

func (c *CircularLinkedList[T]) Get(index int) (T, error) {
	switch {
	case c.IsEmpty() || index < 0 || index >= c.Len():
		var zero T
		return zero, errIndexOutOfBounds(index, c.Len())
	case index == 0:
		return c.GetFirst()
	case index == c.Len()-1:
		return c.GetLast()
	case index == 1:
		return c.head.next.value, nil
	case index == c.Len()-2:
		return c.tail.prev.value, nil
	default:
		cur := c.tail.prev.prev
		if index > (c.Len() / 2) {
			// index right side of mid, so start from tail
			for i := c.Len() - 3; cur != nil && i > index; i-- {
				cur = cur.prev
			}
		} else {
			// index left of mid, so start from head
			cur = c.head.next.next
			for i := 2; cur != nil && i < index; i++ {
				cur = cur.next
			}
		}
		value := cur.value
		cur = nil // avoid memory leaks
		return value, nil
	}
}

func (c *CircularLinkedList[T]) GetHeadNode() (ImmutableNode[T], error) {
	if c.Len() == 0 {
		return nil, errNoSuchElement()
	}
	return c.head, nil
}

func (c *CircularLinkedList[T]) GetTailNode() (ImmutableNode[T], error) {
	if c.Len() == 0 {
		return nil, errNoSuchElement()
	}
	return c.tail, nil
}

func (c *CircularLinkedList[T]) Len() int {
	return c.len
}

func (c *CircularLinkedList[T]) IsEmpty() bool {
	return c.Len() == 0
}

func (c *CircularLinkedList[T]) RemoveFirst() (T, error) {
	if c.Len() <= 1 {
		return c.removeZeroOrOne()
	}
	rm := c.head

	c.head = rm.next
	c.head.prev = c.tail
	c.tail.next = c.head

	rm.next, rm.prev = nil, nil // avoid memory leaks
	ret := rm.value
	c.len, rm = c.len-1, nil
	return ret, nil
}

func (c *CircularLinkedList[T]) RemoveLast() (T, error) {
	if c.Len() <= 1 {
		return c.removeZeroOrOne()
	}
	rm := c.tail
	c.tail = rm.prev
	c.head.prev = rm.prev
	rm.prev.next = c.head
	ret := rm.value
	// avoid memory leaks
	rm.next, rm.prev = nil, nil
	c.len, rm = c.len-1, nil
	return ret, nil
}

func (c *CircularLinkedList[T]) RemoveAt(index int) (T, error) {
	switch {
	case c.IsEmpty() || index < 0 || index >= c.Len():
		var zero T
		return zero, errIndexOutOfBounds(index, c.Len())
	case index == 0:
		return c.RemoveFirst()
	case index == c.Len()-1:
		return c.RemoveLast()
	case index == 1:
		removed, rv := c.head.next, c.head.next.value
		c.head.next, removed.next.prev = removed.next, c.head
		c.len, removed.next, removed.prev, removed = c.len-1, nil, nil, nil
		return rv, nil
	case index == c.Len()-2:
		removed, rv := c.tail.prev, c.tail.prev.value
		c.tail.prev, removed.prev.next = removed.prev, c.tail
		c.len, removed.next, removed.prev, removed = c.len-1, nil, nil, nil
		return rv, nil
	default:
		cur := c.tail.prev.prev
		if index > (c.Len() / 2) {
			// index right side of mid, so start from tail
			for i := c.Len() - 3; cur != nil && i > index; i-- {
				cur = cur.prev
			}
		} else {
			// index left of mid, so start from head
			cur = c.head.next.next
			for i := 2; cur != nil && i < index; i++ {
				cur = cur.next
			}
		}
		rv := cur.value
		cur.prev.next, cur.next.prev = cur.next, cur.prev
		cur.next, cur.prev, cur, c.len = nil, nil, nil, c.len-1
		return rv, nil
	}
}

func (c *CircularLinkedList[T]) All() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i, cur := 0, c.head; i < c.Len(); i, cur = i+1, cur.next {
			if !yield(i, cur.value) {
				return
			}
		}
	}
}

func (c *CircularLinkedList[T]) Values() iter.Seq[T] {
	return func(yield func(T) bool) {
		for i, cur := 0, c.head; i < c.Len(); i, cur = i+1, cur.next {
			if !yield(cur.value) {
				return
			}
		}
	}
}

func (c *CircularLinkedList[T]) ReverseAll() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i, cur := c.Len()-1, c.tail; i >= 0; i, cur = i-1, cur.prev {
			if !yield(i, cur.value) {
				return
			}
		}
	}
}

func (c *CircularLinkedList[T]) ToSlice() []T {
	switch {
	case c == nil:
		return nil
	case c.IsEmpty():
		return make([]T, 0)
	default:
		slice := make([]T, c.Len())
		for i, v := range c.All() {
			slice[i] = v
		}
		return slice
	}
}

func (c *CircularLinkedList[T]) String() string {
	if c == nil {
		return "nil"
	}
	var sb strings.Builder
	for i, v := range c.All() {
		sb.WriteString(fmt.Sprintf("%v", v))
		if i+1 < c.Len() {
			sb.WriteString(" <=> ")
		}
	}
	return sb.String()
}

func (c *CircularLinkedList[T]) removeZeroOrOne() (T, error) {
	if c.IsEmpty() {
		var zero T
		return zero, errNoSuchElement()
	}
	value := c.head.value
	c.len, c.head, c.tail = c.len-1, nil, nil
	return value, nil
}
