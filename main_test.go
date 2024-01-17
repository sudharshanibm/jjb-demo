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

	// Now you can perform assertions on the 'course' object
}
