//Time-stamp: <2017-01-25 19:13:25 hamada>
package tut

import (
	"fmt"
	"io"
	"strings"
)

func Reader() {
	r := strings.NewReader("Hello, Reader!")

	fmt.Printf("%v, %T\n", r, r)

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
