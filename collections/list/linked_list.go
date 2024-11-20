package list

import (
	"fmt"
	"iter"
)

type linkedListType string

func (lt linkedListType) String() string {
	return string(lt)
}

const (
	SinglyLinked linkedListType = "SINGLY_LINKED"
	DoublyLinked linkedListType = "DOUBLY_LINKED"
	Circular     linkedListType = "CIRCULAR"
)

// The LinkedList interface defines the functions for the LinkedList Abstract Data Type (ADT).
// Any type that implements this interface can function as a LinkedList.
// [SinglyLinkedList], [DoublyLinkedList] and [CircularLinkedList] are the known implementations.
type LinkedList[T any] interface {
	// The AddLast method appends the given element to the end of the LinkedList.
	//The new element becomes the tail of the list.
	//The method returns the same LinkedList instance to support method chaining.
	// Ex:
	//	sList := NewSinglyLinkedList[string]()
	//	sList.AddLast("hello").AddLast("world").AddFirst("Greeting:")
	// "hello" -> "world" -> "Greeting:"
	AddLast(e T) LinkedList[T]

	// The AddFirst method adds a new element to the beginning of this LinkedList.
	// As a result, the new element becomes the head of the list.
	// This method returns the same LinkedList instance to allow method chaining.
	// Ex:
	//	sList := NewSinglyLinkedList[string]()
	//	sList.AddFirst("hello").AddFirst("Greeting:").AddLast("world")
	// "Greeting:" -> "hello" -> "world"
	AddFirst(e T) LinkedList[T]

	// The Insert method inserts an element into the LinkedList at a specified position.
	// It accepts the element to be inserted and a zero-based index.
	// This method returns a boolean indicating whether the insertion was successful and [ErrIndexOutOfBounds]
	// error is returned if the index value is less than 0 or greater than the length of the LinkedList.
	Insert(e T, index int) (bool, error)

	// The GetFirst method returns the first element, or the head of the list.
	// If the list is nil or empty, it returns [ErrNoSuchElement] error.
	// In such cases, the first return value will be the zero value of the specified type T.
	GetFirst() (T, error)

	// The GetLast method returns the last element, or the tail of the list.
	// If the list is nil or empty, it returns [ErrNoSuchElement] error.
	//In this case, the first return value will be the zero value of the specified type T.
	GetLast() (T, error)

	// The Get method retrieves an element at the specified index from the LinkedList.
	// A zero-based index is used. [ErrIndexOutOfBounds] error is returned
	// if the index value is less than 0 or greater than the length of the LinkedList.
	Get(index int) (T, error)

	// The GetHeadNode method returns the first (head) node in the LinkedList.
	// The returned node is of type [ImmutableNode] to prevent modifications
	// to its value or its references to the next and previous nodes.
	// If the LinkedList is nil or empty, the method returns an [ErrNoSuchElement] error.
	GetHeadNode() (ImmutableNode[T], error)

	// The GetTailNode method returns the last (tail) node in the LinkedList.
	// The returned node is of type [ImmutableNode] to prevent modifications
	// to its value or its references to the next and previous nodes.
	// If the LinkedList is nil or empty, the method returns an [ErrNoSuchElement] error.
	GetTailNode() (ImmutableNode[T], error)

	// The Len method returns the number of elements in the LinkedList.
	// If the list is nil or this method is invoked on a nil reference, it will cause a panic.
	Len() int

	// The IsEmpty method returns a boolean indicating whether this LinkedList is empty.
	// It returns true if the list is empty and false otherwise.
	IsEmpty() bool

	// The RemoveFirst method deletes the first (leftmost) element in the LinkedList and returns the removed element.
	// It returns an [ErrNoSuchElement] error if the list is empty.
	RemoveFirst() (T, error)

	// The RemoveLast method deletes the last (rightmost) element in the LinkedList and returns the removed element.
	// If the list is empty, it returns an [ErrNoSuchElement] error.
	RemoveLast() (T, error)

	// The RemoveAt method deletes an element at a specified index and returns the removed element.
	// A zero-based index is used. If the index is invalid, an [ErrIndexOutOfBounds] error is returned.
	RemoveAt(index int) (T, error)

	// The All method returns an iterator for the [LinkedList].
	// This can be used with a for-range loop.
	// When using this with a for-range, the first return value is the zero-based index of the element,
	// and the second value is the data at that index.
	//All method returns an Iterator to the [LinkedList]
	// todo: Add example
	All() iter.Seq2[int, T]

	// The Values method returns an iterator for the values in the LinkedList.
	// This can also be used with a for-range loop.
	// Unlike [All], this method only returns the value, not the index.
	// todo: Add example
	Values() iter.Seq[T]

	// The ReverseAll method returns an iterator for the LinkedList that iterates through the list in reverse order.
	// This can be used with a for-range loop.
	// Todo: Add example
	ReverseAll() iter.Seq2[int, T]

	// The ToSlice method is a convenient way to convert a [LinkedList] of type T into a slice of type T.
	// The length of the returned slice will be equal to the length of the [LinkedList].
	// Element positions are preserved in the returned slice; therefore, linkedList.get(i) == returnedSlice[i].
	//
	//If the LinkedList is nil, this method returns nil. If the LinkedList is empty, it returns an empty slice of type T.
	ToSlice() []T

	fmt.Stringer
}

// CONSTRUCTORS / Factory functions.

// The NewLinkedList function is a factory function that returns a reference to a newly created [LinkedList]
// implementation based on the specified linkedListType. The accepted values for [linkedListType] are [SinglyLinked],
// [DoublyLinked], or [Circular]
// - If the value for linkedListType is [SinglyLinked], a reference to a newly created [SinglyLinkedList] is returned.
// - If the value for linkedListType is [DoublyLinked], a reference to a newly created [DoublyLinkedList] is returned.
// - If the value for linkedListType is [Circular], a reference to a newly created [CircularLinkedList] is returned.
//
// If an invalid value is passed for linkedListType, the input is ignored,
// and a [DoublyLinkedList] implementation is returned instead of throwing an error.
// This approach is adopted to favor method chaining on the returned [LinkedList].
func NewLinkedList[T any](linkedListType linkedListType) (linkedList LinkedList[T]) {
	switch linkedListType {
	case SinglyLinked:
		linkedList = &SinglyLinkedList[T]{}
	case Circular:
		linkedList = &CircularLinkedList[T]{}
	default:
		linkedList = &DoublyLinkedList[T]{}
	}
	return
}

// The NewLinkedListFromSlice function is a factory function that returns a reference to a newly created [LinkedList]
// implementation based on the specified linkedListType. The accepted values for [linkedListType] are [SinglyLinked],
// [DoublyLinked], or [Circular]. As in the name, this method accepts a slice of type T along with linkedListType.
// - If the value for linkedListType is [SinglyLinked], a reference to a newly created [SinglyLinkedList] is returned.
// - If the value for linkedListType is [DoublyLinked], a reference to a newly created [DoublyLinkedList] is returned.
// - If the value for linkedListType is [Circular], a reference to a newly created [CircularLinkedList] is returned.
//
// If an invalid value is passed for linkedListType, the input is ignored,
// and a [DoublyLinkedList] implementation is returned instead of throwing an error.
// This approach is adopted to favor method chaining on the returned [LinkedList].
func NewLinkedListFromSlice[T any](linkedListType linkedListType, slice []T) (linkedList LinkedList[T]) {
	switch linkedListType {
	case SinglyLinked:
		linkedList = &SinglyLinkedList[T]{}
	case Circular:
		linkedList = &CircularLinkedList[T]{}
	default:
		linkedList = &DoublyLinkedList[T]{}
	}
	for _, v := range slice {
		linkedList.AddLast(v)
	}
	return
}

// NewLinkedListFrom is a factory function to create an Implementation of [LinkedList] from one or elements.
// It is just a convenient wrapper function over [NewLinkedListFromSlice].
func NewLinkedListFrom[T any](linkedListType linkedListType, elements ...T) (linkedList LinkedList[T]) {
	return NewLinkedListFromSlice[T](linkedListType, elements)
}
