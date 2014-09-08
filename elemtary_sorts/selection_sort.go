package sorts

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
