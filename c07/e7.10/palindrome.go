package palindrome

import (
	"sort"
)

type StringSort struct {
	str string
}

func (s StringSort) Len() int           { return len(s.str)}
func (s StringSort) Less(i, j int) bool { return s.str[i] < s.str[j]}
func (s StringSort) Swap(i, j int) {

}


func IsPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len() - 1; i < j; i, j = i + 1, j -1 {
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}

	return true
}