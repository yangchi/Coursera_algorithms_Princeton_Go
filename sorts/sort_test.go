package sorts

import (
	"testing"
	"fmt"
)

func TestSelectionSortIntSlice (t *testing.T) {
	myslice := []int{5,3,2,6,7,8,2,6,7}
	selectionSortIntSlice(myslice)
	result := verifySortedIntSlice(myslice)
	if !result {
		t.Error("Failed")
	}
}

func TestMergeIntSlice (t *testing.T) {
	myslice := []int{1,5,7,13,2,4,11,12,19,27}
	temp := make([]int, len(myslice))
	mergeIntSlice(myslice, temp, 0, 3, uint(len(myslice)-1))
	result := verifySortedIntSlice(myslice)
	if !result {
		for _, elem := range myslice {
			fmt.Println(elem)
		}
		t.Error("Merge failed")
	}
}

func TestMergeSortIntSlice (t *testing.T) {
	myslice := []int{1,5,7,13,2,4,11,12,19,27}
	mergeSortIntSlice(myslice)
	result := verifySortedIntSlice(myslice)
	if !result {
		t.Error("Merge sort failed")
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

func genericVerify (data Sortable) (bool) {
	var index uint
	for index = 0; index < data.Len(); index++ {
		if index == 0 {
			continue
		}
		if data.Less(index, index - 1) {
			return false
		}
	}
	return true
}

func TestInsertionSort (t *testing.T) {
	var sort Sortable
	data := []int{5,3,2,6,7,8,2,6,7}
	intSlice := IntSlice{data}
	sort = &intSlice
	InsertionSort(sort)
	result := genericVerify(sort)
	if !result {
		t.Error("Insertion sort failed")
	}
}

func TestGenericSelectionSort (t *testing.T) {
	var sort Sortable
	intSlice := IntSlice{[]int{5,3,2,6,7,8,2,6,7}}
	sort = &intSlice
	SelectionSort(sort)
	result := genericVerify(sort)
	if !result {
		t.Error("Selection sort failed")
	}
}
