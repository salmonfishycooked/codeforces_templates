package seg_tree_lazy

import (
	"testing"
)

var a = []int{5, 7, 2, 12, 9, 5, 13, 15}

func TestSegTreeLazy_Query(t *testing.T) {
	seg := NewSegTreeLazy(a)
	if v := seg.Query(0, 3); v != 14 {
		t.Log("In TestSegTree_Query(): the result of seg.Query(0, 3) should be 14, but got", v)
		t.FailNow()
	}
	if v := seg.Query(0, 8); v != 68 {
		t.Log("In TestSegTree_Query(): the result of seg.Query(0, 8) should be 68, but got", v)
		t.FailNow()
	}
}

func TestSegTreeLazy_UpdateArea(t *testing.T) {
	seg := NewSegTreeLazy(a)
	seg.UpdateArea(0, 3, 5)
	if v := seg.Query(0, 3); v != 29 {
		t.Log("In TestSegTree_Update(): the result of seg.Query(0, 3) should be 29, but got", v)
		t.FailNow()
	}
	seg.UpdateArea(1, 3, 2)
	if v := seg.Query(0, 8); v != 87 {
		t.Log("In TestSegTree_Update(): the result of seg.Query(0, 8) should be 87, but got", v)
		t.FailNow()
	}
}
