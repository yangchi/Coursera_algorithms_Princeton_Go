package dynamiccon

import "fmt"

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

func genericInitUF (objects []int, size int) {
	for index, _ := range objects{
		objects[index] = index
	}
}

func genericPrint (objects []int) {
	for _, obj := range objects {
		fmt.Printf("%d  ", obj)
	}
	fmt.Println()
}

