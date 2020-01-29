package main

import "testing"

func TestExecute(t *testing.T) {
	testCases := []struct {
		input map[int]int
		want  map[int]int
	}{
		{
			map[int]int{0: 1, 1: 0, 2: 0, 3: 0, 4: 99},
			map[int]int{0: 2, 1: 0, 2: 0, 3: 0, 4: 99},
		},
		{
			map[int]int{0: 2, 1: 3, 2: 0, 3: 3, 4: 99},
			map[int]int{0: 2, 1: 3, 2: 0, 3: 6, 4: 99},
		},
		{
			map[int]int{0: 2, 1: 4, 2: 4, 3: 5, 4: 99, 5: 0},
			map[int]int{0: 2, 1: 4, 2: 4, 3: 5, 4: 99, 5: 9801},
		},
		{
			map[int]int{0: 1, 1: 1, 2: 1, 3: 4, 4: 99, 5: 5, 6: 6, 7: 0, 8: 99},
			map[int]int{0: 30, 1: 1, 2: 1, 3: 4, 4: 2, 5: 5, 6: 6, 7: 0, 8: 99},
		},
	}

	for _, tc := range testCases {
		alt := execute(tc.input)
		for i, v := range alt {
			if tc.want[i] != v {
				t.Errorf("expect %d, but got %d", tc.want[i], v)
			}
		}

	}
}
