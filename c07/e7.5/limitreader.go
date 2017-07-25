package limitreader

import "io"

type limitreader struct {
	limit int64
	inner io.Reader
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitreader{n, r}
}

func (r *limitreader) Read(p []byte) (n int, err error) {
	n, err = r.inner.Read(p[:r.limit])
	if err != nil {
		return
	}

	r.limit -= int64(n)
	if r.limit == 0 {
		err = io.EOF
	}

	return
}
