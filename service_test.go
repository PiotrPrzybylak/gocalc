package main

import "testing"

type testpair struct {
	input  string
	result string
}

var tests = []testpair{
	{"A", "a"},
	{"b", "b"},
	{"ŁódŹ", "łódź"},
}

func TestLowercase(t *testing.T) {

	svc := stringService{}

	for _, pair := range tests {
		result, _ := svc.Lowercase(pair.input)
		if result != pair.result {
			t.Error("For", pair.input,
				"expected", pair.result,
				"got", result)
		}
	}

}
