package dynamiccon

import ("testing"; "os"; "strings"; "bufio"; "strconv")

type unionFunc func (first, second int) ()

func fileReader () (total int, pairs []Pair)  {
	f, _ := os.Open("test.input")
	bf := bufio.NewReader(f)
	defer f.Close()
	str, ok := bf.ReadString('\n')
	str = strings.Trim(str, "\n")
	if ok != nil { return }
	str_slice := strings.Split(str, " ")
	if len(str_slice) != 1 {
		return
	} else {
		atmost, ok := strconv.Atoi(str_slice[0])
		if ok != nil {
			return
		}
		total = atmost
		pairs = make([]Pair, 0, total)
	}
	counter := 0
	for counter < total {
		str, ok = bf.ReadString('\n')
		str = strings.Trim(str, "\n")
		if ok != nil { 
			break 
		}
		str_slice := strings.Split(str, " ")
		first, _ := strconv.Atoi(str_slice[0])
		second, _ := strconv.Atoi(str_slice[1])
		pair := Pair{first, second}
		pairs = append(pairs, pair)
		counter += 1
	}
	return
}

func TestFileClientQuickFind(t *testing.T) {
	var finder UnionFinder
	var quickFinder QuickFind
	finder = &quickFinder
	finder.initUF(100)
	finder = &quickFinder
	_, pairs := fileReader()
	for _, pair := range pairs {
		finder.unionPair(pair)
		if !finder.connectedPair(pair) {
			t.Errorf("%d and %d should be connected", pair.first, pair.second)
		}
	}
}

func moreGenericUnionFindTest (finder QuickUnion, unioner unionFunc) (bool, string){
	unioner(4, 3)
	unioner(3, 8)
	unioner(6, 5)
	unioner(9, 4)
	unioner(2, 1)
	if finder.connected(0, 7) {
		return false, "0 and 7 should not be connected"
	}
	if !finder.connected(8, 9) {
		return false, "8 and 9 should be connected"
	}
	unioner(0, 5)
	unioner(7, 2)
	unioner(6, 1)
	unioner(1, 0)
	if !finder.connected(0, 7) {
		return false, "0 and 7 should be connected"
	}
	return true, ""
}

func genericUnionFindTest (finder UnionFinder) (bool, string){
	finder.union(4, 3)
	finder.union(3, 8)
	finder.union(6, 5)
	finder.union(9, 4)
	finder.union(2, 1)
	if finder.connected(0, 7) {
		return false, "0 and 7 should not be connected"
	}
	if !finder.connected(8, 9) {
		return false, "8 and 9 should be connected"
	}
	finder.union(0, 5)
	finder.union(7, 2)
	finder.union(6, 1)
	finder.union(1, 0)
	if !finder.connected(0, 7) {
		return false, "0 and 7 should be connected"
	}
	return true, ""
}

func TestQuickUnion(t *testing.T) {
	var finder UnionFinder
	var quickUnion QuickUnion
	finder = &quickUnion
	finder.initUF(10)
	res, msg := genericUnionFindTest(finder)
	if !res {
		t.Error(msg)
	}
}

func TestQuickFind(t *testing.T) {
	var finder UnionFinder
	var quickFinder QuickFind
	finder = &quickFinder
	finder.initUF(10)
	res, msg := genericUnionFindTest(finder)
	if !res {
		t.Error(msg)
	}
}

func BenchmarkQuickUnion(b *testing.B) {
	var quickUnion QuickUnion
	for i := 0; i < b.N; i++ {
		quickUnion.initUF(10)
		moreGenericUnionFindTest(quickUnion, quickUnion.union)
	}
}

func BenchmarkQuickUnionWeighted(b *testing.B) {
	var quickUnion QuickUnion
	for i := 0; i < b.N; i++ {
		quickUnion.initUF(10)
		moreGenericUnionFindTest(quickUnion, quickUnion.weightedUnion)
	}
}

func BenchmarkQuickFind(b *testing.B) {
	var finder UnionFinder
	var quickFinder QuickFind
	finder = &quickFinder
	for i := 0; i < b.N; i++ {
		finder.initUF(10)
		genericUnionFindTest(finder)
	}
}
