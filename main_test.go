package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeHomeurl(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ServeHomeurl)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "<h1>WELCOME TO GO LMS</h1><p>Create, Add, Update & Delete Courses</p>"
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body:\n got  %v\n want %v", rr.Body.String(), expected)
	}
}

func TestGetAllCourses(t *testing.T) {
	resetCourses()

	req, err := http.NewRequest("GET", "/courses", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getAllCourses)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var courses []Courses
	err = json.Unmarshal(rr.Body.Bytes(), &courses)
	if err != nil {
		t.Errorf("Error unmarshalling JSON: %v", err)
	}

	if len(courses) != 2 {
		t.Errorf("Expected 2 courses, got %v", len(courses))
	}

	// You can add more specific assertions based on your API response
}

func TestGetOneCourse(t *testing.T) {
	resetCourses()

	req, err := http.NewRequest("GET", "/courses/1329dfs", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getOneCourse)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var course Courses
	err = json.Unmarshal(rr.Body.Bytes(), &course)
	if err != nil {
		t.Errorf("Error unmarshalling JSON: %v", err)
	}

	expectedCourse := Courses{
		CourseID:   "1329dfs",
		CourseName: "React JS",
		Price:      38.98,
		Author: &Author{
			Name:    "Sudharshan",
			Website: "http://github.com/sudharshan3",
		},
	}

	if !coursesEqual(expectedCourse, course) {
		t.Errorf("Handler returned unexpected body:\n got  %v\n want %v", course, expectedCourse)
	}
}

func TestAddOneCourse(t *testing.T) {
	resetCourses()

	newCourse := Courses{
		CourseName: "New Course",
		Price:      99.99,
		Author: &Author{
			Name:    "New Author",
			Website: "http://example.com",
		},
	}

	jsonBody, _ := json.Marshal(newCourse)

	req, err := http.NewRequest("POST", "/courses", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(addOneCourse)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var responseCourses []Courses
	err = json.Unmarshal(rr.Body.Bytes(), &responseCourses)
	if err != nil {
		t.Errorf("Error unmarshalling JSON: %v", err)
	}

	if len(responseCourses) != 1 {
		t.Errorf("Expected 1 course, got %v", len(responseCourses))
	}

	// You can add more specific assertions based on your API response
}

func TestUpdateCourse(t *testing.T) {
	resetCourses()

	updateCourse := Courses{
		CourseName: "Updated Course",
		Price:      55.55,
		Author: &Author{
			Name:    "Updated Author",
			Website: "http://updated.com",
		},
	}

	jsonBody, _ := json.Marshal(updateCourse)

	req, err := http.NewRequest("PUT", "/courses/1329dfs", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(updateCourse)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var responseCourses []Courses
	err = json.Unmarshal(rr.Body.Bytes(), &responseCourses)
	if err != nil {
		t.Errorf("Error unmarshalling JSON: %v", err)
	}

	if len(responseCourses) != 1 {
		t.Errorf("Expected 1 course, got %v", len(responseCourses))
	}

	// You can add more specific assertions based on your API response
}

func TestDeleteCourse(t *testing.T) {
	resetCourses()

	req, err := http.NewRequest("DELETE", "/courses/1329dfs", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(deleteCourse)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// You can add more specific assertions based on your API response
}

// Add similar test functions for other API endpoints
// Remember to handle edge cases and error scenarios
func resetCourses() {
	courses = []Courses{
		{CourseID: "1329dfs", CourseName: "React JS", Price: 38.98, Author: &Author{Name: "Sudharshan", Website: "http://github.com/sudharshan3"}},
		{CourseID: "sdaff4321", CourseName: "GO Lang", Price: 156.98, Author: &Author{Name: "Sudharshan", Website: "http://github.com/sudharshan3"}},
	}
}

func coursesEqual(c1, c2 Courses) bool {
	return c1.CourseID == c2.CourseID &&
		c1.CourseName == c2.CourseName &&
		c1.Price == c2.Price &&
		authorEqual(c1.Author, c2.Author)
}

func authorEqual(a1, a2 *Author) bool {
    if a1 == nil || a2 == nil {
        return a1 == a2
    }
    return a1.Name == a2.Name && a1.Website == a2.Website
}
