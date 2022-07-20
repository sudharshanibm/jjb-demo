package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	courses = append(courses, Courses{CourseID: "1329dfs", CourseName: "React JS", Price: 38.98, Author: &Author{Name: "Sudharshan", Website: "http://github.com/sudharshan3"}})
	courses = append(courses, Courses{CourseID: "sdaff4321", CourseName: "GO Lang", Price: 156.98, Author: &Author{Name: "Sudharshan", Website: "http://github.com/sudharshan3"}})
	r.HandleFunc("/", ServeHomeurl).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/courses", addOneCourse).Methods("POST")
	r.HandleFunc("/courses/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/courses/{id}", deleteCourse).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":4001", r))
	defer fmt.Println("Server Started Successfully \n Goto: http://localhost:4000")
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
	Name    string `json:"fullname"`
	Website string `json:"website"`
}

//fake DB
var courses []Courses

//Check course id and name are present
func (c *Courses) IsEmpty() bool {
	return c.CourseName == ""
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

//ADD new Course
func addOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please Enter some Data!")
	}
	var course Courses
	_ = json.NewDecoder(r.Body).Decode(&course)
	//checking course name is present
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Please Enter some Data!")
		return
	}
	//generating random string for course ID
	rand.Seed(time.Now().UnixNano())
	course.CourseID = strconv.Itoa(rand.Intn(1000))
	courses = append(courses, course)
	json.NewEncoder(w).Encode("Course Added Successfully")
	json.NewEncoder(w).Encode(course)

	return

}

//UPDATE Course
func updateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, course := range courses {
		if course.CourseID == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Courses
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseID = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode("Course Updated Successfully")
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("Course ID not found!")
}

//DELETE Course
func deleteCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, course := range courses {
		if course.CourseID == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course Deleted Successfully")
			break
		}
	}
}
