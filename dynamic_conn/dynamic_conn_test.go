package dynamiccon

import ("testing"; "os"; "strings"; "bufio"; "strconv";)

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

func TestQuickFind(t *testing.T) {
	var finder UnionFinder
	var quickFinder QuickFind
	finder = &quickFinder
	finder.initUF(10)
	finder.print()
	finder.union(4, 3)
	finder.print()
	finder.union(3, 8)
	finder.print()
	finder.union(6, 5)
	finder.union(9, 4)
	finder.print()
	finder.union(2, 1)
	finder.print()
	if finder.connected(0, 7) {
		t.Error("0 and 7 should not be connected")
	}
	if !finder.connected(8, 9) {
		t.Error("8 and 9 should be connected")
	}
	finder.union(0, 5)
	finder.union(7, 2)
	finder.union(6, 1)
	finder.union(1, 0)
	if !finder.connected(0, 7) {
		t.Error("0 and 7 should be connected")
	}
}
