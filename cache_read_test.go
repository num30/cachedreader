package cachereader

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

var str = []byte("the quick brown fox jumps over the lazy dog")

func Test_Read(t *testing.T) {
	reader := NewCachedReader(bytes.NewReader(str))
	b, err := io.ReadAll(reader)
	if err != nil {
		t.Error(err)
	} else if !bytes.Equal(str, b) {
		t.Errorf("expected %s, got %s", str, b)
	}
}

func Test(t *testing.T) {
	tests := []struct {
		name     string
		preReads func(io.Reader)
	}{
		{
			name: "PreReadByte",
			preReads: func(r io.Reader) {
				r.Read(make([]byte, 1))
			},
		},
		{
			name: "PreReadAll",
			preReads: func(r io.Reader) {
				r.Read(make([]byte, len(str)))
			},
		},
		{
			name: "PreReadNothing",
			preReads: func(r io.Reader) {
				r.Read(make([]byte, 0))
			},
		},
		{
			name: "PreReadMoreThenStream",
			preReads: func(r io.Reader) {
				r.Read(make([]byte, 1024))
			},
		},
		{
			name: "PreReadTwice",
			preReads: func(r io.Reader) {
				r.Read(make([]byte, 1))
				r.Read(make([]byte, 2))
			},
		},
		{
			name: "PreReadAllByOne",
			preReads: func(r io.Reader) {
				for i := 0; i < len(str); i++ {
					r.Read(make([]byte, 1))
				}
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			reader := NewCachedReader(bytes.NewReader(str))
			test.preReads(reader)
			reader.Reset()
			b, err := io.ReadAll(reader)
			if err != nil {
				t.Error(err)
			} else if !bytes.Equal(str, b) {
				t.Errorf("expected %s, got %s", str, b)
			}
		})
	}
}

func Test_ResetTwice(t *testing.T) {
	reader := NewCachedReader(bytes.NewReader(str))
	reader.Reset()

	defer func() {
		// recover from panic if one occured. Set err to nil otherwise.
		if recover() == nil {
			t.Error("expected panic")
		}
	}()
	reader.Reset()
}

// read reader until EOF
func readAll(r io.Reader) {
	b, err := io.ReadAll(r)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s", string(b))
}
