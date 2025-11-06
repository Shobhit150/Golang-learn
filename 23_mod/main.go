package main

import (
	"fmt"
	"net/http"
)

func main()  {
	fmt.Println("Hello Mod")
	greeter()
}

func greeter() {
	fmt.Println("Hey ther")
}

func serveHome(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("<h1>Welcome to goLang</h1>"))
}