package sorts

import (
	"testing"
)

func TestSelectionSortIntSlice (t *testing.T) {
	myslice := []int{5,3,2,6,7,8,2,6,7}
	selectionSortIntSlice(myslice)
	result := verifySortedIntSlice(myslice)
	if !result {
		t.Error("Failed")
	}
}

func verifySortedIntSlice (data []int) (bool) {
	for index, elem := range data {
		if index == 0 {
			continue
		}
		if elem < data[index-1] {
			return false
		}
	}
	return true
}
