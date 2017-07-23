package htmlstringreader

import (
	"io"
)

type htmlstringreader struct {
	s string
}

func (r *htmlstringreader) Read(p []byte) (count int, err error) {
	n := copy(p, r.s)
	r.s = r.s[n:]

	if len(r.s) == 0 {
		return n, io.EOF
	}
	return n, nil
}
