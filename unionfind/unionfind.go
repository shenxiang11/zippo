package unionfind

type UnionFind struct {
	slice []int
	rank  []int
}

func NewUnionFind(size int) *UnionFind {
	slice := make([]int, size)
	rank := make([]int, size)

	for i := 0; i < size; i++ {
		slice[i] = i
		rank[i] = 1
	}

	return &UnionFind{slice: slice, rank: rank}
}

func (qf *UnionFind) Find(x int) int {
	if x == qf.slice[x] {
		return x
	}
	qf.slice[x] = qf.Find(qf.slice[x])
	return qf.slice[x]
}

func (qf *UnionFind) Union(x int, y int) {
	rootX := qf.Find(x)
	rootY := qf.Find(y)

	if rootX != rootY {
		if qf.rank[rootX] > qf.rank[rootY] {
			qf.slice[rootY] = rootX
		} else if qf.rank[rootX] < qf.rank[rootY] {
			qf.slice[rootX] = rootY
		} else {
			qf.slice[rootX] = rootY
			qf.rank[rootY] += 1
		}
	}
}

func (qf *UnionFind) Connected(x int, y int) bool {
	return qf.Find(x) == qf.Find(y)
}
