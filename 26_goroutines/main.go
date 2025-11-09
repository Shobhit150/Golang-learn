package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var wb = []string{}
var mu sync.Mutex

func main()  {
	arr := []string{
		"https://www.google.com",
		"https://www.fb.com",
		"https://www.github.com",
	}
	
	for _,i:= range arr {
		wg.Add(1)
		go greeter(i)
		wg.Add(1)
		go greeter(i)
		
	}

	wg.Wait()

	fmt.Print(wb)
	
}

func greeter(str string)  {
	defer wg.Done()
	_, err := http.Get(str)
	if err != nil {
		panic(err)
	}
	mu.Lock()
	wb = append(wb, str)
	mu.Unlock()
	// fmt.Printf("%d and %s\n", res.StatusCode, str)
	
}