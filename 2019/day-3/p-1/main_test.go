package main

import "testing"

func TestGetDistance(t *testing.T) {
	testCases := []struct {
		input  [][]vector
		output int
	}{
		{
			[][]vector{
				{
					{RIGHT, 75}, {DOWN, 30}, {RIGHT, 83}, {UP, 83}, {LEFT, 12}, {DOWN, 49}, {RIGHT, 71}, {UP, 7}, {LEFT, 72},
				},
				{
					{UP, 62}, {RIGHT, 66}, {UP, 55}, {RIGHT, 34}, {DOWN, 71}, {RIGHT, 55}, {DOWN, 58}, {RIGHT, 83},
				},
			},
			159,
		},
		{
			[][]vector{
				{
					{RIGHT, 98}, {UP, 47}, {RIGHT, 26}, {DOWN, 63}, {RIGHT, 33}, {UP, 87}, {LEFT, 62}, {DOWN, 20}, {RIGHT, 33}, {UP, 53}, {RIGHT, 51},
				},
				{
					{UP, 98}, {RIGHT, 91}, {DOWN, 20}, {RIGHT, 16}, {DOWN, 67}, {RIGHT, 40}, {UP, 7}, {RIGHT, 15}, {UP, 6}, {RIGHT, 7},
				},
			},
			135,
		},
	}

	_ = testCases
}
