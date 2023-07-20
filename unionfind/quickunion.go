package unionfind

type QuickUnion struct {
	slice []int
}

func NewQuickUnion(size int) *QuickUnion {
	slice := make([]int, size)

	for i := 0; i < size; i++ {
		slice[i] = i
	}

	return &QuickUnion{slice: slice}
}

func (qf *QuickUnion) Find(x int) int {
	for x != qf.slice[x] {
		x = qf.slice[x]
	}
	return x
}

func (qf *QuickUnion) Union(x int, y int) {
	rootX := qf.Find(x)
	rootY := qf.Find(y)

	if rootX != rootY {
		qf.slice[rootY] = rootX
	}
}

func (qf *QuickUnion) Connected(x int, y int) bool {
	return qf.Find(x) == qf.Find(y)
}
