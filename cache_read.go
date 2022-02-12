package cachedreader

import (
	"io"
)

type CachedReader struct {
	r          io.Reader
	buff       []byte
	afterReset bool
}

// NewCachedReader returns a new instance of CachedReader.
func NewCachedReader(r io.Reader) *CachedReader {
	return &CachedReader{
		r:    r,
		buff: []byte{},
	}
}

func (s *CachedReader) Read(p []byte) (n int, err error) {
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

// Reset resets the reader to the bigining of the stream.
func (s *CachedReader) Reset() {
	if s.afterReset {
		panic("CachedReader.Reset() called twice. Only one call of Reset() is allowed")
	}
	s.afterReset = true
}
