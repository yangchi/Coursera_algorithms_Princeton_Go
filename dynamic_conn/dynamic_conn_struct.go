package dynamiccon

type Pair struct{
	first int
	second int
}

type UnionFind struct {
	objs []int
}

type UnionFinder interface {
	initUF(int)
	unionPair(Pair)
	union(int, int)
	connectedPair(Pair) (bool)
	connected(int, int) (bool)
	print()
}

