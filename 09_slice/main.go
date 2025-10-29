package main

import (
	"fmt"
	"sort"
)

func main() {
	num := []int{1,55,6,3}
	num = append(num, 4)
	fmt.Println(num)

	// highscore := make([]int, 4)
	// fmt.Println(highscore)

	sort.Ints(num)
	fmt.Println(sort.IntsAreSorted(num))
	fmt.Println(num)
}
