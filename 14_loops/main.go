package main

import "fmt"

func main()  {
	days := []string{"mon", "tue", "thur"}

	for d:=0; d<len(days); d++ {
		fmt.Println(days[d])
	}

	for i,j := range days {
		fmt.Println(i, " ", j)
	}
	a := 1;
	for a<10 {
		if(a==5) {
			goto lco
		}
		fmt.Println(a)

		a++
	}
	lco: 
		fmt.Println(99)
		fmt.Println(99)
	fmt.Println(99)
}