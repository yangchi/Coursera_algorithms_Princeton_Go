package dynamiccon

import "testing"

func laterTestOps(t *testing.T) {
	union(4, 3)
	union(3, 8)
	union(6, 5)
	union(9, 4)
	union(2, 1)
	if connected(0, 7) {
		t.Error("0 and 7 should not be connected")
	}
	if !connected(8, 9) {
		t.Error("8 and 9 should be connected")
	}
	union(0, 5)
	union(7, 2)
	union(6, 1)
	union(1, 0)
	if !connected(0, 7) {
		t.Error("0 and 7 should be connected")
	}
}

func TestUFInit(t *testing.T) {
	uf := UnionFind{make([]int, 100)}
	var uf2 UnionFind
	uf2.objs = make([]int, 200)
	var uf3 UnionFind
	uf3.initUF(300)
	if len(uf.objs) != 100 {
		t.Error("uf definition failed")
	}
	if len(uf2.objs) != 200 {
		t.Error("uf2 slice making failed")
	}
	if len(uf3.objs) != 300 {
		t.Error("uf3 init failed")
	}
}
