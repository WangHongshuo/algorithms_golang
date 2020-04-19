package test

// SortByInt :
type SortByInt []int

// Len :
func (s SortByInt) Len() int {
	return len(s)
}

// Swap:
func (s SortByInt) Swap(i, j int) {
	SwapCounter++
	s[i], s[j] = s[j], s[i]
}

// Less :
func (s SortByInt) Less(i, j int) bool {
	LessCounter++
	return s[i] < s[j]
}

// Point2 : (x, y)
type Point2 struct {
	x int
	y int
}

// SortByPoint2 :
type SortByPoint2 []Point2

// Len :
func (s SortByPoint2) Len() int {
	return len(s)
}

// Swap :
func (s SortByPoint2) Swap(i, j int) {
	SwapCounter++
	s[i], s[j] = s[j], s[i]
}

// Less :
func (s SortByPoint2) Less(i, j int) bool {
	LessCounter++
	if s[i].x < s[j].x {
		return true
	}
	if s[i].x > s[j].x {
		return false
	}
	if s[i].y < s[j].y {
		return true
	}
	return false
}
