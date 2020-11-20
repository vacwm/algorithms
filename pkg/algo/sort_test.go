package algo

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		arr    []int
		k      int
		expect int
	}{
		{[]int{1, 2}, 2, 1},
		{[]int{1, 3, 5}, 1, 0},
		{[]int{1, 1, 1, 6, 8}, 6, 3},
		{[]int{}, 1, -1},
		{[]int{1, 3, 5, 7, 9}, 10, -1},
		{[]int{1, 3, 5, 7, 9}, 0, -1},
	}

	for i, test := range tests {
		defer func() {
			if recover() != nil {
				t.Errorf("TestBinarySearch #%d panic'd", i)
			}
		}()
		result, _ := BinarySearch(test.arr, test.k)
		if result != test.expect {
			t.Errorf("TestBinarySearch #%d failed: expected: %d, got: %d", i, test.expect, result)
		}
	}
}
