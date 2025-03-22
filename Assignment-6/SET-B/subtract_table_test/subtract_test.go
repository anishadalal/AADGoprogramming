package subtract

import "testing"

func TestSubtract(t *testing.T) {
	cases := []struct {
		a, b, expected int
	}{
		{5, 3, 2},
		{10, 4, 6},
	}

	for _, c := range cases {
		if Subtract(c.a, c.b) != c.expected {
			t.Errorf("Expected %d", c.expected)
		}
	}
}
