package dynamiccon

type UnionFind struct {
	objs []int
}

func (ufp *UnionFind) initUF (size int) {
	ufp.objs = make([]int, size)
}

/*
type UnionFinder interface {
	initUF(int)
	union(int, int)
	connected(int, int)
}
*/

