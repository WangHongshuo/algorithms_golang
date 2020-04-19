package test

import (
	"fmt"
	"reflect"
	"testing"

	sortalg "github.com/WangHongshuo/algorithms_golang/sort_alg"
)

func Test_QuickSortV1_1(t *testing.T) {
	SwapCounter, LessCounter = 0, 0
	sortalg.QuickSortV1(SortByInt(IntsArray1))
	fmt.Printf("%v\n", IntsArray1)
	fmt.Println("SwapCounter: ", SwapCounter, " LessCounter: ", LessCounter)
	if !reflect.DeepEqual(IntsArray1, IntsArraySortResult) {
		t.Error("result error")
	}
}

func Test_QuickSortV1_2(t *testing.T) {
	SwapCounter, LessCounter = 0, 0
	sortalg.QuickSortV1(SortByInt(IntsArray2))
	fmt.Printf("%v\n", IntsArray2)
	fmt.Println("SwapCounter: ", SwapCounter, " LessCounter: ", LessCounter)
	if !reflect.DeepEqual(IntsArray2, IntsArraySortResult) {
		t.Error("result error")
	}
}

func Test_QuickSortV1_3(t *testing.T) {
	SwapCounter, LessCounter = 0, 0
	sortalg.QuickSortV1(SortByPoint2(Point2Array))
	fmt.Printf("%v\n", Point2Array)
	fmt.Println("SwapCounter: ", SwapCounter, " LessCounter: ", LessCounter)
	if !reflect.DeepEqual(Point2Array, Point2ArraySortResult) {
		t.Error("result error")
	}
}
