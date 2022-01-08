package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	//string lib
	var str = "hello how are you?"
	fmt.Println(str)
	fmt.Println(strings.Split(str, " "))
	fmt.Println(str)

	//sort lib

	nums := []int{10, 35, 78, 2, 34, 56}
	sort.Ints(nums)
	fmt.Println(nums)
	fmt.Println(sort.SearchInts(nums, 10))
}
