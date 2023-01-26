package main

import (
	"fmt"
	"sort"
)

func main() {
	tmpSlice := []string{
		"bbbbbb",
		"baaabb",
		"baadbb",
		"cccacc",
		"aaaaaa",
	}
	fmt.Println(tmpSlice)

	SortByAlgorithmNarayana(tmpSlice)

	fmt.Println(tmpSlice)
}
func SortByAlgorithmNarayana(ar []string) { //"ABC" sort
	sort.Slice(ar, func(i, j int) bool {
		for k := 0; true; k++ {
			if k == len(ar[i]) {
				return false
			}
			if k == len(ar[j]) {
				return false
			}
			if ar[i][k] == ar[j][k] {
				continue
			}
			return ar[i][k] < ar[j][k]
		}
		return false
	})
}


//тот же принцип сортировки для [][]int
func SortByAlgorithmNarayanaINT(num [][]int) { 
	sort.Slice(num, func(i, j int) bool {
		for k := 0; true; k++ {
			if k == len(num[i]) {
				return false
			}
			if k == len(num[j]) {
				return false
			}
			if num[i][k] == num[j][k] {
				continue
			}
			return num[i][k] < num[j][k]
		}
		return false
	})
}

