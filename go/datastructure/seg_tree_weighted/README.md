# Example Questions

Box In Box: https://atcoder.jp/contests/abc309/tasks/abc309_f

```go
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)

	n := getI()
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, 3)
		a[i][0], a[i][1], a[i][2] = getI(), getI(), getI()
		sort.Ints(a[i])
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i][0] < a[j][0] || a[i][0] == a[j][0] && a[i][1] > a[j][1]
	})
	w := make([]int, n)
	for i, v := range a {
		w[i] = v[1]
	}
	com := NewCompressor(w)
	seg := NewSegTreeWeighted(com.Len())
	for _, v := range a {
		if seg.Query(0, com.Get(v[1])) < v[2] {
			out("Yes")
			return
		}
		seg.Insert(com.Get(v[1]), v[2])
	}
	out("No")
}

func out(x ...interface{}) {
	fmt.Fprintln(wr, x...)
}

func getI() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type SegTreeWeighted struct {
	ele []segnode
	up  int
}

type segnode struct {
	v int
}

func NewSegTreeWeighted(up int) *SegTreeWeighted {
	tree := &SegTreeWeighted{
		ele: make([]segnode, 4*up),
		up:  up,
	}
	return tree
}

func (t *SegTreeWeighted) Insert(v, d int) {
	t.insert(0, v, d, 0, t.up)
}

func (t *SegTreeWeighted) insert(k, v, d, l, r int) {
	if l+1 == r {
		t.updateNode(k, d)
		return
	}
	mid := (l + r) >> 1
	lc, rc := 2*k+1, 2*k+2
	if v < mid {
		t.insert(lc, v, d, l, mid)
	} else {
		t.insert(rc, v, d, mid, r)
	}
	t.updateNode(k, d)
}

func (t *SegTreeWeighted) updateNode(k int, d int) {
	if t.ele[k].v == 0 {
		t.ele[k].v = d
		return
	}
	t.ele[k].v = min(t.ele[k].v, d)
}

func (t *SegTreeWeighted) Query(l, r int) int {
	return t.query(0, 0, t.up, l, r)
}

func (t *SegTreeWeighted) query(k, l, r, ql, qr int) int {
	if l == ql && r == qr {
		if t.ele[k].v == 0 {
			return math.MaxInt32
		}
		return t.ele[k].v
	}
	mid := (l + r) >> 1
	if qr <= mid {
		return t.query(2*k+1, l, mid, ql, qr)
	} else if ql >= mid {
		return t.query(2*k+2, mid, r, ql, qr)
	}
	return min(t.query(2*k+1, l, mid, ql, mid), t.query(2*k+2, mid, r, mid, qr))
}

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

func (c *Compressor) Get(v int) int {
	return c.f[v]
}

func (c *Compressor) Len() int {
	return len(c.f)
}

```

