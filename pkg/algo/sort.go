package algo

import (
	"errors"
	"fmt"
)

// BinarySearch(A, k)
// Input: A presorted array A, and element k
// Output: If A contains the element k, return the index i such A[i] = k, else -1
func BinarySearch(arr []int, k int) (int, error) {
	if len(arr) == 0 {
		return -1, errors.New("Element could not be found, reason: empty array")
	}
	i := 0
	j := len(arr) - 1
	mid := i + (j-i)/2

	for {
		if i > j {
			break
		}
		if arr[mid] == k {
			return mid, nil
		} else if arr[mid] < k {
			i = mid + 1
			mid = i + (j-i)/2
		} else {
			j = mid - 1
			mid = i + (j-i)/2
		}
	}
	return -1, errors.New(fmt.Sprintf("Element, %d, could not be found.", k))
}
