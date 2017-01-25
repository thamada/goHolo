//Time-stamp: <2017-01-25 14:22:46 hamada>
package exer

import (
	"fmt"
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	sf := strings.Fields(s)
	m := make(map[string]int)
	for i, c := range sf {
		m[c] = m[c] + 1
		if false {
			fmt.Printf("%v: %v, %T\n", i, sf[i], c)
		}
	}
	if false {
		fmt.Printf("%v, %T\n", m, m)
	}
	return m
}

func Maps() {

	wc.Test(WordCount)

}
