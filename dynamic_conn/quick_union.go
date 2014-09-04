package dynamiccon

type QuickUnion struct {
	objs []int
	sizes []int
}

func (ufp *QuickUnion) initUF (size int) {
	genericInitUF(&ufp.objs, size)
	ufp.sizes = make([]int, size)
	for index, _ := range ufp.sizes {
		ufp.sizes[index] = 1
	}
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

func (ufp *QuickUnion) weightedUnion (first, second int) {
	first_root := ufp.root(first)
	second_root := ufp.root(second)
	if first_root == second_root {
		return
	}
	if ufp.sizes[first_root] < ufp.sizes[second_root] {
		ufp.objs[first_root] = second_root
		ufp.sizes[second_root] += ufp.sizes[first_root]
	} else {
		ufp.objs[second_root] = first_root
		ufp.sizes[first_root] += ufp.sizes[second_root]
	}
}
