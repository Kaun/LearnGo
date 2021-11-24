package main

import "fmt"

func SearchBinary(slice []int, item int) (int, bool) {
	h := len(slice) - 1
	l := 0
	var point int
	for i := 0; i < len(slice); i++ {
		point = (h + l) / 2
		if l == h && item != slice[point] {
			return 0, false
		} else if item == slice[point] {
			return point, true
		} else if item > slice[point] {
			l = point + 1
		} else if item < slice[point] {
			h = point - 1
		}
	}
	return 0, false
}

func MinItem(slc []int) {
	for i, v := range slc {
		fmt.Println(i, v)
	}

}

func main() {

	slc := []int{1, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20}

	MinItem(slc)

}
