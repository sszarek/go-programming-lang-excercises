package palindrome

import (
	"testing"
	"sort"
)

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		arg sort.Interface
		want bool
	} {
		{
			StringSort{"aba"}, true,
		},
		{
			StringSort{"ab"}, false,
		},
		{
			StringSort{"abcdaffadcba"}, true,
		},
		{
			StringSort{"abcdaffadcb"}, false,
		},
	}

	for _, c := range cases {
		cur := IsPalindrome(c.arg)
		if cur != c.want {
			t.Errorf("IsPalindrome: %v, expected %v to be %v", c.arg, cur, c.want)
		}
	}
}
