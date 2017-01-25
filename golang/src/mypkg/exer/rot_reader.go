//Time-stamp: <2017-01-26 01:42:18 hamada>
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

func rot13(c byte) byte {
	var crot byte
	switch {
	case ('A' <= c && c <= 'Z'):
		crot = (c-'A'+13)%26 + 'A'
	case ('a' <= c && c <= 'z'):
		crot = (c-'a'+13)%26 + 'a'
	default:
		crot = c
	}
	return crot
}

func (rd rot13Reader) Read(buf []byte) (n int, err error) {
	var DEBUG bool = false // true //

	buf0 := make([]byte, 10)
	n, err = rd.r.Read(buf0)

	if DEBUG {
		fmt.Printf("%v, %v\n", n, err)
		fmt.Printf("%v\n", buf0)
	}

	if err != nil {
		return 0, err
	}

	for i := range buf0 {
		b := rot13(buf0[i])
		buf[i] = b
	}

	return
}

func Rot_reader() {

	//		s := strings.NewReader("Lbh penpxrq gur pbqr!")
	s := strings.NewReader("Neg jnfurf njnl sebz gur fbhy gur qhfg bs rirelqnl yvsr. -- Cnoyb Cvpnffb")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
	fmt.Printf("\n")

}
