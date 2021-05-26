package main

import (
	"strings"
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestGenerateHandler(t *testing.T) {
	type test struct {
		body string
		expectedStatus   int
		expectedBody  string
	}

	tests := []test{
		{
			body: "",
			expectedStatus: http.StatusOK,
			expectedBody: "",
		},
		{
			body: "foobar",
			expectedStatus: http.StatusOK,
			expectedBody: "",
		},
		{
			body: "### TEMPLATE ###\nfoobar",
			expectedStatus: http.StatusOK,
			expectedBody: "foobar\n",
		},
		{
			body: "### VALUES ###\nexample: testing\n### TEMPLATE ###\nfoobar",
			expectedStatus: http.StatusOK,
			expectedBody: "foobar\n",
		},
		{
			body: "### VALUES ###\nexample: variable value\n### TEMPLATE ###\n{{ .Values.example }}",
			expectedStatus: http.StatusOK,
			expectedBody: "variable value\n",
		},
	}

	for _, tc := range tests {
		reader := strings.NewReader(tc.body)
		req, err := http.NewRequest("POST", "/generate", reader)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(GenerateHandler)
		handler.ServeHTTP(rr, req)
		if status := rr.Code; status != tc.expectedStatus {
			t.Errorf("handler returned wrong status code: got %v want %v", status, tc.expectedStatus)
		}
		if rr.Body.String() != tc.expectedBody {
			t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), tc.expectedBody)
		}
	}
}
