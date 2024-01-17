package main

import (
	"net/http"
	"net/http/httptest"
	"os"
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
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

// Additional test functions for other API endpoints can be added similarly

func resetCourses() {
	courses = []Courses{}
}

func TestMain(m *testing.M) {
	// Set up test environment, if needed

	// Run tests
	code := m.Run()

	// Clean up test environment, if needed

	// Exit with the test status
	os.Exit(code)
}
