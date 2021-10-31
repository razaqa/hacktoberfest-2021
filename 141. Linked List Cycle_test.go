package linkedlist

import "testing"

func TestHasCycle(t *testing.T) {
	var testCases = []struct {
		nums *ListNode
		want bool
	}{
		{listNode1(), true},
		{listNode2(), true},
		{listNode3(), false},
	}

	for _, tc := range testCases {
		got := HasCycle(tc.nums)
		if got != tc.want {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}

}

func listNode1() *ListNode {
	node1 := &ListNode{Val: 3}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 0}
	node4 := &ListNode{Val: 4}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2

	return node1
}

func listNode2() *ListNode {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}

	node1.Next = node2
	node2.Next = node1

	return node1
}

func listNode3() *ListNode {
	node1 := &ListNode{Val: 1}
	return node1
}
