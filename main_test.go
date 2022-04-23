package main

import (
	"testing"
)

// table tests
var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isError  bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0.0, true},
	{"expect-5", 100.0, 20.0, 5.0, false},
}

func TestDivision(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := divide(test.dividend, test.divisor)
			if test.isError {
				if err == nil {
					t.Error("expected an error but got nil")
				}
			} else {
				if err != nil {
					t.Error("expected no error but got", err.Error())
				}
			}

			if result != test.expected {
				t.Errorf("expected %f, but got %f", test.expected, result)
			}
		})
	}
}
