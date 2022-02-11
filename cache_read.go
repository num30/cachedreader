package cachereader

import (
	"io"
)

type cachedReader struct {
	r          io.Reader
	buff       []byte
	afterReset bool
}

func NewCachedReader(r io.Reader) *cachedReader {
	return &cachedReader{
		r:    r,
		buff: []byte{},
	}
}

func (s *cachedReader) Read(p []byte) (n int, err error) {
	if s.afterReset && len(s.buff) > 0 {
		n = copy(p, s.buff)
		s.buff = s.buff[n:]
		return n, nil
	}

	n, err = s.r.Read(p)
	if !s.afterReset {
		s.buff = append(s.buff, p[:n]...)
	}

	return n, err
}

func (s *cachedReader) Reset() {
	if s.afterReset {
		panic("cachedReader.Reset() called twice. Only one call of Reset() is allowed")
	}
	s.afterReset = true
}
