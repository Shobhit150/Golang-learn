package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name string `json:"course_name"`
	Price int `json:"price"`
	Platform string `json:"platform"`
	Password string `json:"-"`
	Tag []string `json:"tag,omitempty"`
}
func main()  {
	EncodingJson()
}
func EncodingJson()  {
	lcoCourses := []Course{
		{"React", 12, "udemy", "123", []string{"Web","js"}}, 
		{"Angular", 22, "udem1", "123", []string{"Web","js"}},
		{"Js", 13, "udmy", "123", nil},
	}

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err!= nil {
		panic(err)
	}
	fmt.Printf("%s",finalJson)
}
