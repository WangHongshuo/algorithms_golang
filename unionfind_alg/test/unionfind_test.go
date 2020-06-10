package test

import (
	"fmt"
	"testing"

	uf "github.com/WangHongshuo/algorithms_golang/unionfind_alg"
)

func Test_UnionFindV1_1(t *testing.T) {
	e := uf.NewUnionFind1()
	e.Union(1, 2)
	fmt.Println(e.Find(2))
	e.Union(3, 4)
	fmt.Println(e.Find(4))
	e.Union(1, 4)
	fmt.Println(e.Find(4))
	e.Union(4, 5)
	fmt.Println(e.Find(5))
	e.Union(6, 6)
	fmt.Println(e.IsInSameGroup(1, 6))
	fmt.Println(e.IsInSameGroup(6, 7))
	fmt.Println(e)
}
