package dynamiccon

import "testing"

func Test(t *testing.T) {
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
