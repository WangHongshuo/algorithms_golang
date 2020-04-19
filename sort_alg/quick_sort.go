package sortalg

// QuickSortV1 :
func QuickSortV1(data Interface) {
	n := data.Len()
	lo := 0
	hi := n - 1
	quickSortV1(data, lo, hi)
}

func quickSortV1(data Interface, lo, hi int) {
	if lo >= hi {
		return
	}
	v, l, r := lo, lo, hi
	for {
		for {
			if l == r {
				break
			}
			if data.Less(r, v) {
				data.Swap(r, v)
				v = r
				l++
				break
			}
			r--
		}
		for {
			if l == r {
				break
			}
			if data.Less(v, l) {
				data.Swap(l, v)
				v = l
				r--
				break
			}
			l++
		}
		if l == r {
			break
		}
	}
	quickSortV1(data, lo, l-1)
	quickSortV1(data, l+1, hi)
}

// QuickSortV2 :
func QuickSortV2(data Interface) {
	n := data.Len()
	lo := 0
	hi := n - 1
	quickSortV2(data, lo, hi)
}

func quickSortV2(data Interface, lo, hi int) {

}
