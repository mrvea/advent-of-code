package vacation42

import (
	"strconv"
	"testing"
)

func TestByr(t *testing.T) { //": maxMin(2001, 1920, 4), //four digits; at least 1920 and at most 2002
	testCases := []struct {
		input string
		want  bool
	}{
		{
			"2002",
			true,
		},
		{
			"2003",
			false,
		},
	}
	for _, tc := range testCases {
		alt := byr(tc.input)
		if alt != tc.want {
			t.Errorf("expect %s, but got %s => %s", strconv.FormatBool(tc.want), strconv.FormatBool(alt), tc.input)
		}

	}
}
func TestIyr(t *testing.T) { //": maxMin(2020, 2010, 4), //four digits; at least 2010 and at most 2020

}
func TestEyr(t *testing.T) { //": maxMin(2030, 2020, 4), //four digits; at least 2020 and at most 2030

}
func TestHgt(t *testing.T) {
	testCases := []struct {
		input string
		want  bool
	}{
		{
			"60in",
			true,
		},
		{
			"190cm",
			true,
		},
		{
			"190in",
			false,
		},
		{
			"190",
			false,
		},
	}

	for _, tc := range testCases {
		alt := hgt(tc.input)
		if alt != tc.want {
			t.Errorf("expect %s, but got %s => %s", strconv.FormatBool(tc.want), strconv.FormatBool(alt), tc.input)
		}

	}
}
func TestHcl(t *testing.T) {
	testCases := []struct {
		input string
		want  bool
	}{
		{
			"#123abc",
			true,
		},
		{
			"#123abz",
			false,
		},
		{
			"123abc",
			false,
		},
	}

	for _, tc := range testCases {
		alt := hcl(tc.input)
		if alt != tc.want {
			t.Errorf("expect %s, but got %s => %s", strconv.FormatBool(tc.want), strconv.FormatBool(alt), tc.input)
		}

	}
}
func TestEcl(t *testing.T) {
	testCases := []struct {
		input string
		want  bool
	}{
		{
			"brn",
			true,
		},
		{
			"amb",
			true,
		},
		{
			"blu",
			true,
		},
		{
			"gry",
			true,
		},
		{
			"grn",
			true,
		},
		{
			"hzl",
			true,
		},
		{
			"oth",
			true,
		},
		{
			"wat",
			false,
		},
	}
	for _, tc := range testCases {
		alt := ecl(tc.input)
		if alt != tc.want {
			t.Errorf("expect %s, but got %s", strconv.FormatBool(tc.want), strconv.FormatBool(alt))
		}

	}
}
func TestPid(t *testing.T) {
	testCases := []struct {
		input string
		want  bool
	}{
		{
			"000000001",
			true,
		},
		{
			"0123456789",
			false,
		},
	}
	for _, tc := range testCases {
		alt := pid(tc.input)
		if alt != tc.want {
			t.Errorf("expect %s, but got %s", strconv.FormatBool(tc.want), strconv.FormatBool(alt))
		}

	}
}
