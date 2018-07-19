package lab001

import "testing"

var cases = []struct {
	x, y   int
	result int
}{
	{1, 2, 3},
	{0, 0, 0},
}

func TestAdd(t *testing.T) {
	for _, c := range cases {
		r := add(c.x, c.y)
		if r != c.result {
			t.Errorf("add(%d,%d) == %d, want %d", c.x, c.y, r, c.result)
		}
	}
}
