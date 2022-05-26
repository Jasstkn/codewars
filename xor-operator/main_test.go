package main

import "testing"

func TestXor(t *testing.T) {
	tables := []struct {
		x   bool
		y   bool
		exp bool
	}{
		{false, false, false},
		{false, true, true},
		{true, false, true},
		{true, true, false},
		{true, Xor(true, true), true},
		{Xor(true, true), false, false},
		{Xor(true, false), false, true},
		{false, Xor(false, false), false},
		{Xor(false, false), Xor(false, true), true},
		{Xor(true, false), Xor(false, false), true},
		{Xor(false, false), Xor(false, false), false},
	}

	for _, table := range tables {
		ans := Xor(table.x, table.y)
		if ans != table.exp {
			t.Errorf("Xor of (%t+%t) was incorrect, got: %t, want: %t.", table.x, table.y, ans, table.exp)
		}
	}
}
