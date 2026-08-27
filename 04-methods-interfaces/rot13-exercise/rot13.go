package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r13 rot13Reader) Read(b []byte) (int, error) {
	n, err := r13.r.Read(b)
	if err != nil {
		return n, err
	}
	for i := 0; i < n; i++ { // n being number of items read
		switch {
		case 'A' <= b[i] && b[i] <= 'Z': // b is the buffer we are working with
			b[i] = 'A' + (b[i]-'A'+13)%26 // add 13, mod 26 to wrap around, add 'A' to get back to ASCII range
		case 'a' <= b[i] && b[i] <= 'z': // must handle upper and lowercase
			b[i] = 'a' + (b[i]-'a'+13)%26
		}
	}
	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

// rot13 - cipher - rotate 13 characters
// special because 2 consecutive applications result in original text
