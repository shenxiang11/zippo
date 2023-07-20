package unionfind

import (
	"testing"
)

func TestQuickUnion(t *testing.T) {
	qf := NewQuickUnion(5)

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
