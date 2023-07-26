package seg_tree_weighted

import "sort"

type SegTreeWeighted struct {
	ele []segNode
	up  int
}

type segNode struct {
	v int
}

// NewSegTreeWeighted [0, up)
func NewSegTreeWeighted(up int) *SegTreeWeighted {
	tree := &SegTreeWeighted{
		ele: make([]segNode, 4*up),
		up:  up,
	}
	return tree
}

func (t *SegTreeWeighted) Insert(v int) {
	t.insert(0, v, 0, t.up)
}

func (t *SegTreeWeighted) insert(k, v, l, r int) {
	if l+1 == r {
		t.ele[k].v++
		return
	}
	mid := (l + r) >> 1
	lc, rc := 2*k+1, 2*k+2
	if v < mid {
		t.insert(lc, v, l, mid)
	} else {
		t.insert(rc, v, mid, r)
	}
	t.ele[k].v++
}

// Query [l, r)
func (t *SegTreeWeighted) Query(l, r int) int {
	return t.query(0, 0, t.up, l, r)
}

func (t *SegTreeWeighted) query(k, l, r, ql, qr int) int {
	if l == ql && r == qr {
		return t.ele[k].v
	}
	mid := (l + r) >> 1
	if qr <= mid {
		return t.query(2*k+1, l, mid, ql, qr)
	} else if ql >= mid {
		return t.query(2*k+2, mid, r, ql, qr)
	}
	return t.query(2*k+1, l, mid, ql, mid) + t.query(2*k+2, mid, r, mid, qr)
}

// Compressor do the discretizing of original data
// Compressor 将原始数据进行离散化
type Compressor struct {
	f map[int]int
}

func NewCompressor(a []int) *Compressor {
	inst := &Compressor{}
	inst.mapTo(a)
	return inst
}

func (c *Compressor) mapTo(a []int) {
	tmp := make([]int, len(a))
	copy(tmp, a)
	sort.Ints(tmp)
	mp := make(map[int]int, len(a))
	idx := 0
	for i, n := 0, len(tmp); i < n; i++ {
		if i > 0 && tmp[i] == tmp[i-1] {
			continue
		}
		mp[tmp[i]] = idx
		idx++
	}
	c.f = mp
}

// Get returns the mapped value by original value(v)
func (c *Compressor) Get(v int) int {
	return c.f[v]
}

func (c *Compressor) Len() int {
	return len(c.f)
}

// Compressor Version 2 for common questions
//type Compressor struct {
//	f []int
//}
//
//func NewCompressor(a []int) *Compressor {
//	inst := &Compressor{}
//	inst.mapTo(a)
//	return inst
//}
//
//func (c *Compressor) mapTo(a []int) {
//	mp := make(map[int]struct{})
//	for _, v := range a {
//		mp[v] = struct{}{}
//	}
//	c.f = make([]int, len(mp))
//	idx := 0
//	for i := range mp {
//		c.f[idx] = i
//		idx++
//	}
//	sort.Ints(c.f)
//}
//
//func (c *Compressor) up(v int) int {
//	idx := sort.Search(len(c.f), func(i int) bool {
//		return c.f[i] > v
//	})
//	return idx
//}
//
//func (c *Compressor) lo(v int) int {
//	idx := sort.Search(len(c.f), func(i int) bool {
//		return c.f[i] >= v
//	})
//	return idx
//}
//
//func (c *Compressor) Len() int {
//	return len(c.f)
//}
