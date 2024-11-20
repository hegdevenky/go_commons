package list

import (
	"fmt"
	"iter"
	"strings"
)

// A SinglyLinkedList is a type of linked list where each node in the linked list contains
// data / actual value and pointer (reference) to next data node.
//
// Linked list also has len property to keep track of its length / size.
// SinglyLinkedList implements list.LinkedList interface.
type SinglyLinkedList[T any] struct {
	head, tail *singlyLinkedNode[T]
	len        int
}

// The AddLast method appends the given element to the end of the LinkedList.
// The new element becomes the tail of the list.
// The method returns the same LinkedList instance to support method chaining.
// Ex:
//
//	sList := NewSinglyLinkedList[string]()
//	sList.AddLast("hello").AddLast("world").AddFirst("Greeting:")
//
// "hello" -> "world" -> "Greeting:"
func (s *SinglyLinkedList[T]) AddLast(e T) LinkedList[T] {
	if s.IsEmpty() {
		s.addFirstNode(e)
		return s
	}
	newNode := &singlyLinkedNode[T]{value: e}
	s.tail.next = newNode
	s.tail = newNode
	s.len, newNode = s.len+1, nil
	return s
}

func (s *SinglyLinkedList[T]) AddFirst(e T) LinkedList[T] {
	// check if list is empty, if so this will be the first element
	if s.IsEmpty() {
		s.addFirstNode(e)
		return s
	}
	newNode := &singlyLinkedNode[T]{e, s.head}
	// make the new node as head
	s.head = newNode
	s.len, newNode = s.len+1, nil
	return s
}

func (s *SinglyLinkedList[T]) Insert(e T, index int) (bool, error) {
	switch {
	case index < 0 || index > s.Len():
		return false, errIndexOutOfBounds(index, s.Len())
	case index == 0:
		s.AddFirst(e)
		return true, nil
	case index == s.Len():
		s.AddLast(e)
		return true, nil
	case index == 1:
		newNode := &singlyLinkedNode[T]{e, s.head.next}
		s.head.next, s.len = newNode, s.len+1
		newNode = nil // avoid memory leak
		return true, nil
	default:
		prev, cur := s.head.next, s.head.next.next
		for i := 2; cur != nil; {
			if i == index {
				break
			}
			i++
			prev = cur
			cur = cur.next
		}
		newNode := &singlyLinkedNode[T]{e, cur}
		prev.next, s.len = newNode, s.len+1
		return true, nil
	}
}

func (s *SinglyLinkedList[T]) GetFirst() (T, error) {
	if s.Len() == 0 {
		var zero T
		return zero, errNoSuchElement()
	}
	return s.head.value, nil
}

func (s *SinglyLinkedList[T]) GetLast() (T, error) {
	if s.Len() == 0 {
		var zero T
		return zero, errNoSuchElement()
	}
	return s.tail.value, nil
}

func (s *SinglyLinkedList[T]) Get(index int) (T, error) {
	switch {
	case s.IsEmpty() || index < 0 || index >= s.Len(): // if given index is out of bound.
		var t T
		return t, errIndexOutOfBounds(index, s.Len())
	case index == 0: // if index is 0, i.e first element
		return s.GetFirst()
	case index == s.Len()-1:
		return s.GetLast()
	case index == 1:
		return s.head.next.value, nil
	default:
		cur := s.head.next
		defer func() {
			cur = nil
		}()
		i := 0
		for cur != nil && cur.next != nil {
			if i+2 == index {
				return cur.next.value, nil
			}
			cur = cur.next
			i++
		}
		var t T
		return t, nil
	}
}

func (s *SinglyLinkedList[T]) GetHeadNode() (ImmutableNode[T], error) {
	if s == nil || s.Len() == 0 {
		return nil, errNoSuchElement()
	}
	return s.head, nil
}

func (s *SinglyLinkedList[T]) GetTailNode() (ImmutableNode[T], error) {
	if s == nil || s.Len() == 0 {
		return nil, errNoSuchElement()
	}
	return s.tail, nil
}

func (s *SinglyLinkedList[T]) Len() int {
	return s.len
}

func (s *SinglyLinkedList[T]) IsEmpty() bool {
	return s.Len() == 0
}

func (s *SinglyLinkedList[T]) addFirstNode(e T) {
	node := &singlyLinkedNode[T]{value: e}
	s.head = node
	s.tail = s.head
	s.len, node = s.len+1, nil
}

func (s *SinglyLinkedList[T]) RemoveFirst() (T, error) {
	// item can't be removed if list is empty
	if s.IsEmpty() {
		var t T
		return t, errNoSuchElement()
	}
	head := s.head
	s.head = head.next
	ret := head.value
	s.len, head = s.len-1, nil
	return ret, nil
}

func (s *SinglyLinkedList[T]) RemoveLast() (T, error) {
	if s.IsEmpty() {
		var t T
		return t, errNoSuchElement()
	}
	cur := s.head
	prev := cur
	for cur != nil && cur.next != nil {
		prev = cur
		cur = cur.next
	}
	s.tail = prev
	s.tail.next = nil // tail's next is nil
	ret := cur.value
	s.len, cur = s.len-1, nil
	return ret, nil
}

func (s *SinglyLinkedList[T]) RemoveAt(index int) (T, error) {
	switch {
	case s.IsEmpty() || index < 0 || index >= s.Len():
		var zero T
		return zero, errIndexOutOfBounds(index, s.Len())
	case index == 0:
		return s.RemoveFirst()
	case index == s.Len()-1:
		return s.RemoveLast()
	default:
		var prev *singlyLinkedNode[T] = nil
		i, cur := 0, s.head
		for ; cur != nil && i < index; i++ {
			prev = cur
			cur = cur.next
		}
		rv := cur.value
		prev.next = cur.next
		cur.next, cur, s.len = nil, nil, s.len-1
		return rv, nil

	}
}

func (s *SinglyLinkedList[T]) All() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		cur := s.head
		for i := 0; cur != nil; i++ {
			if !yield(i, cur.value) {
				return
			}
			cur = cur.next
		}
	}
}

func (s *SinglyLinkedList[T]) Values() iter.Seq[T] {
	return func(yield func(T) bool) {
		cur := s.head
		for cur != nil {
			if !yield(cur.value) {
				return
			}
			cur = cur.next
		}
	}
}

func (s *SinglyLinkedList[T]) ReverseAll() iter.Seq2[int, T] {
	sl := s.ToSlice()
	return func(yield func(int, T) bool) {
		for i := len(sl) - 1; i >= 0; i-- {
			if !yield(i, sl[i]) {
				return
			}
		}
	}
}

func (s *SinglyLinkedList[T]) ToSlice() []T {
	switch {
	case s == nil:
		return nil
	case s.IsEmpty():
		return make([]T, 0)
	default:
		slice := make([]T, s.Len())
		for i, v := range s.All() {
			slice[i] = v
		}
		return slice
	}
}

func (s *SinglyLinkedList[T]) String() string {
	if s == nil {
		return "<nil>"
	}
	var sb strings.Builder
	for _, v := range s.All() {
		sb.WriteString(fmt.Sprintf("[%v|—]⃓——→ ", v))
	}
	sb.WriteString("<nil>")
	return sb.String()
}
