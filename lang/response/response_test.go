package response

import (
	"net/http"
	"testing"
)

func TestBadRequest_DroppedConnection(t *testing.T) {
	if WasBadRequest(&http.Response{}) {
		t.Fatalf("WasBadRequest should return `false` for an empty response")
	}
	if WasBadRequest(nil) {
		t.Fatalf("WasBadRequest should return `false` for a dropped connection")
	}
}

func TestBadRequest_StatusCodes(t *testing.T) {
	testCases := []struct {
		statusCode     int
		expectedResult bool
	}{
		{http.StatusOK, false},
		{http.StatusInternalServerError, false},
		{http.StatusNotFound, false},
		{http.StatusBadRequest, true},
	}

	for _, test := range testCases {
		resp := http.Response{
			StatusCode: test.statusCode,
		}
		result := WasBadRequest(&resp)
		if test.expectedResult != result {
			t.Fatalf("Expected '%+v' for status code '%d' - got '%+v'",
				test.expectedResult, test.statusCode, result)
		}
	}
}

func TestConflict_DroppedConnection(t *testing.T) {
	if WasConflict(&http.Response{}) {
		t.Fatalf("WasConflict should return `false` for an empty response")
	}
	if WasConflict(nil) {
		t.Fatalf("WasConflict should return `false` for a dropped connection")
	}
}

func TestConflict_StatusCodes(t *testing.T) {
	testCases := []struct {
		statusCode     int
		expectedResult bool
	}{
		{http.StatusOK, false},
		{http.StatusInternalServerError, false},
		{http.StatusNotFound, false},
		{http.StatusConflict, true},
	}

	for _, test := range testCases {
		resp := http.Response{
			StatusCode: test.statusCode,
		}
		result := WasConflict(&resp)
		if test.expectedResult != result {
			t.Fatalf("Expected '%+v' for status code '%d' - got '%+v'",
				test.expectedResult, test.statusCode, result)
		}
	}
}

func TestNotFound_DroppedConnection(t *testing.T) {
	if WasNotFound(&http.Response{}) {
		t.Fatalf("WasNotFound should return `false` for an empty response")
	}
	if WasNotFound(nil) {
		t.Fatalf("WasNotFound should return `false` for a dropped connection")
	}
}

func TestNotFound_StatusCodes(t *testing.T) {
	testCases := []struct {
		statusCode     int
		expectedResult bool
	}{
		{http.StatusOK, false},
		{http.StatusInternalServerError, false},
		{http.StatusNotFound, true},
	}

	for _, test := range testCases {
		resp := http.Response{
			StatusCode: test.statusCode,
		}
		result := WasNotFound(&resp)
		if test.expectedResult != result {
			t.Fatalf("Expected '%+v' for status code '%d' - got '%+v'",
				test.expectedResult, test.statusCode, result)
		}
	}
}
