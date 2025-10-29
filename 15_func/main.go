package main

import "fmt"

func adder(a int, b ...int) int {
	tot := a
	for _,i := range b {
		tot += i
	}
	return tot
}
func main()  {
	func () {
		fmt.Println("ok")
	}()
	fmt.Println(adder(1,2,3,));
}
