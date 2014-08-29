package dynamiccon


type QuickFind UnionFind


func (ufp *QuickFind) initUF (size int) {
	genericInitUF(&ufp.objs, size)
}

func (ufp *QuickFind) unionPair (pair Pair) {
	ufp.union(pair.first, pair.second)
}

func (ufp *QuickFind) union (first, second int) {
	pid := ufp.objs[first]
	qid := ufp.objs[second]
	for index, obj := range ufp.objs {
		if obj == pid {
			ufp.objs[index] = qid
		}
	}
}

func (ufp *QuickFind) connected (first, second int) (bool) {
	return ufp.objs[first] == ufp.objs[second]
}

func (ufp *QuickFind) connectedPair (pair Pair) (bool) {
	return ufp.objs[pair.first] == ufp.objs[pair.second]
}

func (ufp *QuickFind) print () () {
	genericPrint(ufp.objs)
}
