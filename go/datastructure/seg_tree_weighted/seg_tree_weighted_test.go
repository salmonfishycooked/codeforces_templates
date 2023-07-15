package seg_tree_weighted

import (
	"testing"
)

var a = []int{6, 10, 12, 20, 23, 17, 32, 40}

func TestSegTreeWeighted(t *testing.T) {
	com := NewCompressor(a)
	seg := NewSegTreeWeighted(com.Len())

	seg.Insert(com.Get(a[0]))
	if v := seg.Query(0, 3); v != 1 {
		t.Log("the result of seg.Query(0, 3) should be 1, but got", v)
		t.FailNow()
	}
	seg.Insert(com.Get(a[4]))
	if v := seg.Query(0, 6); v != 2 {
		t.Log("the result of seg.Query(0, 6) should be 2, but got", v)
		t.FailNow()
	}
	seg.Insert(com.Get(a[5]))
	if v := seg.Query(0, 8); v != 3 {
		t.Log("the result of seg.Query(0, 8) should be 3, but got", v)
		t.FailNow()
	}
}
