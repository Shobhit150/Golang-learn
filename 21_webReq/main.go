package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Web Server......")

	PerformPostJsonRequest()
}

func PerformGetRequest() {
	myUrl := "http://localhost:8000/get"

	res, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	fmt.Println(res.StatusCode)
	content, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var stringbuilder strings.Builder
	stringbuilder.Write(content)
	fmt.Println(stringbuilder.String())
	// fmt.Println(string(content))
}

func PerformPostJsonRequest()  {
	const myurl = "http://localhost:8000/post"

	requestBody := strings.NewReader(`
		{
			"coursename": "Golang",
			"price": 0
		}
	`)

	res, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	var stringbuilder strings.Builder
	content, err := io.ReadAll(res.Body)
	stringbuilder.Write(content)
	fmt.Println(stringbuilder.String())

}