package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

}

//Course format
type Courses struct {
	CourseID   string  `json:"courseid"`
	CourseName string  `json:"coursename"`
	Price      float64 `json:"price"`
	Author     *Author `json:"author"`
}

//Author format
type Author struct {
	Name    string
	website string
}

//fake DB
var courses []Courses

//Check course id and name are present
func (c *Courses) IsEmpty() bool {
	return c.CourseID == "" && c.CourseName == ""
}

//HomeURL
func ServeHomeurl(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>WELCOME TO GO LMS</h1><p>Create, Add, Update & Delete Courses</p>"))
}

// GET all courses
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

//GET one Course
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseID == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode(`No Course Found with requested ID`)
	return
}
