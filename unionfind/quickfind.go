package unionfind

type QuickFind struct {
	slice []int
}

func NewQuickFind(size int) *QuickFind {
	slice := make([]int, size)

	for i := 0; i < size; i++ {
		slice[i] = i
	}

	return &QuickFind{slice: slice}
}

func (qf *QuickFind) Find(x int) int {
	return qf.slice[x]
}

func (qf *QuickFind) Union(x int, y int) {
	rootX := qf.Find(x)
	rootY := qf.Find(y)

	if rootX != rootY {
		for i := 0; i < len(qf.slice); i++ {
			if qf.slice[i] == rootY {
				qf.slice[i] = rootX
			}
		}
	}
}

func (qf *QuickFind) Connected(x int, y int) bool {
	return qf.Find(x) == qf.Find(y)
}
