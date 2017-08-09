package uiaatoi

import "testing"

var atoi_tests = []struct {
	a        string
	expected int
}{
	{"42", 42},
	{"-42", -42},
	{"0", 0},
}

func TestUiaAtoi(t *testing.T) {
	for _, v := range atoi_tests {
		if val, _ := Uiaatoi(v.a); val != v.expected {
			//t.Fatalf("at SomeSig(%f) returned %t, expected %t", v.f, val, v.expected)
			t.Errorf("UiaAtoi(%s) returned %d, expected %d", v.a, val, v.expected)
		}
	}
}
