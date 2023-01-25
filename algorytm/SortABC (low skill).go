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
	sort.Slice(tmpSlice, func(i, j int) bool { return tmpSlice[i-1] > tmpSlice[j+1] })
	fmt.Println("By name:", tmpSlice)

	//Output:
	//By name: [{Alice 55} {Bob 75} {Gopher 7} {Vera 24}]
	//By age: [{Gopher 7} {Vera 24} {Alice 55} {Bob 75}]
}
