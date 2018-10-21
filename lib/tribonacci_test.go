package lib

import (
	"testing"
)

func TestTribonacci(t *testing.T) {
	testCases := map[uint]uint{
		0:  0,
		1:  0,
		2:  1,
		3:  1,
		4:  2,
		5:  4,
		6:  7,
		7:  13,
		8:  24,
		9:  44,
		10: 81,
		15: 1705,
		20: 35890}

	for n, expected := range testCases {
		got := Tribonacci(n)
		if got != expected {
			t.Errorf("Expected %v tribonacci number for n=%v - got %v", expected, n, got)
		}
	}
}
