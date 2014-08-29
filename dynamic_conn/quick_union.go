package dynamiccon

type QuickUnion UnionFind

func (ufp *QuickUnion) initUF (size int) {
	ufp.objs = make([]int, size)
	genericInitUF(ufp.objs, size)
}

func (ufp *QuickUnion) unionPair (pair Pair) {
	ufp.union(pair.first, pair.second)
}

func (ufp *QuickUnion) union (first, second int) {
	ufp.objs[first] = ufp.objs[second]
}

func (ufp *QuickUnion) connectedPair (pair Pair) bool {
	return ufp.connected(pair.first, pair.second)
}

func (ufp *QuickUnion) connected (first, second int) bool {
	firstRoot := ufp.objs[first]
	for firstRoot != ufp.objs[firstRoot] {
		firstRoot = ufp.objs[firstRoot]
	}
	secondRoot := ufp.objs[second]
	for secondRoot != ufp.objs[secondRoot] {
		secondRoot = ufp.objs[secondRoot]
	}
	return firstRoot == secondRoot
}

func (ufp *QuickUnion) print () {
	genericPrint(ufp.objs)
}
