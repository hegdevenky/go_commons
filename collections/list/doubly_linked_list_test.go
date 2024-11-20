package list

import (
	"errors"
	"math/rand/v2"
	"reflect"
	"testing"
)

func TestNewDoublyLinkedList(t *testing.T) {
	dLinkedList := NewLinkedList[int](DoublyLinked)
	if dLinkedList == nil {
		t.Errorf("expected non nil Linkedlist but got %v\n", dLinkedList)
	}
	if _, ok := any(dLinkedList).(*DoublyLinkedList[int]); !ok {
		t.Errorf("expected %v type Linkedlist but got %T type\n", reflect.TypeFor[*DoublyLinkedList[int]](), dLinkedList)
	}
	if dLinkedList.Len() != 0 {
		t.Errorf("expected %d length but got %d", 0, dLinkedList.Len())
	}
	if !dLinkedList.IsEmpty() {
		t.Errorf("expected empty list but got %d", dLinkedList.Len())
	}
	t.Cleanup(func() {
		dLinkedList = nil
	})
}

func TestNewDoublyLinkedListFromSlice(t *testing.T) {
	// nil slice
	intList := NewLinkedListFromSlice[int](DoublyLinked, nil)
	if intList == nil {
		t.Errorf("got %v expected = %v", intList, &DoublyLinkedList[int]{})
	}

	// empty slice
	intList = NewLinkedListFromSlice[int](DoublyLinked, []int{})
	if intList == nil {
		t.Errorf("got nil expected zero value %v", &DoublyLinkedList[int]{})
	}
	if _, ok := any(intList).(*DoublyLinkedList[int]); !ok {
		t.Errorf("expected %v type Linkedlist but got %T type\n", reflect.TypeFor[*DoublyLinkedList[int]](), intList)
	}
	if intList.Len() != 0 {
		t.Errorf("got Len %d, expected Len %d", intList.Len(), 0)
	}
	if !intList.IsEmpty() {
		t.Errorf("IsEmpty: Got %t, expected %t", intList.IsEmpty(), true)
	}

	// slice with one element
	strList := NewLinkedListFromSlice[string](DoublyLinked, []string{"hello world"})
	if _, ok := any(strList).(*DoublyLinkedList[string]); !ok {
		t.Errorf("expected %v type Linkedlist but got %T type\n", reflect.TypeFor[*DoublyLinkedList[string]](), strList)
	}
	if strList.Len() != 1 {
		t.Errorf("got Len %d, expected Len %d", strList.Len(), 1)
	}
	if strList.IsEmpty() {
		t.Errorf("IsEmpty: Got %t, expected %t", strList.IsEmpty(), false)
	}
	// todo: test for head and tail

	// slice with multiple elements
	evens := []int{4, 8, 6, 2}
	intList = NewLinkedListFromSlice[int](DoublyLinked, evens)
	if _, ok := any(intList).(*DoublyLinkedList[int]); !ok {
		t.Errorf("expected %v type Linkedlist but got %T type\n", reflect.TypeFor[*DoublyLinkedList[int]](), intList)
	}
	if intList.Len() != len(evens) {
		t.Errorf("got Len %d, expected Len %d", intList.Len(), len(evens))
	}
	if intList.IsEmpty() {
		t.Errorf("IsEmpty: Got %t, expected %t", intList.IsEmpty(), false)
	}
	// todo: test for head and tail
}

func TestDoublyLinkedList_GetFirst(t *testing.T) {
	list0 := NewLinkedList[int](DoublyLinked)
	list1 := NewLinkedListFrom[int](DoublyLinked, 1234)
	list2 := NewLinkedListFrom[int](DoublyLinked, 13462, 76354192, 3645, 13262, 76354193, 43645, 4, 5, 8, 3, 14, 98, 3425, 56777, 24352)

	type testCase[T any] struct {
		name    string
		s       LinkedList[T]
		want    T
		wantErr bool
		err     error
	}

	tests := []testCase[int]{
		{"expect ErrNoSuchElement", list0, 0, true, ErrNoSuchElement},
		{"list with only element", list1, 1234, false, nil},
		{"list with more elements", list2, 13462, false, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetFirst()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFirst() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFirst() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoublyLinkedList_GetLast(t *testing.T) {
	list0 := NewLinkedList[int](DoublyLinked)
	list1 := NewLinkedListFrom[int](DoublyLinked, 1234)
	list2 := NewLinkedListFrom[int](DoublyLinked, 13462, 76354192, 3645, 13262, 76354193, 43645, 4, 5, 8, 3, 14, 98, 3425, 56777, 24352)

	type testCase[T any] struct {
		name    string
		s       LinkedList[T]
		want    T
		wantErr bool
		err     error
	}

	tests := []testCase[int]{
		{"expect ErrNoSuchElement", list0, 0, true, ErrNoSuchElement},
		{"list with only element", list1, 1234, false, nil},
		{"list with more elements", list2, 24352, false, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetLast()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFirst() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFirst() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoublyLinkedList_AddFirst(t *testing.T) {
	// add first on nil list
	// add first when no elements
	// add first when more than element

	type testCase[T any] struct {
		name           string
		list           LinkedList[T]
		input          T
		lengthExpected int
		headExpected   T
		tailExpected   T
	}
	list1 := NewLinkedList[int](DoublyLinked)

	tests := []testCase[int]{
		{"add first item into list", list1, 1, 1, 1, 1},
		{"add second item into list", list1, 4, 2, 4, 1},
		{"add third item into list", list1, 5, 3, 5, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.list.AddFirst(tt.input)
			// check len
			if l := list1.Len(); l != tt.lengthExpected {
				t.Errorf("test %s, list length: got %d, expected %d", tt.name, l, tt.lengthExpected)
			}
			// check head
			if h, err := list1.GetFirst(); err != nil || h != tt.headExpected {
				t.Errorf("test %s, list GetFirst(): err: got %s, expected nil\nhead: got %d, expected %d",
					tt.name, err.Error(), h, tt.headExpected)
			}
			// check tail
			if tl, err := list1.GetLast(); err != nil || tl != tt.tailExpected {
				t.Errorf("test %s, list GetLast(): err: got %s, expected nil\ntail: got %d, expected %d",
					tt.name, err.Error(), tl, tt.headExpected)
			}
		})
	}

	t.Cleanup(func() {
		list1 = nil
	})
}

func TestDoublyLinkedList_AddLast(t *testing.T) {
	// add last on nil list
	// add last when no elements
	// add last when more than element

	type testCase[T any] struct {
		name           string
		list           LinkedList[T]
		input          T
		lengthExpected int
		headExpected   T
		tailExpected   T
	}
	list1 := NewLinkedList[int](DoublyLinked)

	tests := []testCase[int]{
		{"add first item into list", list1, 1, 1, 1, 1},
		{"add second item into list", list1, 4, 2, 1, 4},
		{"add third item into list", list1, 5, 3, 1, 5},
		{"add more item into list", list1, 13, 4, 1, 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.list.AddLast(tt.input)
			// check len
			if l := list1.Len(); l != tt.lengthExpected {
				t.Errorf("test %s, list length: got %d, expected %d", tt.name, l, tt.lengthExpected)
			}
			// check head
			if h, err := list1.GetFirst(); err != nil || h != tt.headExpected {
				t.Errorf("test %s, list GetFirst(): err: got %s, expected nil\nhead: got %d, expected %d",
					tt.name, err.Error(), h, tt.headExpected)
			}
			// check tail
			if tl, err := list1.GetLast(); err != nil || tl != tt.tailExpected {
				t.Errorf("test %s, list GetLast(): err: got %s, expected nil\ntail: got %d, expected %d",
					tt.name, err.Error(), tl, tt.headExpected)
			}
		})
	}

	t.Cleanup(func() {
		list1 = nil
	})
}

func TestDoublyLinkedList_Get(t *testing.T) {
	list0 := NewLinkedList[int](DoublyLinked)
	oneEleList := NewLinkedListFrom(DoublyLinked, 4)
	list1 := NewLinkedListFrom(DoublyLinked, 13462, 76354192, 3645, 13262, 76354193, 43645, 4, 5, 8, 3, 14, 98, 3425, 56777, 24352)
	type testCase[T any] struct {
		name    string
		list    LinkedList[T]
		index   int
		want    T
		wantErr bool
		err     error
	}
	tests := []testCase[int]{
		{"get from invalid index", list1, -1, 0, true, ErrIndexOutOfBounds},
		{"get from invalid index 2", list0, list0.Len() + 1, 0, true, ErrIndexOutOfBounds},
		{"get 0th element from empty list", list0, 0, 0, true, ErrIndexOutOfBounds},
		{"get 0th element from normal list", list1, 0, 13462, false, nil},
		{"get 1th element from single element list", oneEleList, 1, 0, true, ErrIndexOutOfBounds},
		{"get last element from normal list", list1, list1.Len() - 1, 24352, false, nil},
		{"get element at a index", list1, 3, 13262, false, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.list.Get(tt.index)
			// validate error
			if err != nil {
				if !tt.wantErr {
					t.Errorf("Get() got unexpected error = %v", err)
				}
				if tt.wantErr && !errors.Is(err, tt.err) {
					t.Errorf("Get() got unexpected error = %v, expected %v", err, tt.err)
				}
			}
			// validate the value
			if got != tt.want {
				t.Errorf("Get() got = %v, expected %v", got, tt.want)
			}
		})
	}
}

func TestDoublyLinkedList_Insert(t *testing.T) {
	type testCase[T any] struct {
		name    string
		list    LinkedList[T]
		index   int
		element T
		result  bool
		wantErr bool
		err     error
		head    T
		tail    T
	}
	strList := NewLinkedList[string](DoublyLinked)
	tests := []testCase[string]{
		// insert into 0th index when list is empty, verify both head and tail are same element.
		{"insert to 0th index when list is empty", strList, 0, "2", true, false, nil, "2", "2"},
		// insert element at 0th index when list is not empty and verify head is new and tail is old element.
		{"insert to 0th index when list is not empty", strList, 0, "4", true, false, nil, "4", "2"},
		// insert into last index when list is not empty, verify the tail is the same element.
		{"insert to last index when list is not empty", strList, 2, "6", true, false, nil, "4", "6"},
		// insert into some middle index and verify expected element is at right index and head and tails are intact.
		{"insert at the middle", strList, 2, "8", true, false, nil, "4", "6"},
		// try to insert into invalid index, should get ErrIndexOutOfBounds
		{"insert to invalid index", strList, -1, "8", false, true, ErrIndexOutOfBounds, "4", "6"},
		{"insert to invalid index", strList, 5, "8", false, true, ErrIndexOutOfBounds, "4", "6"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ok, err := tt.list.Insert(tt.element, tt.index)
			if ok != tt.result {
				t.Errorf("Insert(%v, %v) got = %v expected %v", tt.element, tt.index, ok, tt.result)
			}
			if err != nil && !tt.wantErr {
				t.Errorf("Insert(%v, %v) unexpected error %v", tt.element, tt.index, err)
			}
			if tt.wantErr && !errors.Is(err, tt.err) {
				t.Errorf("Insert(%v, %v) GotErr = %v expectedErr %v", tt.element, tt.index, err, tt.err)
			}
			if !tt.wantErr && err == nil {
				// Get element from the inserted index and check returning element and inserted element are same
				if get, _ := tt.list.Get(tt.index); get != tt.element {
					t.Errorf("Insert(%v, %v) Got %v at index %d expected %v at this index", tt.element, tt.index, get, tt.index, tt.element)
				}

				// check head and tail
				if head, _ := tt.list.GetFirst(); head != tt.head {
					t.Errorf("Insert(%v, %v) Got %v at head expected %v", tt.element, tt.index, head, tt.head)
				}
				if tail, _ := tt.list.GetLast(); tail != tt.tail {
					t.Errorf("Insert(%v, %v) Got %v at the tail expected %v", tt.element, tt.index, tail, tt.tail)
				}
			}
		})
	}

	t.Cleanup(func() {
		strList = nil
	})
}

func TestDoublyLinkedList_RemoveFirst(t *testing.T) {
	list := NewLinkedList[int](DoublyLinked)
	// remove first from empty list - should return ErrNoSuchElement
	first, err := list.RemoveFirst()
	if err == nil || !errors.Is(err, ErrNoSuchElement) {
		t.Errorf("RemoveFirst() gotErr = %v, expectedErr %v", err, ErrNoSuchElement)
	}
	if first != 0 {
		t.Errorf("RemoveFirst() got = %v expected %v", first, 0)
	}

	// remove first from list with only element - verify list is empty
	valueTobeInserted := 43
	list.AddLast(valueTobeInserted)
	if v, err := list.RemoveFirst(); v != valueTobeInserted || err != nil {
		if valueTobeInserted != v {
			t.Errorf("RemoveFirst() got = %v expected %v", v, valueTobeInserted)
		}
		if err != nil {
			t.Errorf("RemoveFirst() unexpected error %v expected err %v", err, nil)
		}
	}
	if l := list.Len(); l != 0 {
		t.Errorf("RemoveFirst() got len %v expected len %v", l, 0)
	}
	// remove first from list with more elements - verify length and next head and tail
	for v := range 6 {
		list.AddLast(v + 1)
	}
	if v, err := list.RemoveFirst(); v != 1 || err != nil {
		if 1 != v {
			t.Errorf("RemoveFirst() got = %v expected %v", v, 1)
		}
		if err != nil {
			t.Errorf("RemoveFirst() unexpected error %v expected err %v", err, nil)
		}
	}
	if l := list.Len(); l != 5 {
		t.Errorf("RemoveFirst() got len %v expected len %v", l, 5)
	}
	if first, err := list.GetFirst(); first != 2 || err != nil {
		if 2 != first {
			t.Errorf("RemoveFirst() got = %v expected %v", first, 2)
		}
		if err != nil {
			t.Errorf("RemoveFirst() unexpected error %v expected err %v", err, nil)
		}
	}
}

func TestDoublyLinkedList_RemoveLast(t *testing.T) {
	// remove last from empty list - should return ErrNoSuchElement
	list := NewLinkedList[int](DoublyLinked)
	// remove first from empty list - should return ErrNoSuchElement
	last, err := list.RemoveLast()
	if err == nil || !errors.Is(err, ErrNoSuchElement) {
		t.Errorf("RemoveFirst() gotErr = %v, expectedErr %v", err, ErrNoSuchElement)
	}
	if last != 0 {
		t.Errorf("RemoveLast() got = %v expected %v", last, 0)
	}

	// remove last from list with only element - verify list is empty
	valueTobeInserted := 43
	list.AddLast(valueTobeInserted)
	if last, err := list.RemoveLast(); last != valueTobeInserted || err != nil {
		if valueTobeInserted != last {
			t.Errorf("RemoveLast() got = %v expected %v", last, valueTobeInserted)
		}
		if err != nil {
			t.Errorf("RemoveLast() unexpected error %v expected err %v", err, nil)
		}
	}
	if l := list.Len(); l != 0 {
		t.Errorf("RemoveLast() got len %v expected len %v", l, 0)
	}
	// remove last from list with more elements - verify length and next head and tail
	for v := range 6 {
		list.AddLast(v + 1)
	}
	if last, err := list.RemoveLast(); last != 6 || err != nil {
		if 6 != last {
			t.Errorf("RemoveLast() got = %v expected %v", last, 6)
		}
		if err != nil {
			t.Errorf("RemoveLast() unexpected error %v expected err %v", err, nil)
		}
	}
	if l := list.Len(); l != 5 {
		t.Errorf("RemoveLast() got len %v expected len %v", l, 5)
	}
	if first, err := list.GetFirst(); first != 2 || err != nil {
		if 1 != first {
			t.Errorf("RemoveLast() got = %v expected %v", first, 1)
		}
		if err != nil {
			t.Errorf("RemoveLast() unexpected error %v expected err %v", err, nil)
		}
	}
}

func TestDoublyLinkedList_RemoveAt(t *testing.T) {
	list1 := NewLinkedListFrom(DoublyLinked, 2, 3, 4, 5, 6, 7, 8, 1, 12, 24, 23, 54, 90)
	list2 := NewLinkedListFrom(DoublyLinked, 2, 4, 6, 8, 10, 12, 14)
	type testCase[T any] struct {
		name     string
		list     LinkedList[T]
		index    int
		want     T
		wantErr  bool
		err      error
		wantHead T
		wantTail T
		wantLen  int
	}
	tests := []testCase[int]{
		{"remove from index 0", list1, 0, 2, false, nil, 3, 90, list1.Len() - 1},                   // list will be 3, 4, 5, 6, 7, 8, 1, 12, 24, 23, 54, 90
		{"remove from last index", list1, list1.Len() - 2, 90, false, nil, 3, 54, list1.Len() - 2}, // list after this 3, 4, 5, 6, 7, 8, 1, 12, 24, 23, 54
		{"remove at index 3", list2, 3, 8, false, nil, 2, 14, list2.Len() - 1},
		{"remove from invalid index", list2, list2.Len(), 0, true, ErrIndexOutOfBounds, 2, 14, list2.Len() - 1},
		{"remove from empty list", NewLinkedList[int](DoublyLinked), 0, 0, true, ErrIndexOutOfBounds, 0, 0, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.list.RemoveAt(tt.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("RemoveAt() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveAt() got = %v, want %v", got, tt.want)
			}
			if f, _ := tt.list.GetFirst(); f != tt.wantHead {
				t.Errorf("RemoveAt() head got = %v, want %v", f, tt.wantHead)
			}
			if tl, _ := tt.list.GetLast(); tl != tt.wantTail {
				t.Errorf("RemoveAt() tail got = %v, want %v", tl, tt.wantTail)
			}
			if l := tt.list.Len(); l != tt.wantLen {
				t.Errorf("RemoveAt() length got = %v, want %v", l, tt.wantLen)
			}
		})
	}

	t.Cleanup(func() {
		list1, list2 = nil, nil
	})
}

func TestDoublyLinkedList_GetHeadAndTailNode(t *testing.T) {
	list1 := NewLinkedList[string](DoublyLinked)
	head, err := list1.GetHeadNode()
	// test with empty list
	if head != nil || !errors.Is(err, ErrNoSuchElement) {
		t.Errorf("got error %v expected error %v", err, ErrNoSuchElement)
	}
	tail, err := list1.GetTailNode()
	if tail != nil || !errors.Is(err, ErrNoSuchElement) {
		t.Errorf("got error %v expected error %v", err, ErrNoSuchElement)
	}

	// only one item in list
	list1.AddFirst("hello")
	head, err = list1.GetHeadNode()
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if head.Value() != "hello" || head.Next() != nil || head.Prev() != nil {
		t.Errorf("got head value %v next %v prev %v expected head value %v next %v prev %v",
			head.Value(), head.Next(), head.Prev(), "hello", nil, nil)
	}
	tail, err = list1.GetTailNode()
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if tail.Value() != "hello" || tail.Next() != nil || tail.Prev() != nil {
		t.Errorf("got tail value %v next %v prev %v expected tail value %v next %v prev %v",
			tail.Value(), tail.Next(), tail.Prev(), "hello", nil, nil)
	}
	if head != tail {
		t.Errorf("expected %v == %v", head, tail)
	}

	// 2 items in the list
	list1.AddLast("world")
	head, err = list1.GetHeadNode()
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if head.Value() != "hello" || head.Next().Value() != "world" || head.Prev() != nil {
		t.Errorf("got head value %v next %v prev %v expected head value %v next %v prev %v",
			head.Value(), head.Next().Value(), head.Prev(), "hello", "world", nil)
	}
	tail, err = list1.GetTailNode()
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if tail.Value() != "world" || tail.Next() != nil || tail.Prev().Value() != "hello" {
		t.Errorf("got tail value %v next %v prev %v expected tail value %v next %v prev %v",
			tail.Value(), tail.Next(), tail.Prev().Value(), "world", nil, "hello")
	}

	// 3 items in the list
	list1.AddLast("last")
	head, err = list1.GetHeadNode()
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if head.Value() != "hello" || head.Next().Value() != "world" || head.Prev() != nil {
		t.Errorf("got head value %v next %v prev %v expected head value %v next %v prev %v",
			head.Value(), head.Next().Value(), head.Prev(), "hello", "world", nil)
	}
	tail, err = list1.GetTailNode()
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if tail.Value() != "last" || tail.Next() != nil || tail.Prev().Value() != "world" {
		t.Errorf("got tail value %v next %v prev %v expected tail value %v next %v prev %v",
			tail.Value(), tail.Next(), tail.Prev().Value(), "last", nil, "world")
	}
}

func TestDoublyLinkedList_ToSlice(t *testing.T) {
	// nil list
	var nilList *DoublyLinkedList[int]
	if s := nilList.ToSlice(); s != nil {
		t.Errorf("ToSlice() got = %v expected %v", s, nil)
	}
	// empty list
	sList := NewLinkedList[int](DoublyLinked)
	if s := sList.ToSlice(); len(s) != 0 {
		t.Errorf("ToSlice() got = %v expected %v", s, make([]int, 0))
	}

	// single element
	if s := sList.AddLast(2).ToSlice(); len(s) != 1 || s[0] != 2 {
		t.Errorf("ToSlice() got = %v expected %v", s, []int{2})
	}
	// 2 elements
	if s := sList.AddLast(4).ToSlice(); len(s) != 2 || s[0] != 2 || s[1] != 4 {
		t.Errorf("ToSlice() got = %v expected %v", s, []int{2, 4})
	}
	// more elements
	if s := sList.AddLast(6).AddLast(8).AddLast(10).ToSlice(); len(s) != 5 || s[0] != 2 || s[len(s)-1] != 10 || s[1] != 4 {
		t.Errorf("ToSlice() got = %v expected %v", s, []int{2, 4, 6, 8, 10})
	}
}

func TestDoublyLinkedList_Values(t *testing.T) {
	nums := make([]int, 10000)
	for i := 0; i < cap(nums); i++ {
		nums[i] = rand.IntN(cap(nums))
	}

	dList := NewLinkedListFromSlice(DoublyLinked, nums)
	i := 0
	for v := range dList.Values() {
		if v != nums[i] {
			t.Errorf("Values() got = %v, expected %v", v, nums[i])
		}
		i++
	}
}

func TestDoublyLinkedList_ReverseAll(t *testing.T) {
	nums := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		nums[i] = rand.IntN(1000)
	}

	sList := NewLinkedListFromSlice(DoublyLinked, nums)
	j := len(nums) - 1
	for i, v := range sList.ReverseAll() {
		if i != j || v != nums[j] {
			t.Errorf("ReverseAll() got index %d value %d expected index %d value %d", i, v, j, nums[j])
		}
		j--
	}
}

func TestDoublyLinkedList_String(t *testing.T) {
	var nilList *SinglyLinkedList[int]
	list0 := NewLinkedList[int](DoublyLinked)
	list1 := NewLinkedListFrom[int](DoublyLinked, 12)
	list2 := NewLinkedListFrom[int](DoublyLinked, 1, 2, 3, 4)
	type testCase[T any] struct {
		name string
		s    LinkedList[T]
		want string
	}
	tests := []testCase[int]{
		{"nil list", nilList, "<nil>"},
		{"empty list", list0, ""},
		{"list with one element", list1, "12"},
		{"list with more elements", list2, "1 <=> 2 <=> 3 <=> 4"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.String(); got != tt.want {
				t.Errorf("String() = got = %v, expected %v", got, tt.want)
			}
		})
	}
}
