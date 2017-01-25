//Time-stamp: <2017-01-26 00:25:14 hamada>
package exer

import (
	"io"
	"os"
	"strings"
	"fmt"
)

type rot13Reader struct {
	r io.Reader
}

func (rd rot13Reader) Read([]byte) (int, error) {
	fmt.Printf("%v\n", rd.r)
	return 0, nil
}

func Rot_reader() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
