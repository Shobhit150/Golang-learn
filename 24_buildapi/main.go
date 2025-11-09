package main

type Course struct {
	CourseId string `json:"courseid"`
	CourseName string `json:"coursename"`
	Price int `json:"price"`
	Author *Author `json:"author"`
}
type Author struct {
	Fullname string `json:"fullname"`
	Website string 	`json:"website"`
}

var course []Course

func (c *Course) isEmpty() bool {
	return c.CourseId == "" && c.CourseName==""
}