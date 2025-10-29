package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("./MyName")
	if err != nil {
		panic(err)
	}
	lenght, err := io.WriteString(file, "My name is Shobhit")
	if err != nil {
		panic(err)
	}
	fmt.Println("Length is ", lenght)
	file.Close()
	readFile("./MyName")
}

func readFile(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	
	fmt.Println(string(data))
	
}
