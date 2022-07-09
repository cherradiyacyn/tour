package main

import "golang.org/x/tour/reader"

type myReader struct{}

// TODO: Add a Read([]byte) (int, error) method to myReader.
func (r myReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	reader.Validate(myReader{})
}
