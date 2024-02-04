package word

import "testing"

func TestUseCount(t *testing.T) {
	m := UseCount("one two three three")
	for k, v := range m {
		switch k {
		case "one":
			if v != 1 {
				t.Error("expected:", 1, "got:", v)
			}
		case "two":
			if v != 1 {
				t.Error("expected:", 1, "got:", v)
			}
		case "three":
			if v != 2 {
				t.Error("expected:", 2, "got:", v)
			}
		}
	}
}
