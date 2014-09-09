package sorts

/*
	Merge data[lo...mid] and data[mid+x...high] each of which
	are sorted
*/
func mergeIntSlice (data, temp []int, lo, mid, high uint) {
	//temp := make([]int, len(data))
	second_half_index := mid + 1
	index := 0
	for index < len(temp) {
		if data[lo] < data[second_half_index] {
			temp[index] = data[lo]
			lo++
			index++
			if lo > mid {
				break
			}
		} else {
			temp[index] = data[second_half_index]
			second_half_index++
			index++
			if second_half_index > high {
				break
			}
		}
	}
	temp_slice := temp[index:]
	var temp_data_slice []int
	if lo > mid {
		temp_data_slice = data[second_half_index:]
	} else if second_half_index > high {
		temp_data_slice = data[lo:mid]
	}
	if len(temp_data_slice) > 0 {
		copy(temp_slice, temp_data_slice)
	}
	copy(data, temp)
}

func mergeSortHelperIntSlice (data, temp []int, lo, hi uint) {
	if lo >= hi {
		return
	}
	mid := lo + (hi - lo) / 2
	mergeSortHelperIntSlice(data, temp, lo, mid)
	mergeSortHelperIntSlice(data, temp, mid + 1, hi)
	mergeIntSlice(data, temp, lo, mid, hi)
}

func mergeSortIntSlice (data []int) {
	temp := make([]int, len(data))
	mergeSortHelperIntSlice(data, temp, 0, uint(len(data) - 1))
}
