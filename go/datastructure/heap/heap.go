package heap

type Heap struct {
	ele []int
}

func NewHeap(nums []int) *Heap {
	heap := &Heap{}
	// will modify on original data
	heap.Build(nums)

	// or use copy
	// tmp := make([]int, len(nums))
	// copy(tmp, nums)
	// heap.Build(tmp)
	return heap
}

func (h *Heap) Build(nums []int) {
	n := len(nums)
	h.ele = nums
	for idx := n/2 - 1; idx >= 0; idx-- {
		h.down(idx)
	}
}

func (h *Heap) Pop() int {
	n := len(h.ele)
	v := h.ele[0]
	h.swap(0, n-1)
	h.ele = h.ele[:n-1]
	h.down(0)
	return v
}

func (h *Heap) Push(x int) {
	h.ele = append(h.ele, x)
	h.up(len(h.ele) - 1)
}

func (h *Heap) down(i int) {
	n := len(h.ele)
	for {
		left := i<<1 + 1
		if left >= n {
			break
		}
		right := left + 1
		// idx is an index which has the higher priority in the son nodes of i
		idx := left
		if left < n && right < n && h.compare(right, left) {
			idx = right
		}
		if !h.compare(idx, i) {
			break
		}
		h.swap(i, idx)
		i = idx
	}
}

func (h *Heap) up(i int) {
	for {
		dad := (i - 1) / 2 // parent
		if i == dad || !h.compare(i, dad) {
			break
		}
		h.swap(i, dad)
		i = dad
	}
}

func (h *Heap) swap(i, j int) {
	h.ele[i], h.ele[j] = h.ele[j], h.ele[i]
}

func (h *Heap) compare(i, j int) bool {
	return h.ele[i] > h.ele[j]
}
