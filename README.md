# Cached Reader
Chached Reader is an implimination of a reader that could be re-read from the biggining. 

One of the scenarios is when you need to verify file or request data by examining content in the beginning. 

## Install 
``` 
  go get github.com/testhub-io/cachedreader
```

## Example 1. Minimal example 

``` go
package main

import (
	"bytes"
	"fmt"
	"io"

	"github.com/testhub-io/cachedreader"
)

func main() {
	var str = []byte("the quick brown fox jumps over the lazy dog")
	r := cachedreader.NewCachedReader(bytes.NewReader(str)) // create cached reader
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
}
```

## Example 2. Check if http request body is a valid tar.gz file and store it 

``` go
// http handler 
func storeFile(w http.ResponseWriter, req *http.Request) {

	bufReader := cachedreader.NewCachedReader(req.Body)

	uncompressedStream, err := gzip.NewReader(bufReader) // uncompress the gzip stream
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Failed to read gz stream. %v", err)))
		return
	}

	tr := tar.NewReader(uncompressedStream) // create tar reader
	_, err = tr.Next() // read the file header
	if err != nil { // if error then stream is no a valid tar file
		w.Write([]byte(fmt.Sprintf("Not a tar.gz stream. %v", err)))
		return
	}
	bufReader.Reset() // reset reader to the biggining

	out, err := os.Create("filename.tar.gz")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error creating file. %v", err)))
		return
	}
	defer out.Close() 

	io.Copy(out, bufReader) // write content to a file form the beginning of the stream
}

```
