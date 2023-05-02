// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

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

func TestForbidden_DroppedConnection(t *testing.T) {
	if WasForbidden(&http.Response{}) {
		t.Fatalf("WasForbidden should return `false` for an empty response")
	}
	if WasForbidden(nil) {
		t.Fatalf("WasForbidden should return `false` for a dropped connection")
	}
}

func TestForbidden_StatusCodes(t *testing.T) {
	testCases := []struct {
		statusCode     int
		expectedResult bool
	}{
		{http.StatusOK, false},
		{http.StatusInternalServerError, false},
		{http.StatusNotFound, false},
		{http.StatusForbidden, true},
	}

	for _, test := range testCases {
		resp := http.Response{
			StatusCode: test.statusCode,
		}
		result := WasForbidden(&resp)
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

func TestWasStatusCode_DroppedConnection(t *testing.T) {
	if WasStatusCode(&http.Response{}, http.StatusOK) {
		t.Fatalf("WasStatusCode should return `false` for an empty response")
	}
	if WasStatusCode(nil, http.StatusOK) {
		t.Fatalf("WasStatusCode should return `false` for a dropped connection")
	}
}

func TestWasStatusCode_StatusCodes(t *testing.T) {
	testCases := []struct {
		returnedStatusCode int
		checkForStatusCode int
		result             bool
	}{
		{
			checkForStatusCode: http.StatusOK,
			returnedStatusCode: http.StatusOK,
			result:             true,
		},
		{
			checkForStatusCode: http.StatusOK,
			returnedStatusCode: http.StatusInternalServerError,
			result:             false,
		},
		{
			checkForStatusCode: http.StatusInternalServerError,
			returnedStatusCode: http.StatusInternalServerError,
			result:             true,
		},
		{
			checkForStatusCode: http.StatusInternalServerError,
			returnedStatusCode: http.StatusOK,
			result:             false,
		},
		{
			checkForStatusCode: http.StatusNotFound,
			returnedStatusCode: http.StatusNotFound,
			result:             true,
		},
		{
			checkForStatusCode: http.StatusOK,
			returnedStatusCode: http.StatusNotFound,
			result:             false,
		},
	}

	for _, test := range testCases {
		resp := http.Response{
			StatusCode: test.returnedStatusCode,
		}
		actual := WasStatusCode(&resp, test.checkForStatusCode)
		if test.result != actual {
			t.Fatalf("expected %t but got %t for status codes %d (returned) and %d (checking for)", test.result, actual, test.returnedStatusCode, test.checkForStatusCode)
		}
	}
}
