package test

import "testing"

func AssertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatal()
	}
}
