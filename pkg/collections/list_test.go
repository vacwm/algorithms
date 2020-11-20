package collections

import (
	"testing"
)

func sortedList() *ListNode {
	c := &ListNode{3, nil}
	b := &ListNode{2, c}
	a := &ListNode{1, b}
	return a
}

func alternatingSortedList() *ListNode {
	z := &ListNode{3, nil}
	y := &ListNode{4, z}
	x := &ListNode{2, y}
	w := &ListNode{5, x}
	v := &ListNode{1, w}
	return v
}

func singleElementList() *ListNode {
	return &ListNode{1, nil}
}

func palindromeList() *ListNode {
	d := &ListNode{1, nil}
	c := &ListNode{2, d}
	b := &ListNode{2, c}
	a := &ListNode{1, b}
	return a
}

func TestListNode(t *testing.T) {
	tests := []struct {
		arr []int
		list *ListNode
	}{
		{[]int{1,2,3}, sortedList()},
		{[]int{1,5,2,4,3}, alternatingSortedList()},
		{[]int{1}, singleElementList()},
		{[]int{1,2,2,1}, palindromeList()},
	}

	for testId, test := range tests {
		nums := make([]interface{}, len(test.arr))
		for i, num := range test.arr {
			nums[i] = num
		}
		l := ArrToList(nums)
		if !l.Equals(test.list) {
			t.Errorf("TestListNode #%d failed.", testId)
		}
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		arr []int
		list *ListNode
	}{
		{[]int{3,2,1}, sortedList()},
		{[]int{3,4,2,5,1}, alternatingSortedList()},
		{[]int{1}, singleElementList()},
		{[]int{1,2,2,1}, palindromeList()},
	}

	for testId, test := range tests {
		nums := make([]interface{}, len(test.arr))
		for i, num := range test.arr {
			nums[i] = num
		}
		l := ArrToList(nums)
		if !l.Equals(test.list.Reverse()) {
			t.Errorf("TestReverse #%d failed.", testId)
		}
	}
}

func TestMergeLinkedLists(t *testing.T) {
	tests := []struct {
		arr1 []int
		arr2 []int
		expect []int
	}{
		{[]int{1,2,3}, []int{4,5,6}, []int{1,2,3,4,5,6}},
		{[]int{3,7,9}, []int{4,5,8}, []int{3,4,5,7,8,9}},
		{[]int{1}, []int{}, []int{1}},
		{[]int{}, []int{}, []int{}},
	}

	for testId, test := range tests {
		list1 := make([]interface{}, len(test.arr1))
		list2 := make([]interface{}, len(test.arr2))
		expectList := make([]interface{}, len(test.expect))
		for i, num := range test.arr1 {
			list1[i] = num
		}
		for i, num := range test.arr2 {
			list2[i] = num
		}
		for i, num := range test.expect {
			expectList[i] = num
		}

		l1 := ArrToList(list1)
		l2 := ArrToList(list2)
		e  := ArrToList(expectList)
		if !MergeLinkedLists(l1, l2).Equals(e) {
			t.Errorf("TestMergeLinkedLists #%d failed.", testId)
		}
	}

}

