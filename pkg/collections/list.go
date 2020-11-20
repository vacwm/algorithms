package collections

// Node contains data (int) and a pointer to its successor node.
type ListNode struct {
	Val  interface{}
	Next *ListNode
}

func ArrToList(arr []interface{}) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	if len(arr) == 1 {
		return &ListNode{arr[0], nil}
	}
	return &ListNode{arr[0], ArrToList(arr[1:])}
}

// Reverse changes direction of nodes in the list and returns the new head.
func (l *ListNode) Reverse() *ListNode {
	if l == nil {
		return nil
	}

	var prev *ListNode
	cur := l

	for {
		if cur.Next == nil {
			cur.Next = prev
			return cur
		}
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
}

func (l1 *ListNode) Equals(l2 *ListNode) bool {
	if l1 == nil && l2 == nil {
		return true
	}

	if l1 == nil {
		return false
	}

	if l2 == nil {
		return false
	}

	if l1.Val != l2.Val {
		return false
	}

	return l1.Next.Equals(l2.Next)
}

// MergeLinkedLists
// Precondition:
//   l1 is in non-decreasing order
//   l2 is in non-decreasing order
//   l1 and l2 Value type is int
func MergeLinkedLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Val.(int) < l2.Val.(int) {
		return &ListNode{l1.Val, MergeLinkedLists(l1.Next, l2)}
	} else {
		return &ListNode{l2. Val, MergeLinkedLists(l1, l2.Next)}
	}
}

