package main

import "testing"

func TestSum(t *testing.T) {
	tables := []struct {
		x int
		y int
		r int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{3, 2, 5},
		{25, 26, 51},
	}

	for _, table := range tables {
		total := Sum(table.x, table.y)
		if total != table.r {
			t.Errorf("Sum(%d, %d) was incorrect, got: %d, want: %d.", table.x, table.y, total, table.r)
		}
	}
}

func TestMax(t *testing.T) {
	tables := []struct {
		x int
		y int
		r int
	}{
		{1, 2, 2},
		{2, 2, 2},
		{3, 2, 3},
		{25, 26, 26},
	}

	for _, table := range tables {
		max := GetMax(table.x, table.y)
		if max != table.r {
			t.Errorf("Max(%d, %d) was incorrect, got: %d, want: %d.", table.x, table.y, max, table.r)
		}
	}
}

func TestFib(t *testing.T) {
	tables := []struct {
		a int
		b int
	}{
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{13, 233},
		{14, 377},
		{15, 610},
		{21, 10946},
		{22, 17711},
		{23, 28657},
		{24, 46368},
		{28, 317811},
		{29, 514229},
		{50, 12586269025},
	}

	for _, table := range tables {
		fib := Fibonacci(table.a)
		if fib != table.b {
			t.Errorf("Fib(%d) was incorrect, got: %d, want: %d.", table.a, fib, table.b)
		}
	}
}
