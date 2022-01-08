package main

import "fmt"

func main() {
	//arrays
	var a = [5]int{1, 34, 56, 78, 98}
	for i := 1; i < len(a); i++ {
		fmt.Println(a[i])
	}

	fmt.Println("set:", a)
	a[4] = 100
	fmt.Println("get:", a[4])

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	//slices

	var numbers = []int{2, 5, 7}
	fmt.Println(numbers)
	numbers = append(numbers, 23)
	fmt.Print(numbers)

}
