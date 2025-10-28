package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println(time.Now())
	fmt.Println(time.Now().Format("2006-01-02"))
	fmt.Println(runtime.NumCPU())
}
