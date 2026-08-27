package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	b[0] = 'A'
	return 1, nil
}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func main() {
	reader.Validate(MyReader{})
}

// this was a reading check:
// Read - the interface method we are focusing on, takes in slice
// implementation: populate slice with data, return number of bytes read and error if any
