package unionfind

import "testing"

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
