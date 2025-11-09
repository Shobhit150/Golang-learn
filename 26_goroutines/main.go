package main

import (
	"fmt"
	"net/http"
)

func main()  {
	arr := []string{
		"https://www.google.com",
		"https://www.fb.com",
		"https://www.github.com",
	}
	for _,i:= range arr {
		greeter(i)
	}
	
}

func greeter(str string)  {
	res, err := http.Get(str)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d and %s\n", res.StatusCode, str)
	
}