package sorts

type Sortable interface {
	Len () uint
	Less (first, second uint) bool
	Swap (first, second uint)
}
