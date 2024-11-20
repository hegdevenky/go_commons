package list

import (
	"bytes"
	"fmt"
	"strings"
	"text/tabwriter"
)

// ImmutableNode is an interface that defines method for a Node in Linked list of given type
//
// Value method returns the data stored in the linked list node.
// If no data is present, it returns zero value of the type
//
// Next method returns a pointer to the next data node in the Linked list,
// if there is no next node / data it returns nil.
//
// Prev method returns a pointer to the previous data node in the Linked list,
// if there is no preceding node / data it returns nil.
type ImmutableNode[T any] interface {
	Value() T
	Next() ImmutableNode[T]
	Prev() ImmutableNode[T]
}

// singlyLinkedNode defines properties for a node in Singly linked list
type singlyLinkedNode[T any] struct {
	value T
	next  *singlyLinkedNode[T] // pointer to next
}

func (s *singlyLinkedNode[T]) Value() T {
	return s.value
}

func (s *singlyLinkedNode[T]) Next() ImmutableNode[T] {
	if s == nil || s.next == nil {
		return nil
	}
	return s.next
}

func (s *singlyLinkedNode[T]) Prev() ImmutableNode[T] {
	return nil
}

func (s *singlyLinkedNode[T]) String() string {
	if s == nil {
		return "<nil>"
	}
	var sb strings.Builder
	if s.next != nil {
		sb.WriteString(fmt.Sprintf("*[%v|—]⃓——→ [%v|—]⃓", s.value, s.next.value))
	} else {
		sb.WriteString(fmt.Sprintf("*[%v|—]⃓——→ [%v|—]⃓", s.value, s.next))
	}
	return sb.String()
}

// doublyLinkedNode defines properties for a node in Doubly linked list
type doublyLinkedNode[T any] struct {
	value      T
	next, prev *doublyLinkedNode[T] // pointer to next and prev
}

func (dn *doublyLinkedNode[T]) Value() T {
	return dn.value
}

func (dn *doublyLinkedNode[T]) Next() ImmutableNode[T] {
	if dn.next == nil {
		return nil
	}
	return dn.next
}

func (dn *doublyLinkedNode[T]) Prev() ImmutableNode[T] {
	if dn.prev == nil {
		return nil
	}
	return dn.prev
}

func (dn *doublyLinkedNode[T]) String() string {
	if dn == nil {
		return "<nil>"
	}

	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 0, 0, 0, ' ', tabwriter.Debug)
	switch {
	case dn.prev == nil && dn.next != nil:
		fmt.Fprintf(w, "%v<——\t%v\t——>\t%v\t ", "<nil>", dn.value, dn.next.value)
	case dn.prev != nil && dn.next == nil:
		fmt.Fprintf(w, "\t%v\t<——\t%v\t——>%v ", dn.prev.value, dn.value, "<nil>")
	default:
		fmt.Fprintf(w, "\t%v\t<——\t%v\t——>\t%v\t ", dn.prev.value, dn.value, dn.next.value)
	}
	w.Flush()
	return buf.String()
}

// Singly linked and Doubly linked Nodes

// AsSinglyLinkedNodes is a constructor function, that returns reference to SinglyLinkedNode
func AsSinglyLinkedNodes[T any](values ...T) (head *SinglyLinkedNode[T]) {
	head = &SinglyLinkedNode[T]{}
	cur := head
	for i, v := range values {
		cur.Value = v
		if i != len(values)-1 {
			cur.Next = &SinglyLinkedNode[T]{}
			cur = cur.Next
		}
	}
	cur = nil // avoid memory leak
	return
}

// SinglyLinkedNode represents a Node in SinglyLinkedList
// This can be used to construct SinglyLinkedList
// All properties are exported for quick modification.
// [AsSinglyLinkedNodes] can be used create singly linked nodes.
type SinglyLinkedNode[T any] struct {
	Value T
	Next  *SinglyLinkedNode[T]
}

// Len method returns number of connected nodes to the invoking node.
// Returned length is including the current node.
func (s *SinglyLinkedNode[T]) Len() (l int) {
	cur := s
	for cur != nil {
		l, cur = l+1, cur.Next
	}
	cur = nil // avoid memory leak
	return
}

func (s *SinglyLinkedNode[T]) String() string {
	if s == nil {
		return "<nil>"
	}
	var sb strings.Builder
	cur := s
	for cur != nil {
		sb.WriteString(fmt.Sprintf("[%v|—]——→", cur.Value))
		cur = cur.Next
	}
	sb.WriteString("<nil>")
	return sb.String()
}

func AsDoubleLinkedNodes[T any](values ...T) (head *DoublyLinkedNode[T]) {
	var prev *DoublyLinkedNode[T] = nil
	head = &DoublyLinkedNode[T]{}
	cur := head

	for i, v := range values {
		cur.Value, cur.Prev = v, prev
		if i != len(values)-1 {
			cur.Next = &DoublyLinkedNode[T]{}
			prev = cur
			cur = cur.Next
		}
	}
	cur = nil
	return head
}

// DoublyLinkedNode represents a Node in DoublyLinkedList
// This can be used to construct DoublyLinkedList
// All properties are exported for quick modification.
// [AsDoubleLinkedNodes] can be used create singly linked nodes.
type DoublyLinkedNode[T any] struct {
	Value      T
	Next, Prev *DoublyLinkedNode[T]
}

// Len method returns number of nodes starting from current node (d).
// Returns 0 if the current node is nil
// Returns 1 if there are no next nodes linked current node
func (d *DoublyLinkedNode[T]) Len() (l int) {
	cur := d
	for cur != nil {
		l, cur = l+1, cur.Next
	}
	return
}

// String is a Stringer method returns string representation of DoublyLinkedNodes.
func (d *DoublyLinkedNode[T]) String() string {
	if d == nil {
		return "<nil>"
	}
	var sb strings.Builder
	sb.WriteString("<nil>←——")
	for l, count, cur := d.Len(), 0, d; cur != nil; count, cur = count+1, cur.Next {
		if count < l-1 {
			sb.WriteString(fmt.Sprintf("[—|%v|—]←——→", cur.Value))
		} else {
			sb.WriteString(fmt.Sprintf("[—|%v|—]", cur.Value))
		}
	}
	sb.WriteString("——→<nil>")
	return sb.String()
}
