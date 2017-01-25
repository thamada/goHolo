//Time-stamp: <2017-01-26 01:51:26 hamada>
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

func nrot(c byte, n int) byte {
	var crot byte
	switch {
	case ('A' <= c && c <= 'Z'):
		crot = (c-'A'+byte(n))%26 + 'A'
	case ('a' <= c && c <= 'z'):
		crot = (c-'a'+byte(n))%26 + 'a'
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
		b := nrot(buf0[i], 13)
		buf[i] = b
	}

	return
}

func do (msg string) {
	s := strings.NewReader(msg)
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
	fmt.Printf("\n")
}

func Rot_reader() {

	do("Lbh penpxrq gur pbqr!")
	do("Neg jnfurf njnl sebz gur fbhy gur qhfg bs rirelqnl yvsr. -- Cnoyb Cvpnffb")
	do("Vg qbrf abg znggre ubj fybjyl lbh tb fb ybat nf lbh qb abg fgbc. -- Pbhshpvhf")
	do("Vs bar nqinaprf pbasvqragyl va gur qverpgvba bs uvf qernzf, naq raqrnibef gb yvir gur yvsr juvpu ur unf vzntvarq, ur jvyy zrrg jvgu n fhpprff harkcrpgrq va pbzzba ubhef. -- Urael Qnivq Gubernh")
	do("Bhe terngrfg tybel vf abg va arire snyyvat, ohg va evfvat hc rirel gvzr jr snvy. -- Enycu Jnyqb Rzrefba")

}
