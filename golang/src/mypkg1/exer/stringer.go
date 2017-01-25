//Time-stamp: <2017-01-25 14:33:11 hamada>
package exer

import "fmt"

type IPAddr [4]byte

func (i IPAddr) String() string {
	s := fmt.Sprintf("%v.%v.%v.%v", i[0], i[1], i[2], i[3])
	return s
}

func Stringer() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
