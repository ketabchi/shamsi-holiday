package main

import "testing"

func TestToDateFormat(t *testing.T) {
	j, m := toDateFormat(12, 29)

	if j != "1399/12/29" {
		t.Errorf("Expected 1398/12/29 got %s", j)
	}
	if m != "2021/03/19" {
		t.Errorf("Expected 2020/03/19 got %s", m)
	}
}
