package seg_tree_lazy

// NOTE:
// the example lazy segment tree records sum of [l, r)
// the function UpdateArea() is used by increasing original values by v

type SegTreeLazy struct {
	ele []segNode
}

type segNode struct {
	l, r int
	v    int
	// lazy flags
	incr int
}

func NewSegTreeLazy(nums []int) *SegTreeLazy {
	n := len(nums)
	tree := &SegTreeLazy{
		ele: make([]segNode, 4*n),
	}
	tree.build(0, 0, n, nums)
	return tree
}

// build a segment tree
// node k saves the information of [l, r)
func (t *SegTreeLazy) build(k, l, r int, nums []int) {
	t.ele[k].l, t.ele[k].r = l, r
	if l+1 == r {
		t.ele[k].v = nums[l]
		return
	}
	mid := (l + r) >> 1
	t.build(t.leftChild(k), l, mid, nums)
	t.build(t.rightChild(k), mid, r, nums)
	t.updateNode(k)
}

func (t *SegTreeLazy) leftChild(k int) int {
	return 2*k + 1
}

func (t *SegTreeLazy) rightChild(k int) int {
	return 2*k + 2
}

// AreaIncr changes value of [l, r)
// using lazy flag
func (t *SegTreeLazy) AreaIncr(l, r, v int) {
	t.areaIncr(0, l, r, v)
}

func (t *SegTreeLazy) areaIncr(k, l, r, v int) {
	if t.ele[k].l == l && t.ele[k].r == r {
		t.lazyIncr(k, v)
		return
	}
	if t.isLazy(k) {
		t.lazyDown(k)
	}
	mid := (t.ele[k].l + t.ele[k].r) >> 1
	lc, rc := t.leftChild(k), t.rightChild(k)
	if r <= mid {
		t.areaIncr(lc, l, r, v)
	} else if l >= mid {
		t.areaIncr(rc, l, r, v)
	} else {
		t.areaIncr(lc, l, mid, v)
		t.areaIncr(rc, mid, r, v)
	}
	t.updateNode(k)
}

func (t *SegTreeLazy) updateNode(k int) {
	t.ele[k].v = t.ele[t.leftChild(k)].v + t.ele[t.rightChild(k)].v
}

// isLazy returns true if the node k is lazy.
func (t *SegTreeLazy) isLazy(k int) bool {
	return t.ele[k].incr != 0
}

// lazyIncr makes the node k being lazy
func (t *SegTreeLazy) lazyIncr(k, v int) {
	t.ele[k].incr += v
	t.ele[k].v += (t.ele[k].r - t.ele[k].l) * v
}

// lazyDown clears lazy flag of the node k and passes it to its children
func (t *SegTreeLazy) lazyDown(k int) {
	t.lazyIncr(t.leftChild(k), t.ele[k].incr)
	t.lazyIncr(t.rightChild(k), t.ele[k].incr)
	t.ele[k].incr = 0
}

// Query returns the sum of [l, r)
// using lazy flag
func (t *SegTreeLazy) Query(l, r int) int {
	return t.query(0, l, r)
}

func (t *SegTreeLazy) query(k, l, r int) int {
	if t.ele[k].l == l && t.ele[k].r == r {
		return t.ele[k].v
	}
	if t.isLazy(k) {
		t.lazyDown(k)
	}
	mid := (t.ele[k].l + t.ele[k].r) >> 1
	if r <= mid {
		return t.query(t.leftChild(k), l, r)
	}
	if l >= mid {
		return t.query(t.rightChild(k), l, r)
	}
	return t.query(t.leftChild(k), l, mid) + t.query(t.rightChild(k), mid, r)
}
