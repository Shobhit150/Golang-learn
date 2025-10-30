package main

import (
	"fmt"
	"net/url"
)

func main()  {
	// myurl := "https://www.youtube.com/watch?v=cl7_ouTMFh0&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=26"

	// result,_ := url.Parse(myurl)
	// fmt.Println(result)
	// fmt.Println("Scheme: ",result.Scheme)
	// fmt.Println("Path: ",result.Path)
	// fmt.Println("Host: ",result.Host)
	// fmt.Println("Port: ",result.Port())
	// fmt.Println("Raw Query: ",resul t.RawQuery)

	// qparams := result.Query()
	// for i,j := range qparams {
	// 	fmt.Println(i, " ", j)
	// }

	partofurls := &url.URL{
		Scheme: "https",
		Host: "google.com",
		// Path: "/",
		// RawQuery: "",
	}
	
	anotherurl := partofurls.String()
	fmt.Println(anotherurl)
}