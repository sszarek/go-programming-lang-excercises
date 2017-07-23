package limitreader

import (
	"testing"
	"strings"
	"bufio"
)

func TestLimitreader_Read(t *testing.T) {
	cases := []struct {
		str, want string
		n int64
	} {
		{"test", "te", 2},
		{"test", "t", 1},
	}

	for _, c := range cases {
		reader := LimitReader(strings.NewReader(c.str), c.n)
		scanner := bufio.NewScanner(reader)

		if scanner.Scan() {
			s := scanner.Text()
			if s != c.want {
				t.Errorf("Expected %s to equal %s", s, c.want)
			}
		}
	}
}
