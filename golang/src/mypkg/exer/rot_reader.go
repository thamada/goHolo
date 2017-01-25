//Time-stamp: <2017-01-26 01:13:12 hamada>
package exer

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rd rot13Reader) Read(buf []byte) (n int, err error) {
	var DEBUG bool = true //false

	buf0 := make([]byte, 10)
	n, err = rd.r.Read(buf0)
	if DEBUG {
		fmt.Printf("%v, %v\n", n, err)
		fmt.Printf("%v\n", buf0)
	}
	for i := 0; i < n; i++ {
		buf[i] = (((buf0[i] - 97) + 13) % 26) + 97
	}
	return
}

func Rot_reader() {

	if true {
//		s := strings.NewReader("Lbh penpxrq gur pbqr!")
//		s := strings.NewReader("a b c d e f g")
		s := strings.NewReader("Lbhpenpxrqgurpbqr!")
		r := rot13Reader{s}
		io.Copy(os.Stdout, &r)
	}
}
