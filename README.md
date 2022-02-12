# Cached Reader
Chached Reader is an implimination of a reader that could be re-read from the biggining. 

One of the scenarios is when you need to verify file or request data by examining content in the beginning. 

## Install 
``` 
  go get github.com/testhub-io/cachedreader
```

## Example 1. Minimal example 

``` golang
	var str = []byte("the quick brown fox jumps over the lazy dog")
	r := NewCachedReader(bytes.NewReader(str)) // create cached reader
	the := make([]byte, 3)
	_, err := r.Read(the) // Read first 3 bytes
	if err != nil {
		panic(err)
	}
	fmt.Println(string(the)) // print "the"

	b, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b)) // print whole string

```

