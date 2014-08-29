package dynamiccon

type QuickUnion UnionFind

func (ufp *QuickUnion) initUF (size int) {
	ufp.objs = make([]int, size)
	genericInitUF(ufp.objs, size)
}

func (ufp *QuickUnion) root (index int) (root int) {
	root = ufp.objs[index]
	for root != ufp.objs[root] {
		root = ufp.objs[root]
	}
	return
}

func (ufp *QuickUnion) unionPair (pair Pair) {
	ufp.union(pair.first, pair.second)
}

func (ufp *QuickUnion) union (first, second int) {
	ufp.objs[first] = ufp.root(second)
}

func (ufp *QuickUnion) connectedPair (pair Pair) bool {
	return ufp.connected(pair.first, pair.second)
}

func (ufp *QuickUnion) connected (first, second int) bool {
	firstRoot := ufp.root(first)
	secondRoot := ufp.root(second)
	return firstRoot == secondRoot
}

func (ufp *QuickUnion) print () {
	genericPrint(ufp.objs)
}
