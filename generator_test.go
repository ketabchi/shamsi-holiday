package main

import "testing"

func TestIsFriday(t *testing.T) {
	tests := []struct {
		m, d int
		exp  bool
	}{
		{7, 26, true},
		{7, 3, false},
		{12, 29, false},
		{1, 2, true},
	}

	for i, test := range tests {
		if res := isFriday(test.m, test.d); res != test.exp {
			t.Errorf("Test %d: Expected isFriday return %t, but got %t",
				i, test.exp, res)
		}
	}
}

func TestIsShamsiHoliday(t *testing.T) {
	tests := []struct {
		m, d int
		exp  bool
	}{
		{8, 5, true},
		{7, 3, false},
		{12, 29, true},
		{1, 1, true},
	}

	for i, test := range tests {
		if res := isShamsiHoliday(test.m, test.d); res != test.exp {
			t.Errorf("Test %d: Expected isShamsiHoliday return %t, but got %t",
				i, test.exp, res)
		}
	}
}
