package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a := rand.Intn(6) + 1
	switch a {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(1)
		fallthrough
	case 3:
		fmt.Println(1)
	case 4:
		fmt.Println(1)
	case 5:
		fmt.Println(1)
	case 6:
		fmt.Println(1)
	default:
		fmt.Println(1)
	}
}
