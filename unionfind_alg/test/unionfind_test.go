package test

import (
	"testing"

	uf "github.com/WangHongshuo/algorithms_golang/unionfind_alg"
)

func Test_UnionFindV1_1(t *testing.T) {
	const errMsg = "result error"
	e := uf.NewUnionFind1()
	e.Union(1, 2)
	if l, err := e.Find(2); l != 1 || err != nil {
		t.Error(errMsg)
	}
	e.Union(3, 4)
	if l, err := e.Find(4); l != 3 || err != nil {
		t.Error(errMsg)
	}
	e.Union(1, 4)
	if l, err := e.Find(4); l != 1 || err != nil {
		t.Error(errMsg)
	}
	e.Union(4, 5)
	if l, err := e.Find(5); l != 1 || err != nil {
		t.Error(errMsg)
	}
	e.Union(6, 6)
	if e.IsInSameGroup(1, 6) != false {
		t.Error(errMsg)
	}
	if e.IsInSameGroup(6, 7) != false {
		t.Error(errMsg)
	}
	if e.IsInSameGroup(1, 5) != true {
		t.Error(errMsg)
	}
}
