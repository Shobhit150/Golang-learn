package main

import "fmt"

func main()  {
	a := 2
	if a==2 {
		fmt.Println("jo")
	}
	if a:=2; a==2 {
		fmt.Println(2)
	}
}