package unionfind

import "testing"

func TestNewUnionFind(t *testing.T) {
	qf := NewUnionFind(5)

	// 测试初始状态下是否正确地连接
	if qf.Connected(0, 1) || qf.Connected(1, 2) || qf.Connected(3, 4) {
		t.Errorf("TestQuickFind failed. Initial connections are incorrect.")
	}

	// 连接元素
	qf.Union(0, 1)
	qf.Union(1, 2)
	qf.Union(3, 4)

	// 测试连接是否正确
	if !qf.Connected(0, 1) || !qf.Connected(1, 2) || !qf.Connected(3, 4) {
		t.Errorf("TestQuickFind failed. Connections after unions are incorrect.")
	}

	if !qf.Connected(0, 2) || qf.Connected(1, 4) {
		t.Errorf("TestQuickFind failed. Incorrect connections.")
	}
}

func BenchmarkQuickFind(b *testing.B) {
	size := 10000
	qu := NewQuickFind(size)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// 进行一些合并和查找操作
		for j := 0; j < size/2; j++ {
			qu.Union(j, size-1-j)
		}
		for j := 0; j < size/2; j++ {
			qu.Connected(j, size-1-j)
		}
	}
}

// BenchmarkQuickUnion measures the performance of the QuickUnion algorithm.
func BenchmarkQuickUnion(b *testing.B) {
	size := 10000
	qu := NewQuickUnion(size)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 进行一些合并和查找操作
		for j := 0; j < size/2; j++ {
			qu.Union(j, size-1-j)
		}
		for j := 0; j < size/2; j++ {
			qu.Connected(j, size-1-j)
		}
	}
}

// BenchmarkQuickUnion measures the performance of the UnionFind algorithm.
func BenchmarkUnionFind(b *testing.B) {
	size := 10000
	qu := NewUnionFind(size)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 进行一些合并和查找操作
		for j := 0; j < size/2; j++ {
			qu.Union(j, size-1-j)
		}
		for j := 0; j < size/2; j++ {
			qu.Connected(j, size-1-j)
		}
	}
}
