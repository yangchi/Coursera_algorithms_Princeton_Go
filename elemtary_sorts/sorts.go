package sorts

type Sortable interface {
	Len () uint
	Less (first, second uint) bool
	Swap (first, second uint)
}


func InsertionSort (data Sortable) {
	if data.Len() < 2 {
		return
	}
	var index uint
	for index = 1; index < data.Len(); index++ {
		var lookback int
		for lookback = int(index) - 1; lookback >= 0; lookback-- {
			if data.Less(uint(lookback + 1), uint(lookback)) {
				data.Swap(uint(lookback + 1), uint(lookback))
			}
		}
	}
}
