package list

import (
	"testing"
)

func TestImmutableSinglyLinkedNode(t *testing.T) {
	sn := &singlyLinkedNode[int]{3, &singlyLinkedNode[int]{4, nil}}
	snStr := "*[3|—]⃓——→ [4|—]⃓"

	if sn.Value() != 3 {
		t.Errorf("got value %d expected = %d", sn.Value(), 3)
	}
	if nxt := sn.Next(); nxt.Value() != 4 && nxt.Next() != nil && nxt.Prev() != nil {
		t.Errorf("got value %d expected = %d", nxt.Value(), 4)
		t.Errorf("got next's next %v expected = %v", nxt.Next(), nil)
		t.Errorf("got next's prev %v expected = %v", nxt.Prev(), nil)
	}
	if sn.String() != snStr {
		t.Errorf("got string %s expected = %s", sn.String(), snStr)
	}

	sn = nil
	if s := sn.String(); s != "<nil>" {
		t.Errorf("got string %s expected = %s", sn.String(), snStr)
	}
}

func TestImmutableDoublyLinkedNode(t *testing.T) {
	// nil <———— node1 <————> node2 <————> node3 ————> nil
	node1 := &doublyLinkedNode[int]{value: 2}
	node2 := &doublyLinkedNode[int]{value: 4}
	node3 := &doublyLinkedNode[int]{value: 6}

	node1.prev, node1.next = nil, node2
	node2.prev, node2.next = node1, node3
	node3.prev, node3.next = node2, nil

	// test for Value() method
	if node1.Value() != 2 {
		t.Errorf("got value %v, expected value = %v", node1.Value(), 2)
	}
	if node2.Value() != 4 {
		t.Errorf("got value %v, expected value = %v", node2.Value(), 4)
	}
	if node3.Value() != 6 {
		t.Errorf("got value %v, expected value = %v", node3.Value(), 6)
	}

	// test for Next() method
	if nxt := node1.Next(); nxt != node2 {
		t.Errorf("got next %v, expected next = %v", nxt, node2)
	}
	if nxt := node2.Next(); nxt != node3 {
		t.Errorf("got next %v, expected next = %v", nxt, node3)
	}
	if nxt := node3.Next(); nxt != nil {
		t.Errorf("got next %v, expected next = %v", nxt, node3)
	}

	// test for Prev() method
	if prv := node1.Prev(); prv != nil {
		t.Errorf("got Prev %v, expected Prev = %v", prv, nil)
	}
	if prv := node2.Prev(); prv != node1 {
		t.Errorf("got Prev %v, expected Prev = %v", prv, node1)
	}
	if prv := node3.Prev(); prv != node2 {
		t.Errorf("got Prev %v, expected Prev = %v", prv, node2)
	}

	// String method
	if str := node1.String(); str != "<nil><——|2|——>|4| " {
		t.Errorf("got String %s, expected String = %s", str, "<nil><——|2|——>|4| ")
	}
	if str := node2.String(); str != "|2|<——|4|——>|6| " {
		t.Errorf("got String %s, expected String = %s", str, "|2|<——|4|——>|6| ")
	}
	if str := node3.String(); str != "|4|<——|6|——><nil> " {
		t.Errorf("got String %s, expected String = %s", str, "|4|<——|6|——><nil> ")
	}
	t.Cleanup(func() {
		node1, node2, node3 = nil, nil, nil
	})
}

func TestSinglyLinkedNodes(t *testing.T) {
	var nilNodes *SinglyLinkedNode[int]
	if nilNodes.Len() != 0 {
		t.Errorf("got length %v expected 0", nilNodes.Len())
	}
	if nilNodes.String() != "<nil>" {
		t.Errorf("got String() %v expected <nil>", nilNodes.String())
	}

	emptyNodes := AsSinglyLinkedNodes[int]()
	if emptyNodes.Value != 0 || emptyNodes.Len() != 1 || emptyNodes.Next != nil {
		t.Errorf("empty node expected value %v len %v next %v got  value %v len %v next %v",
			0, 0, nil, emptyNodes.Value, emptyNodes.Len(), emptyNodes.Next)
	}
	if emptyNodes.String() != "[0|—]——→<nil>" {
		t.Errorf("got String %v expected [0|—]——→<nil>", emptyNodes.String())
	}

	nums := []int{2, 4, 6, 8, 10}
	head := AsSinglyLinkedNodes[int](nums...)
	if head == nil {
		t.Errorf("expected non nil head but got nil head")
	}
	if head.Len() != 5 {
		t.Errorf("got length %v expected lenght %v", head.Len(), 5)
	}
	if head.String() != "[2|—]——→[4|—]——→[6|—]——→[8|—]——→[10|—]——→<nil>" {
		t.Errorf("got String %v expected [2|—]——→[4|—]——→[6|—]——→[8|—]——→[10|—]——→<nil>", emptyNodes.String())
	}

	for cur, i := head, 0; cur.Next != nil; cur, i = cur.Next, i+1 {
		t.Logf("got %v at inded %d", cur.Value, i)
		if cur.Value != nums[i] {
			t.Errorf("got %v at index %d but expected %v", cur.Value, i, nums[i])
		}
	}
}

func TestDoublyLinkedListNodes(t *testing.T) {
	var nilNodes *DoublyLinkedNode[int]
	if nilNodes.Len() != 0 {
		t.Errorf("got length %v expected 0", nilNodes.Len())
	}
	if nilNodes.String() != "<nil>" {
		t.Errorf("got String() %v expected <nil>", nilNodes.String())
	}

	emptyNodes := AsDoubleLinkedNodes[int]()
	if emptyNodes.Value != 0 || emptyNodes.Len() != 1 || emptyNodes.Next != nil || emptyNodes.Prev != nil {
		t.Errorf("empty node expected value=%v len=%v next=%v prev=%v but got value=%v len=%v next=%v prev=%v",
			0, 1, nil, nil, emptyNodes.Value, emptyNodes.Len(), emptyNodes.Next, emptyNodes.Prev)
	}
	nodeStr := "<nil>←——[—|0|—]——→<nil>"
	if emptyNodes.String() != nodeStr {
		t.Errorf("got String %s but expected %s", emptyNodes.String(), nodeStr)
	}

	nums := []int{2, 4, 6, 8, 10}
	head := AsDoubleLinkedNodes[int](nums...)
	if head == nil {
		t.Errorf("expected non nil head but got nil head")
	}
	if head.Len() != 5 || head.Value != 2 || head.Next.Value != 4 || head.Prev != nil || head.Next.Next.Value != 6 {
		t.Errorf("got value=%v len=%d next=%v prev=%v next.next=%v but expected value=%v len=%d next=%v prev=%v next.next=%v",
			head.Value, head.Len(), head.Next.Value, head.Prev.Value, head.Next.Next.Value, 2, 5, 4, nil, 6)
	}
	nodeStr = "<nil>←——[—|2|—]←——→[—|4|—]←——→[—|6|—]←——→[—|8|—]←——→[—|10|—]——→<nil>"
	if head.String() != nodeStr {
		t.Errorf("got String %v expected %s", head.String(), nodeStr)
	}

	for cur, i := head, 0; cur.Next != nil; cur, i = cur.Next, i+1 {
		t.Logf("got %v at inded %d", cur.Value, i)
		if cur.Value != nums[i] {
			t.Errorf("got %v at index %d but expected %v", cur.Value, i, nums[i])
		}
	}

}
