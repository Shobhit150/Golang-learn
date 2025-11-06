package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name     string   `json:"course_name"`
	Price    int      `json:"price"`
	Platform string   `json:"platform"`
	Password string   `json:"-"`
	Tag      []string `json:"tag,omitempty"`
}

func main() {
	EncodingJson()
}
func EncodingJson() {
	lcoCourses := []Course{
		{"React", 12, "udemy", "123", []string{"Web", "js"}},
		{"Angular", 22, "udem1", "123", []string{"Web", "js"}},
		{"Js", 13, "udmy", "123", nil},
	}

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
		{
            "course_name123": "React",
            "price": 12,
            "platform": "udemy",
            "tag": [
                    "Web",
                    "js"
            ]
		}
	`)
	var lcoCourse Course
	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Is is valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v", lcoCourse)
	} else {
		fmt.Println("Not valid")
	}

	// where i want key value pair

	fmt.Print("\n Map here \n")
	var myonlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myonlineData)
	fmt.Printf("%#v", myonlineData)

	fmt.Print("\n\n\n\n--------\n")
	for k, v := range myonlineData {
		fmt.Printf("Key is %v and value is %v and Type is %T\n", k, v, v)
	}

}
