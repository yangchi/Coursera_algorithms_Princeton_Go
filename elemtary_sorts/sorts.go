package sorts

type Sortable interface {
	Less (first, second uint) bool
	Swap (first, second uint)
}
