package sorts

type IntSlice struct {
	data []int
}

func selectionSortIntSlice (data []int) {
	if len(data) < 2 {
		return
	}
	for target_index, elem := range data {
		mini := elem
		mini_index := target_index
		for search_index := target_index; search_index < len(data); search_index++ {
			if data[search_index] < mini {
				mini = data[search_index]
				mini_index = search_index
			}
		}
		swapIntSlice(data, target_index, mini_index)
	}
}

func swapIntSlice (data []int, first, second int) {
	temp := data[first]
	data[first] = data[second]
	data[second] = temp
}

func InsertionSort (data Sortable) {
	if data.Len() < 2 {
		return
	}
	var target_index uint
	for target_index = 0; target_index < data.Len(); target_index++ {
		mini_index := target_index
		for search_index := target_index; search_index < data.Len(); search_index++ {
			if data.Less(search_index, mini_index) {
				mini_index = search_index
			}
		}
		data.Swap(target_index, mini_index)
	}
}

func (intSlice *IntSlice) Len () uint {
	return uint(len(intSlice.data))
}

func (intSlice *IntSlice) Less (first, second uint) bool {
	if first > uint(len(intSlice.data)) - 1 || 
	   second > uint(len(intSlice.data)) -1 {
		return false
	}
	return intSlice.data[first] < intSlice.data[second]
}

func (intSlice *IntSlice) Swap (first, second uint) {
	temp := intSlice.data[first]
	intSlice.data[first] = intSlice.data[second]
	intSlice.data[second] = temp
}


