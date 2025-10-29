package main

import "fmt"

type User struct{
	a int
	b string
}
func main()  {
	user := User{1, "Shob"}
	fmt.Printf("%+v", user)
}
