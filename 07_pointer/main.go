package main

import "fmt"

func nice(a *int, b *int) int {
	return *a + *b
}

func main() {
	b := 3
	a := &b
	fmt.Println(*a)
}
