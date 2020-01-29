package main

import "testing"

func TestGetFuel(t *testing.T) {
	testCases := []struct {
		input int
		want  int
	}{
		{14, 2},
		{1969, 966},
		{100756, 50346},
	}

	for _, tc := range testCases {
		num := getFuel(tc.input)
		if num != tc.want {
			t.Errorf("expect %d, but got %d", tc.want, num)
		}
	}
}
