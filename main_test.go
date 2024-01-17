package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

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

	expected := `[{"courseid":"1329dfs","coursename":"React JS","price":38.98,"author":{"fullname":"Sudharshan","website":"http://github.com/sudharshan3"}},{"courseid":"sdaff4321","coursename":"GO Lang","price":156.98,"author":{"fullname":"Sudharshan","website":"http://github.com/sudharshan3"}}]`
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body:\n got  %v\n want %v", rr.Body.String(), expected)
	}
}

// Additional test functions for other API endpoints can be added similarly

func resetCourses() {
	courses = []Courses{
		{CourseID: "1329dfs", CourseName: "React JS", Price: 38.98, Author: &Author{Name: "Sudharshan", Website: "http://github.com/sudharshan3"}},
		{CourseID: "sdaff4321", CourseName: "GO Lang", Price: 156.98, Author: &Author{Name: "Sudharshan", Website: "http://github.com/sudharshan3"}},
	}
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

	var expected Courses
	for _, c := range courses {
		if c.CourseID == "1329dfs" {
			expected = c
			break
		}
	}

	var actual Courses
	err = json.NewDecoder(rr.Body).Decode(&actual)
	if err != nil {
		t.Fatal(err)
	}

	if !coursesEqual(actual, expected) {
		t.Errorf("Handler returned unexpected body:\n got  %v\n want %v", actual, expected)
	}
}

// Additional test functions for other API endpoints can be added similarly

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

// Add similar functions for testing other endpoints like addOneCourse, updateCourse, deleteCourse, etc.
