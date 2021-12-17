package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type AddResult struct {
	x 		 int
	y 		 int
	expected int
}

// TestCalculate is simple Test
func TestCalculate(t *testing.T) {
	if Calculate(2) != 4 {
		t.Error("Exoected 2 + 2 equal 4")
	}
}

// It must include Test at Start and rest does not matter
func TestTableCalculate(t *testing.T) {
	// slice of struct
	// Anonymous struct
	var tests = []struct{
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
		{99999, 100001},
	}

	for _, test := range tests {
		if output := Calculate(test.input); output != test.expected {
			t.Errorf("test: Failed: %d inputted, %d Expected %d recieved", test.input, test.expected, output)
		}
	}
}

var addResults = []AddResult {
	{1, 1, 2},
	{0, 1, 1},
	{2, 2, 4},
	{1, 4, 5},
}

// Technique 1: TABLE driven test
// TesAdd is Example of TABLE driven test technique
func TestAdd(t *testing.T) {
	for _, test := range addResults {
		result := Add(test.x, test.y)
		if result != test.expected {
			t.Fatal("Expected Result Not Given")
		}
	}
}

// Technique 2: Use of test data directory
func TestReadFile(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/test.data")
	if err != nil {
		t.Fatal("could not open file")
	}
	if string(data) != "hello world" {
		t.Fatal("String content do not match expected")
	}
}

// Technique 3
// Mocking http request for API
func TestHttpRequest(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{\"status\": \"good\"}")
	}

	req := httptest.NewRequest("GET", "https://tutorialedge.net", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	if 200 != resp.StatusCode {
		t.Fatal("Status code not OK")
	}
}