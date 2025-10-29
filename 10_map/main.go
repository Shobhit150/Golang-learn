package main

import "fmt"

func main()  {
	m := make(map[int]int)
	m[1] = 99
	m[2] = 99
	value, found :=m[1]
	if found{
		fmt.Println(value)
	}
	delete(m, 2)
	value1, found1 :=m[2]
	if found1{
		fmt.Println(value1)
	}else{
		fmt.Println(value1)
	}
}