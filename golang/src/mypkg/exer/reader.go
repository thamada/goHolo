//Time-stamp: <2017-01-25 19:21:01 hamada>
package exer

import (
	"golang.org/x/tour/reader"
)

type MyReader struct{}

func (r MyReader) Read(b []byte) (n int, err error) {
	n = int(1)
	b[0] = 'A'
	err = nil
	return n, err
}

func Reader() {
	reader.Validate(MyReader{})
}
