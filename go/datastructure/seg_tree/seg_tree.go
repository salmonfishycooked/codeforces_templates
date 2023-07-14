package seg_tree

// NOTE:
// the example segment tree records sum of [l, r)

type SegTree struct {
	ele []segnode
	ori []int
}

type segnode struct {
	l, r, v int
}

func NewSegTree(a []int) *SegTree {
	n := len(a)
	tree := &SegTree{
		ele: make([]segnode, 4*n),
		ori: a,
	}
	tree.build(0, 0, n)
	return tree
}

// build a segment tree
// node k saves the information of [l, r)
func (t *SegTree) build(k, l, r int) {
	t.ele[k].l, t.ele[k].r = l, r
	if l+1 == r {
		t.ele[k].v = t.ori[l]
		return
	}
	mid := (l + r) >> 1
	t.build(t.leftChild(k), l, mid)
	t.build(t.rightChild(k), mid, r)
	t.updateNode(k)
}

func (t *SegTree) leftChild(k int) int {
	return 2*k + 1
}

func (t *SegTree) rightChild(k int) int {
	return 2*k + 2
}

// query, the main logic

// Update changes value of the element which index is i to v
func (t *SegTree) Update(i, v int) {
	t.update(0, i, v)
}

func (t *SegTree) update(k, i, v int) {
	if t.ele[k].l+1 == t.ele[k].r && t.ele[k].l == i {
		t.ele[k].v = v
		return
	}
	mid := (t.ele[k].r + t.ele[k].l) >> 1
	if i < mid {
		t.update(t.leftChild(k), i, v)
	} else {
		t.update(t.rightChild(k), i, v)
	}
	t.updateNode(k)
}

func (t *SegTree) updateNode(k int) {
	t.ele[k].v = t.ele[t.leftChild(k)].v + t.ele[t.rightChild(k)].v
}

// Query returns a value which is the information of [l, r)
// the specific information depends on the updateNode(k int)
func (t *SegTree) Query(l, r int) int {
	return t.query(0, l, r)
}

func (t *SegTree) query(k, l, r int) int {
	if t.ele[k].l == l && t.ele[k].r == r {
		return t.ele[k].v
	}
	mid := (t.ele[k].l + t.ele[k].r) >> 1

	// logic (may change)
	lc, rc := t.leftChild(k), t.rightChild(k)
	if r <= mid {
		return t.query(lc, l, r)
	}
	if l >= mid {
		return t.query(rc, l, r)
	}
	return t.query(lc, l, mid) + t.query(rc, mid, r)
}
