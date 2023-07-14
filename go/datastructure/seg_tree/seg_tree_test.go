package seg_tree

import (
	"testing"
)

var a = []int{5, 7, 2, 12, 9, 5, 13, 15}

func TestSegTree_Query(t *testing.T) {
	seg := NewSegTree(a)
	if v := seg.Query(0, 3); v != 14 {
		t.Log("In TestSegTree_Query(): the result of seg.Query(0, 3) should be 14, but got", v)
		t.FailNow()
	}
	if v := seg.Query(0, 8); v != 68 {
		t.Log("In TestSegTree_Query(): the result of seg.Query(0, 8) should be 68, but got", v)
		t.FailNow()
	}
}

func TestSegTree_Update(t *testing.T) {
	seg := NewSegTree(a)
	seg.Update(0, 1)
	if v := seg.Query(0, 3); v != 10 {
		t.Log("In TestSegTree_Update(): the result of seg.Query(0, 3) should be 10, but got", v)
		t.FailNow()
	}
	seg.Update(5, 1)
	if v := seg.Query(0, 8); v != 60 {
		t.Log("In TestSegTree_Update(): the result of seg.Query(0, 8) should be 60, but got", v)
		t.FailNow()
	}
}
