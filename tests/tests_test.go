package main

import "testing"

/*
* Prefix all test functions with Test or it will be ignored
*
* All Test functions will take the argument of (t *testing.T)
*
* Cmd to run tests is:
* $ go test
* OR use verbose flag to get more details about test run
* $ go test -v
*
* For test coverage:
* $ go test -cover
*
* For test coverage report:
* $ go test -coverprofile=coverage.out && go tool cover -html=coverage.out
 */

// Slow way to create tests
func TestDivide(t *testing.T) {
	// check for err in divide function
	_, err := divide(10.0, 1.0)
	if err != nil {
		// use testing error and not a regular error in tests
		t.Error("Got an error when we should not have")
	}
}

func TestBadDivide(t *testing.T) {
	// check for err in divide function
	_, err := divide(10.0, 0)
	if err == nil {
		// use testing error and not a regular error in tests
		t.Error("Got an error when we should have")
	}
}

// faster and more thorough testing method
var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"error-data", 10.0, 0.0, 0, true},
	{"expect-5", 50.0, 10.0, 5.0, false},
	{"expect-fraction", -1.0, -777.0, 0.0012870013, false},
}

func TestDivision(t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.dividend, tt.divisor)
		if tt.isErr {
			if err == nil {
				t.Error(tt.name, "expected an error but did not get one")
			}
		} else {
			if err != nil {
				t.Error(tt.name, "Did not expect an error but got one", err.Error())
			}
		}

		if got != tt.expected {
			t.Errorf("%v expected %f but got %f", tt.name, tt.expected, got)
		}
	}
}
